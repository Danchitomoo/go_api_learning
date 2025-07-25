package apperrors

type MyAppError struct {
	ErrCode // type is ErrCode. when the field name is skipped, it is the same one of type.
	Message string
	Err     error `json:"-"` // prevent from being encoded in json.
}

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}
