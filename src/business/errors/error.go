package errors

// ErrorDuplicate ...
type ErrorDuplicate struct {
	Message string
}

func (e ErrorDuplicate) Error() string {
	return e.Message
}

// ErrUnAuthorize ...
type ErrUnAuthorize struct {
	Message string
}

func (e ErrUnAuthorize) Error() string {
	return e.Message
}

// ErrBadRequest ..
type ErrBadRequest struct {
	Message string
}

func (e ErrBadRequest) Error() string {
	return e.Message
}

// ErrNotFound ...
type ErrNotFound struct {
	Message string
}

func (e ErrNotFound) Error() string {
	return e.Message
}

// ErrNotAllowed ...
type ErrNotAllowed struct {
	Message string
}

func (e ErrNotAllowed) Error() string {
	return e.Message
}
