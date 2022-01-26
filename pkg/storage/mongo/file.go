package storage

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//FileModel represents file model for sql communication
type fileModel struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	Name       string             `bson:"name"`
	Path       string             `bson:"path"`
	Owner      primitive.ObjectID `bson:"owner"`
	SearchTags string             `bson:"searchTags"`

	//auditCreateUpdateInterface
	CreatedAt time.Time          `bson:"createdAt"`
	CreatedBy primitive.ObjectID `bson:"createdBy"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	UpdatedBy primitive.ObjectID `bson:"updatedBy"`
}

func (m fileModel) GetID() primitive.ObjectID {
	return m.ID
}

func (m fileModel) Creating(userID primitive.ObjectID) {
	m.CreatedAt = time.Now()
	m.CreatedBy = userID
	m.UpdatedAt = time.Now()
	m.UpdatedBy = userID
}

func (m fileModel) Updating(userID primitive.ObjectID) {
	m.UpdatedAt = time.Now()
	m.UpdatedBy = userID
}

func (m fileModel) GetUpdatingModelID() primitive.ObjectID {
	return m.GetID()
}
