package post

type ErrInvalidPost struct {
	Err error
}

func (e *ErrInvalidPost) Error() string { return "Invalid post: " + e.Err.Error() }
func (e *ErrInvalidPost) Unwrap() error { return e.Err }
