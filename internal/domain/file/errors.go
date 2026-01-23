package file

import "errors"

var (
	ErrInvalidFileId = errors.New("invalid file id")
	ErrInvalidFileName = errors.New("invalid file name")
	ErrInvalidFolderId = errors.New("invalid folder id")
	ErrInvalidOwner = errors.New("invalid owner")
)