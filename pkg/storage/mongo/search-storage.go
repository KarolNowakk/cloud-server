package storage

import (
	"cloud/pkg/search"
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//NewSearchStorageService return an instance of ScannerStorageService
func NewSearchStorageService(db *mongo.Database) *SearchStorageService {
	return &SearchStorageService{
		coll: db.Collection("files"),
	}
}

//SearchStorageService scans database and looks for updates
type SearchStorageService struct {
	coll *mongo.Collection
}

func (s *SearchStorageService) SearchByPhrase(phrase, userID string) ([]search.SearchedFile, error) {
	ctx := context.Background()

	owner, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	regex := bson.M{"$regex": phrase, "$options": ""}
	filter := bson.M{"owner": owner,
		"$or": []interface{}{
			bson.M{"name": regex},
			bson.M{"searchTags": regex},
		},
	}

	cursor, err := s.coll.Find(ctx, filter)
	if err != nil {
		log.Error("Mongo error: %v", err)
		return nil, err
	}

	files := []search.SearchedFile{}

	for cursor.Next(ctx) {
		var model fileModel

		if err := cursor.Decode(&model); err != nil {
			log.Error("Error decoding: %v", err)
			return nil, err
		}

		var searchedFile search.SearchedFile

		searchedFile.ID = model.ID.Hex()
		searchedFile.SearchTags = model.SearchTags
		searchedFile.Name = model.Name

		files = append(files, searchedFile)
	}

	return files, err
}

func (s *SearchStorageService) GetAllFilesOfUser(userID string) ([]search.SearchedFile, error) {
	ctx := context.Background()

	owner, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"owner": owner}

	cursor, err := s.coll.Find(ctx, filter)
	if err != nil {
		log.Error("Mongo error: %v", err)
		return nil, err
	}

	files := []search.SearchedFile{}

	for cursor.Next(ctx) {
		var model fileModel

		if err := cursor.Decode(&model); err != nil {
			log.Error("Error decoding: %v", err)
			return nil, err
		}

		var searchedFile search.SearchedFile

		searchedFile.ID = model.ID.Hex()
		searchedFile.SearchTags = model.SearchTags
		searchedFile.Name = model.Name

		files = append(files, searchedFile)
	}

	return files, err
}
