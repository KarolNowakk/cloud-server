package download

import (
	"os"
)

//Service provides fiel uploading operations
type Service interface {
	ReadBytes() ([]byte, error)
	RecordDownloadFile() error
	OpenFile(fileInfo *FileDownload) error
}

//Repository is interface that plugged in repo service must satisfy
type Repository interface {
	CanDownload() error
}

//NewService returns new upload handler instance
func NewService(r Repository) Service {
	return &service{r: r}
}

//Handler handles file upload
type service struct {
	r        Repository
	file     *os.File
	fileData *FileDownload
}

//CreateFileIfNotExistsAndOpen creates file and all directories
func (s *service) OpenFile(fileInfo *FileDownload) error {
	fullPath := fileInfo.Path + "/" + fileInfo.Name + "." + fileInfo.Extension

	readFile, err := os.OpenFile("files/"+fullPath, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}

	s.fileData = fileInfo
	s.file = readFile

	return nil
}

//ReadBytes read bytes from a file
func (s *service) ReadBytes() ([]byte, error) {
	bytes := make([]byte, 4*1024)

	_, err := s.file.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

//RecordDownloadFile records file download
func (s *service) RecordDownloadFile() error {
	return nil
}
