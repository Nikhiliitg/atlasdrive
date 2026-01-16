package folder

import "time"

type Folder struct{
	ID string
	Name string
	OwnerID string
	ParentID *string
	CreatedAt time.Time
}

func NewFolder(
	id string , name string , ownerID string , parentID *string )(*Folder , error){

		if id == "" || name == "" || ownerID == "" {
			return nil , ErrInvalidOwner
		}

		if parentID != nil && *parentID == id {
		return nil, ErrCycleDetected
	}
	return &Folder{
		ID: id,
		Name: name,
		OwnerID: ownerID,
		ParentID: parentID,
		CreatedAt: time.Now(),
	}, nil
}