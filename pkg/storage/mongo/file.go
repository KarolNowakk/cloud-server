package storage

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//FileModel represents file model for sql communication
type fileModel struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	Extension  string             `bson:"extension"`
	FullPath   string             `bson:"fullPath"`
	BelongsTo  primitive.ObjectID `bson:"belongsTo"`
	Size       int64              `bson:"size"`
	UploadedAt time.Time          `bson:"uploadedAt"`
	ModifiedAt time.Time          `bson:"modifiedAt"`
}
