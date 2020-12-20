package storage

import (
	"cloud/pkg/upload"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//NewFileStorageService return an instance of FileStorageService
func NewFileStorageService(db *mongo.Database) *FileStorageService {
	return &FileStorageService{
		fileColl: db.Collection("files"),
	}
}

//FileStorageService is app storage for files build on badgerhold
type FileStorageService struct {
	fileColl *mongo.Collection
}

//UpdateOrCreate updates if file exists or creates new file
func (s *FileStorageService) UpdateOrCreate(fileInfo *upload.File) error {
	newFileModel := fileModel{}
	newFileModel.Name = fileInfo.Name
	newFileModel.Extension = fileInfo.Extension
	newFileModel.FullPath = fileInfo.FullPath
	newFileModel.UploadedAt = time.Now()
	newFileModel.ModifiedAt = time.Now()

	//create file if not exists
	existingFile, err := s.checkFilePath(&newFileModel)
	if err != nil {
		res, err := s.fileColl.InsertOne(context.Background(), newFileModel)
		if err != nil {
			return err
		}
		if _, ok := res.InsertedID.(primitive.ObjectID); !ok || ok {
			return ErrInvalidLastInsertedID
		}
		return nil
	}

	//update if do exists
	update := bson.M{"$set": bson.M{"modifiedAt": time.Now()}}

	_, err = s.fileColl.UpdateOne(context.Background(), bson.M{"_id": existingFile.ID}, update)
	if err != nil {
		return err
	}

	return nil
}

func (s *FileStorageService) checkFilePath(newFileModel *fileModel) (*fileModel, error) {
	var file fileModel
	filter := bson.M{"fullPath": newFileModel.FullPath}

	res := s.fileColl.FindOne(context.Background(), filter)
	if err := res.Decode(&file); err != nil {
		return nil, err
	}

	return &file, nil
}
