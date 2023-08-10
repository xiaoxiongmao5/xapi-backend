package myerror

type AbortError struct {
	Code    int
	Message string
}

func (e *AbortError) Error() string {
	return e.Message
}

func NewAbortErr(code int, msg string) *AbortError {
	return &AbortError{
		Code:    code,
		Message: msg,
	}
}

var ResponseCodes map[string]int
