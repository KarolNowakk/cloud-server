package upload

import (
	"cloud/pkg/config"
	"cloud/pkg/upload/uploadpb"
	"context"
	"log"
	"os"
)

//Service provides fiel uploading operations
type Service interface {
	CreateFileIfNotExistsAndOpen(ctx context.Context, file *uploadpb.FileUploadInfo, userID string) error
	WriteBytes(file *uploadpb.FileUploadBody) error
	UpdateOrCreateFile(ctx context.Context, userID string) error
}

//Repository is interface that plugged in repo service must satisfy
type Repository interface {
	UpdateOrCreate(ctx context.Context, fileInfo *File, userID string) error
}

//NewService returns new upload handler instance
func NewService(r Repository) Service {
	return &service{r: r}
}

//Handler handles file upload
type service struct {
	r        Repository
	file     *os.File
	fileData *File
}

//CreateFileIfNotExistsAndOpen creates file and all directories
func (s *service) CreateFileIfNotExistsAndOpen(ctx context.Context, file *uploadpb.FileUploadInfo, userID string) error {
	s.fileData = &File{
		Name:       file.GetName(),
		Owner:      userID,
		SearchTags: file.GetSearchTags(),
	}

	fullPath := config.UploadFolder + "/" + s.fileData.Owner + "-" + file.Name

	s.fileData.Path = fullPath

	writeFile, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}

	s.file = writeFile

	return nil
}

//WriteBytes writes bytes to a file
func (s *service) WriteBytes(file *uploadpb.FileUploadBody) error {
	_, err := s.file.Write(file.GetBytes())
	if err != nil {
		return err
	}

	return nil
}

//UpdateOrCreateFile creates or updates file
func (s *service) UpdateOrCreateFile(ctx context.Context, userID string) error {
	if err := s.r.UpdateOrCreate(ctx, s.fileData, userID); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
