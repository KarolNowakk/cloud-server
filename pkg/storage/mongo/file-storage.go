package storage

import (
	"cloud/pkg/upload"
	"context"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var insertFolder = func(coll *mongo.Collection, newFolderModel *folderModel) {
	res, err := coll.InsertOne(context.Background(), newFolderModel)
	if err != nil {
		log.Println("error while saving folder model: ", err)
		return
	}
	if _, ok := res.InsertedID.(primitive.ObjectID); !ok {
		log.Println("error while saving folder model: ", err)
		return
	}
}

var updateFolder = func(coll *mongo.Collection, newFolderModel *folderModel) {
	update := bson.M{"$set": bson.M{"modifiedAt": time.Now()}}
	filter := bson.M{"_id": newFolderModel.ID}

	_, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("error while updating folder model: ", err)
		return
	}
}

//NewFileStorageService return an instance of FileStorageService
func NewFileStorageService(db *mongo.Database) *FileStorageService {
	return &FileStorageService{
		fileColl:   db.Collection("files"),
		folderColl: db.Collection("folders"),
	}
}

//FileStorageService is app storage for files build on mongo
type FileStorageService struct {
	fileColl   *mongo.Collection
	folderColl *mongo.Collection
}

//UpdateOrCreate updates if file exists or creates new file
func (s *FileStorageService) UpdateOrCreate(fileInfo *upload.File, userID string) error {
	belongsTo, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	var newFileModel fileModel
	mapServiceFileToFileModel(fileInfo, &newFileModel)

	newFileModel.BelongsTo = belongsTo

	//create file if not exists
	existingFile, err := s.checkFilePath(newFileModel.FullPath, userID, newFileModel.BelongsTo)
	if err != nil {
		res, err := s.fileColl.InsertOne(context.Background(), newFileModel)
		if err != nil {
			return err
		}
		if _, ok := res.InsertedID.(primitive.ObjectID); !ok {
			return ErrInvalidLastInsertedID
		}

		go s.doFolderAction(&newFileModel, userID, insertFolder)

		return nil
	}

	//update if do exists
	update := bson.M{"$set": bson.M{"modifiedAt": time.Now()}}

	_, err = s.fileColl.UpdateOne(context.Background(), bson.M{"_id": existingFile.ID}, update)
	if err != nil {
		return err
	}

	go s.doFolderAction(&newFileModel, userID, updateFolder)

	return nil
}

func (s *FileStorageService) checkFilePath(fullPath, userID string, belongsTo primitive.ObjectID) (*fileModel, error) {
	var file fileModel
	filter := bson.M{"fullPath": fullPath, "belongsTo": belongsTo}

	res := s.fileColl.FindOne(context.Background(), filter)
	if err := res.Decode(&file); err != nil {
		return nil, err
	}

	return &file, nil
}

func (s *FileStorageService) doFolderAction(
	newFileModel *fileModel,
	userID string,
	action func(coll *mongo.Collection, newFolderModel *folderModel)) {
	pathSliceNotFiltered := strings.Split(newFileModel.FullPath, "/")
	if len(pathSliceNotFiltered) < 1 {
		log.Println("path folders has not been saved")
		return
	}

	var pathSlice []string

	for _, name := range pathSliceNotFiltered {
		if name == "." || strings.Contains(name, ".") {
			continue
		}

		pathSlice = append(pathSlice, name)

		newFolderModel := folderModel{
			Name:      name,
			FullPath:  strings.Join(pathSlice, "/"),
			BelongsTo: newFileModel.BelongsTo,
		}

		action(s.folderColl, &newFolderModel)
	}
}

func mapServiceFileToFileModel(fileInfo *upload.File, newFileModel *fileModel) {
	newFileModel.Name = fileInfo.Name
	newFileModel.Extension = fileInfo.Extension
	newFileModel.FullPath = fileInfo.FullPath
	newFileModel.UploadedAt = time.Now()
	newFileModel.ModifiedAt = time.Now()
}
