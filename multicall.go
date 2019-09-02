package arigo

import (
	"encoding/json"
)

// MethodCallError represents an error returned by aria2 for a MethodCall
type MethodCallError struct {
	Code    uint   `json:"code,string"`
	Message string `json:"message"`
}

func (e *MethodCallError) Error() string {
	return e.Message
}

// MethodResult represents the result of a MethodCall
// in a MultiCall operation.
type MethodResult struct {
	Result []byte // Raw JSON value returned

	// Error encountered during the method call.
	// This is likely to be a MethodCallError but it's
	// not guaranteed.
	Error error
}

// Unmarshal unmarshals the raw result into v.
// If the result contains an error, it is returned directly
// without ever even attempting to unmarshal the result.
func (res *MethodResult) Unmarshal(v interface{}) error {
	if res.Error != nil {
		return res.Error
	}

	err := json.Unmarshal(res.Result, v)
	return err
}

// MethodCall represents a method call in a multi call operation
type MethodCall struct {
	MethodName string        `json:"methodName"` // Method name to call
	Params     []interface{} `json:"params"`     // Parameters to pass to the method
}

// NewMethodCall creates a new MethodCall
func NewMethodCall(methodName string, params ...interface{}) *MethodCall {
	return &MethodCall{
		MethodName: methodName,
		Params:     params,
	}
}

// TODO create constructors for all methods
