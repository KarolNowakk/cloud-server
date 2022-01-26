package storage

import (
	"cloud/pkg/upload"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//NewFileStorageService return an instance of FileStorageService
func NewFileStorageService(db *mongo.Database) *FileStorageService {
	return &FileStorageService{
		coll: db.Collection("files"),
	}
}

//FileStorageService is app storage for files build on mongo
type FileStorageService struct {
	coll *mongo.Collection
}

//UpdateOrCreate updates if file exists or creates new file
func (s *FileStorageService) UpdateOrCreate(ctx context.Context, fileInfo *upload.File, userID string) error {
	currentUserID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	newFileModel := fileModel{
		Name:       fileInfo.Name,
		Path:       fileInfo.Path,
		SearchTags: fileInfo.SearchTags,
		Owner:      currentUserID,
	}

	err = s.checkFilePath(ctx, newFileModel.Path, newFileModel.Owner)

	if err != nil {
		s.insert(ctx, newFileModel, currentUserID)
	} else {
		s.update(ctx, newFileModel, currentUserID)
	}

	return nil
}

func (s *FileStorageService) checkFilePath(ctx context.Context, path string, belongsTo primitive.ObjectID) error {
	filter := bson.M{"path": path, "owner": belongsTo}

	model := &fileModel{}

	res := s.coll.FindOne(ctx, filter)
	if err := res.Decode(model); err != nil {
		return err
	}

	return nil
}

func (s *FileStorageService) insert(ctx context.Context, newModel auditCreateUpdateInterface, userID primitive.ObjectID) error {
	newModel.Creating(userID)

	res, err := s.coll.InsertOne(ctx, newModel)
	if err != nil {
		return err
	}
	if _, ok := res.InsertedID.(primitive.ObjectID); !ok {
		return ErrInvalidLastInsertedID
	}

	return nil
}

func (s *FileStorageService) update(ctx context.Context, model auditCreateUpdateInterface, userID primitive.ObjectID) error {
	filter := bson.M{"_id": model.GetUpdatingModelID()}

	model.Updating(userID)

	_, err := s.coll.ReplaceOne(ctx, filter, model)
	if err != nil {
		return err
	}

	return nil
}
