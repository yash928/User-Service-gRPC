package user

import "errors"

var (
	ErrDocumentNotFound     = errors.New("document not found")
	ErrInvalidID            = errors.New("invalid id")
	ErrUserNotFound         = errors.New("user not found")
	ErrInvalidMaritalStatus = errors.New("invalid marital status")
	ErrSomethingWentWrong   = errors.New("something went wrong")
)
