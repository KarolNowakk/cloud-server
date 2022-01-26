package storage

import (
	"cloud/pkg/auth"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

//NewAuthStorageService return an instance of UserStorageService
func NewAuthStorageService(db *mongo.Database) *AuthStorageService {
	return &AuthStorageService{
		userColl:  db.Collection("users"),
		tokenColl: db.Collection("tokens"),
	}
}

//AuthStorageService is app storage for users
type AuthStorageService struct {
	userColl  *mongo.Collection
	tokenColl *mongo.Collection
}

//CreateUser creates new user and saves it into database
func (s *AuthStorageService) CreateUser(ctx context.Context, user auth.User) error {
	newUserModel := userModel{}
	newUserModel.Email = user.Email
	newUserModel.JoinedAt = time.Now()

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}

	newUserModel.Password = string(encryptedPassword)

	res, err := s.userColl.InsertOne(ctx, newUserModel)
	if err != nil {
		return err
	}
	if _, ok := res.InsertedID.(primitive.ObjectID); !ok {
		return ErrInvalidLastInsertedID
	}

	return nil
}

//ValueExists tells if value exists on db
func (s *AuthStorageService) ValueExists(ctx context.Context, field, value string) bool {
	var user userModel
	filter := bson.M{field: value}

	if err := s.userColl.FindOne(ctx, filter).Decode(&user); err != nil {
		return false
	}

	return true
}

//FindUser tels if value exists on db
func (s *AuthStorageService) FindUser(ctx context.Context, field, value string) (auth.User, error) {
	var user userModel
	filter := bson.M{field: value}

	if err := s.userColl.FindOne(ctx, filter).Decode(&user); err != nil {
		return auth.User{}, err
	}

	authUser := auth.User{}
	authUser.ID = user.ID.Hex()
	authUser.Email = user.Email
	authUser.Password = user.Password

	return authUser, nil
}

//CreateToken creates new token and saves it into database
func (s *AuthStorageService) CreateToken(ctx context.Context, token auth.Token) error {
	userID, err := primitive.ObjectIDFromHex(token.UserID)
	if err != nil {
		return err
	}

	newTokenModel := tokenModel{}
	newTokenModel.TokenID = token.TokenID
	newTokenModel.UserID = userID
	newTokenModel.Token = token.TokenString

	res, err := s.tokenColl.InsertOne(ctx, newTokenModel)
	if err != nil {
		return err
	}
	if _, ok := res.InsertedID.(primitive.ObjectID); !ok {
		return ErrInvalidLastInsertedID
	}

	return nil
}

//FindUserByHex finds user by hex value
func (s *AuthStorageService) FindUserByHex(ctx context.Context, id string) (auth.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return auth.User{}, err
	}

	var user userModel
	filter := bson.M{"_id": objectID}

	if err := s.userColl.FindOne(ctx, filter).Decode(&user); err != nil {
		return auth.User{}, err
	}

	authUser := auth.User{}
	authUser.ID = user.ID.Hex()
	authUser.Email = user.Email
	authUser.Password = user.Password

	return authUser, nil
}

//FindUsersTokens finds all users tokens
func (s *AuthStorageService) FindUsersTokens(ctx context.Context, userID string) ([]auth.Token, error) {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"userId": objectID}

	cursor, err := s.tokenColl.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var tokens []auth.Token
	var token tokenModel
	var authToken auth.Token

	for cursor.Next(ctx) {
		if err := cursor.Decode(&token); err != nil {
			return nil, err
		}

		authToken.TokenID = token.TokenID
		authToken.UserID = token.UserID.Hex()
		authToken.ExpirationTime = token.ExpirationTime
		authToken.TokenString = token.Token

		tokens = append(tokens, authToken)
	}

	if len(tokens) < 1 {
		return nil, errors.New("not found")
	}

	return tokens, nil
}
