package storage

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//FileModel represents file model for sql communication
type folderModel struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	Name     string             `bson:"name"`
	FullPath string             `bson:"fullPath"`
	Owner    primitive.ObjectID `bson:"owner"`

	//auditCreateUpdateInterface
	CreatedAt time.Time          `bson:"createdAt"`
	CreatedBy primitive.ObjectID `bson:"createdBy"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	UpdatedBy primitive.ObjectID `bson:"updatedBy"`
}

func (m *folderModel) GetID() primitive.ObjectID {
	return m.ID
}

func (m *folderModel) Creating(userID primitive.ObjectID) {
	m.CreatedAt = time.Now()
	m.CreatedBy = userID
	m.UpdatedAt = time.Now()
	m.UpdatedBy = userID
}

func (m *folderModel) Updating(userID primitive.ObjectID) {
	m.UpdatedAt = time.Now()
	m.UpdatedBy = userID
}

func (m *folderModel) GetUpdatingModelID() primitive.ObjectID {
	return m.GetID()
}
