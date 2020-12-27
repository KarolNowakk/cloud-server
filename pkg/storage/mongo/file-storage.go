package storage

import (
	"cloud/pkg/upload"
	"context"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var insert = func(coll *mongo.Collection, newModel auditCreateUpdateInterface, userID primitive.ObjectID) error {
	newModel.Creating(userID)

	res, err := coll.InsertOne(context.Background(), newModel)
	if err != nil {
		return err
	}
	if _, ok := res.InsertedID.(primitive.ObjectID); !ok {
		return ErrInvalidLastInsertedID
	}

	return nil
}

var update = func(coll *mongo.Collection, model auditCreateUpdateInterface, userID primitive.ObjectID) error {
	filter := bson.M{"_id": model.GetUpdatingModelID()}

	model.Updating(userID)

	_, err := coll.ReplaceOne(context.Background(), filter, model)
	if err != nil {
		return err
	}

	return nil
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
	currentUserID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	var newFileModel fileModel
	if err := mapServiceFileToFileModel(fileInfo, &newFileModel); err != nil {
		return err
	}

	//create file if not exists
	err = s.checkFilePath(s.fileColl, &newFileModel, newFileModel.FullPath, newFileModel.Owner)
	if err != nil {
		if err := insert(s.fileColl, &newFileModel, currentUserID); err != nil {
			return err
		}
	} else {
		if err := update(s.fileColl, &newFileModel, currentUserID); err != nil {
			return err
		}
	}

	s.doFolderAction(&newFileModel, currentUserID)

	return nil
}

func (s *FileStorageService) checkFilePath(coll *mongo.Collection, model auditCreateUpdateInterface, fullPath string, belongsTo primitive.ObjectID) error {
	filter := bson.M{"fullPath": fullPath, "owner": belongsTo}

	res := coll.FindOne(context.Background(), filter)
	if err := res.Decode(model); err != nil {
		return err
	}

	return nil
}

func (s *FileStorageService) doFolderAction(newFileModel *fileModel, userID primitive.ObjectID) {
	pathSliceNotFiltered := strings.Split(newFileModel.FullPath, "/")
	if len(pathSliceNotFiltered) < 1 {
		log.Println("path folders has not been saved")
		return
	}

	var pathSlice []string

	for _, name := range pathSliceNotFiltered {
		if strings.Contains(name, ".") || len(name) < 1 {
			continue
		}

		pathSlice = append(pathSlice, name)

		newFolderModel := folderModel{
			Name:     name,
			FullPath: strings.Join(pathSlice, "/"),
			Owner:    newFileModel.Owner,
		}

		err := s.checkFilePath(s.folderColl, &newFolderModel, newFolderModel.FullPath, newFolderModel.Owner)
		if err != nil {
			if err := insert(s.folderColl, &newFolderModel, userID); err != nil {
				log.Println(err)
				return
			}
		} else {
			if err := update(s.folderColl, &newFolderModel, userID); err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func mapServiceFileToFileModel(fileInfo *upload.File, newFileModel *fileModel) error {
	owner, err := primitive.ObjectIDFromHex(fileInfo.Owner)
	if err != nil {
		return err
	}

	newFileModel.Name = fileInfo.Name
	newFileModel.Extension = fileInfo.Extension
	newFileModel.FullPath = fileInfo.FullPath
	newFileModel.Size = fileInfo.Size
	newFileModel.Owner = owner

	return nil
}
