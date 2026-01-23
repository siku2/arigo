package errors

import "fmt"

type JSONRPCError struct {
	Code    ExitStatus `json:"code"`    // error code
	Message string     `json:"message"` // The human readable error message associated to Code
}

func (e *JSONRPCError) Error() string {
	return fmt.Sprintf("code=%d, message=%s", e.Code, e.Message)
}

