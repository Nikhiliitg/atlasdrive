package file

import "time"

type File struct {
	ID        string
	Name      string
	FolderID  string
	OwnerID   string
	CreatedAt time.Time
}

func NewFile(id , name , folderID , OwnerID string)(*File , error){
	if id == ""{
		return nil , ErrInvalidFileId
	}
	if name == ""{
		return nil , ErrInvalidFileName
	}
	if folderID == ""{
		return nil , ErrInvalidFolderId
	}
	if OwnerID == ""{
		return nil , ErrInvalidOwner
	}
	return &File{
		ID: id,
		Name: name,
		FolderID: folderID,
		OwnerID: OwnerID,
		CreatedAt: time.Now(),
	}, nil
}
