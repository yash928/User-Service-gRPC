package user

import "errors"

var (
	ErrDocumentNotFound   = errors.New("document not found")
	ErrInvalidID          = errors.New("invalid id")
	ErrUserNotFound       = errors.New("user not found")
	ErrSomethingWentWrong = errors.New("something went wrong")
)
