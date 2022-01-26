package storage

import (
	"cloud/pkg/download"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//NewDownloadStorageService return an instance of FileStorageService
func NewDownloadStorageService(db *mongo.Database) *DownloadStorageService {
	return &DownloadStorageService{
		coll: db.Collection("files"),
	}
}

//DownloadStorageService is app storage for files build on badgerhold
type DownloadStorageService struct {
	coll *mongo.Collection
}

func (s DownloadStorageService) FindFile(ctx context.Context, fileID string) (download.FileDownload, error) {
	objectID, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		return download.FileDownload{}, err
	}

	var file fileModel
	filter := bson.M{"_id": objectID}

	if err := s.coll.FindOne(ctx, filter).Decode(&file); err != nil {
		return download.FileDownload{}, err
	}

	return download.FileDownload{
		ID:    file.ID.Hex(),
		Owner: file.Owner.Hex(),
		Path:  file.Path,
		Name:  file.Name,
	}, nil
}

func (s DownloadStorageService) DeleteFile(ctx context.Context, fileID string) error {
	objectID, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	if _, err := s.coll.DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}
