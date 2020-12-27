package storage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type auditCreateUpdateInterface interface {
	Creating(userID primitive.ObjectID)
	Updating(userID primitive.ObjectID)
	GetUpdatingModelID() primitive.ObjectID
}

type auditDownloadInterface interface {
	Downloading(userID primitive.ObjectID)
	GetDownloadedModelID() primitive.ObjectID
}

type idInterface interface {
	GetID() primitive.ObjectID
}
