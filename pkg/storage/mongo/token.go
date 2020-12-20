package storage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type tokenModel struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	TokenID        string             `bson:"tokenId"`
	UserID         primitive.ObjectID `bson:"userId"`
	ExpirationTime int64              `bson:"expirationTime"`
	Token          string             `bson:"tokenString"`
}
