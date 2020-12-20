package storage

import "go.mongodb.org/mongo-driver/mongo"

//NewDownloadStorageService return an instance of FileStorageService
func NewDownloadStorageService(db *mongo.Database) *DownloadStorageService {
	return &DownloadStorageService{
		fileColl: db.Collection("files"),
	}
}

//DownloadStorageService is app storage for files build on badgerhold
type DownloadStorageService struct {
	fileColl *mongo.Collection
}

//CanDownload tells if user can download specific file
func (s *DownloadStorageService) CanDownload() error {
	return nil
}
