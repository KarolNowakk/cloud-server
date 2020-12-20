package storage

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userModel struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string             `bson:"username"`
	Email        string             `bson:"email"`
	Password     string             `bson:"password"`
	JoinedAt     time.Time          `bson:"joinedAt"`
	LastActivity time.Time          `bson:"lastActivity"`
	Tokens       []tokenModel       `bson:"tokens"`
}
