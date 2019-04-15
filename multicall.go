package arigo

import "encoding/json"

// MethodCallError represents an error returned by aria2 for a MethodCall
type MethodCallError struct {
	Code    uint `json:",string"`
	Message string
}

func (e *MethodCallError) Error() string {
	return e.Message
}

// MethodResult represents the result of a MethodCall
// in a MultiCall operation.
type MethodResult struct {
	// Raw JSON value returned
	Result []byte

	// Error encountered during the method call.
	// This is likely to be a MethodCallError but it's
	// not guaranteed.
	Error error
}

func (res *MethodResult) Unmarshal(v interface{}) error {
	if res.Error != nil {
		return res.Error
	}

	err := json.Unmarshal(res.Result, v)
	return err
}

type MethodCall struct {
	// Method name to call
	MethodName string `json:"methodName"`
	// Parameters to pass to the method
	Params []interface{} `json:"params"`
}

func NewMethodCall(methodName string, params ...interface{}) MethodCall {
	return MethodCall{
		MethodName: methodName,
		Params:     params,
	}
}
