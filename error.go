package scholarship_api

// ErrorDuplicate ...
type ErrorDuplicate struct {
	Message string
}

func (e ErrorDuplicate) Error() string {
	return e.Message
}

// ErrorUnAuthorize ...
type ErrorUnAuthorize struct {
	Message string
}

func (e ErrorUnAuthorize) Error() string {
	return e.Message
}
