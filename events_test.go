package arigo

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEventTarget(t *testing.T) {
	var evtTarget eventTarget

	var events []*DownloadEvent

	unsub := evtTarget.Subscribe(StartEvent, func(event *DownloadEvent) {
		events = append(events, event)
	})

	evtTarget.Dispatch(StartEvent, &DownloadEvent{"1"})
	evtTarget.Dispatch(CompleteEvent, &DownloadEvent{"2"})
	evtTarget.Dispatch(StartEvent, &DownloadEvent{"3"})

	unsub()
	unsub()

	evtTarget.Dispatch(StartEvent, &DownloadEvent{"4"})

	require.Len(t, events, 2, "should only receive two events")

	first := events[0].GID
	assert.Equal(t, "1", first)

	second := events[1].GID
	assert.Equal(t, "3", second)
}
