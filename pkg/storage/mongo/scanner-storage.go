package storage

import (
	"cloud-server/pkg/scanner"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//NewScannerStorageService return an instance of ScannerStorageService
func NewScannerStorageService(db *mongo.Database) *ScannerStorageService {
	return &ScannerStorageService{
		fileColl: db.Collection("files"),
	}
}

//ScannerStorageService scans database and looks for updates
type ScannerStorageService struct {
	fileColl *mongo.Collection
}

//GetUpdatedAfter returns files info of updated files after specified date
func (s *ScannerStorageService) GetUpdatedAfter(date time.Time) ([]scanner.ScannedFile, error) {
	filter := bson.M{"updatedAt": bson.M{"$lt": date}}

	cursor, err := s.fileColl.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var files []scanner.ScannedFile

	for cursor.Next(context.Background()) {
		var model fileModel

		if err := cursor.Decode(&model); err != nil {
			return nil, err
		}

		var scannedFile scanner.ScannedFile

		scannedFile.ID = model.ID.Hex()
		scannedFile.Path = model.FullPath
		scannedFile.Owner = model.Owner.Hex()
		scannedFile.Name = model.Name
		scannedFile.Extension = model.Extension

		files = append(files, scannedFile)
	}

	return files, nil
}
