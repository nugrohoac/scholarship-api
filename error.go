package scholarship_api

// ErrorDuplicate ...
type ErrorDuplicate struct {
	Message string
}

func (e ErrorDuplicate) Error() string {
	return e.Message
}
