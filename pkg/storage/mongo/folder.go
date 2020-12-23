package storage

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//FileModel represents file model for sql communication
type folderModel struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	FullPath   string             `bson:"fullPath"`
	BelongsTo  primitive.ObjectID `bson:"belongsTo"`
	CreatedAt  time.Time          `bson:"uploadedAt"`
	ModifiedAt time.Time          `bson:"modifiedAt"`
}
