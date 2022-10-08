package arigo

import "sync"

//go:generate stringer -type=EventType

// EventType represents a DownloadEvent which can be subscribed to on the
// Client.
type EventType uint

const (
	// StartEvent is dispatched when a download is started
	StartEvent EventType = iota
	// PauseEvent is dispatched when a download is paused
	PauseEvent
	// StopEvent is dispatched when a download is stopped
	StopEvent
	// CompleteEvent is dispatched when a download is completed
	CompleteEvent
	// BTCompleteEvent is dispatched when a BitTorrent download is completed
	BTCompleteEvent
	// ErrorEvent is dispatched when an error occurs
	ErrorEvent
)

// DownloadEvent represents the event emitted by aria2 concerning downloads.
// It only contains the gid of the download.
type DownloadEvent struct {
	GID string
}

func (e *DownloadEvent) String() string {
	return e.GID
}

// EventListener represents a function which should be called
// when an event occurs.
type EventListener func(event *DownloadEvent)

// UnsubscribeFunc is a function which when called, unsubscribes from an
// event.
type UnsubscribeFunc func() bool

// EventSubscriber is an interface which can be subscribed to.
type EventSubscriber interface {
	Subscribe(evtType EventType, listener EventListener) UnsubscribeFunc
}

// EventDispatcher is an interface capable of dispatching events.
type EventDispatcher interface {
	Dispatch(evtType EventType, event *DownloadEvent)
}

type listenerData struct {
	f  EventListener
	id uint64
}

type eventTarget struct {
	listenerMap map[EventType][]listenerData
	currentID   uint64
	mut         sync.RWMutex
}

func (t *eventTarget) unsubscribe(evtType EventType, id uint64) bool {
	t.mut.Lock()
	defer t.mut.Unlock()

	listeners := t.listenerMap[evtType]
	if len(listeners) == 0 {
		return false
	}

	found := false
	for i, l := range listeners {
		if l.id == id {
			found = true
			listeners[i] = listeners[len(listeners)-1]
			listeners = listeners[:len(listeners)-1]
			break
		}
	}

	if !found {
		return false
	}

	if len(listeners) > 0 {
		t.listenerMap[evtType] = listeners
	} else {
		delete(t.listenerMap, evtType)
	}

	return true
}

func (t *eventTarget) Subscribe(evtType EventType, listener EventListener) UnsubscribeFunc {
	t.mut.Lock()
	defer t.mut.Unlock()

	id := t.currentID
	t.currentID++

	if t.listenerMap == nil {
		t.listenerMap = make(map[EventType][]listenerData)
	}

	t.listenerMap[evtType] = append(t.listenerMap[evtType], listenerData{id: id, f: listener})

	return func() bool {
		return t.unsubscribe(evtType, id)
	}
}

func (t *eventTarget) Dispatch(evtType EventType, event *DownloadEvent) {
	t.mut.RLock()
	defer t.mut.RUnlock()

	var wg sync.WaitGroup

	listeners := t.listenerMap[evtType]

	wg.Add(len(listeners))
	for _, listener := range listeners {
		go func(l listenerData) {
			l.f(event)
			wg.Done()
		}(listener)
	}

	wg.Wait()
}
