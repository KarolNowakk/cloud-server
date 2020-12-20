package upload

import (
	"cloud/pkg/upload/uploadpb"
	"log"
	"os"
)

//Service provides fiel uploading operations
type Service interface {
	CreateFileIfNotExistsAndOpen(file *uploadpb.FileUploadInfo) error
	WriteBytes(file *uploadpb.FileUploadBody) error
	UpdateOrCreateFile() error
}

//Repository is interface that plugged in repo service must satisfy
type Repository interface {
	UpdateOrCreate(fileInfo *File) error
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
func (s *service) CreateFileIfNotExistsAndOpen(file *uploadpb.FileUploadInfo) error {
	s.fileData = &File{
		Name:      file.GetName(),
		FullPath:  file.GetPath() + "/" + file.GetName() + "." + file.GetExtension(),
		Extension: file.GetExtension(),
	}

	if err := os.MkdirAll("files/"+file.GetPath(), os.ModePerm); err != nil {
		return err
	}

	filePath := s.fileData.FullPath

	writeFile, err := os.OpenFile("files/"+filePath, os.O_CREATE|os.O_WRONLY, 0777)
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
func (s *service) UpdateOrCreateFile() error {
	if err := s.r.UpdateOrCreate(s.fileData); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
