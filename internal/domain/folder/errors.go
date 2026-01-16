package folder

import "errors"

var (
	ErrInvalidFolderID = errors.New("invalid folder id")
	ErrInvalidFolderName = errors.New("invalid folder name")
	ErrInvalidOwner = errors.New("invalid owner")
	ErrCycleDetected = errors.New("cycle detected")
)
