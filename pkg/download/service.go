package download

import (
	"cloud/pkg/download/downloadpb"
	"os"
)

//Service provides fiel uploading operations
type Service interface {
	ReadBytes() ([]byte, error)
	// RecordDownloadFile(userID string)
	OpenFile(req *downloadpb.FileDownloadRequest, userID string) error
}

//Repository is interface that plugged in repo service must satisfy
type Repository interface {
	// GetFileInfo(fileInfo *FileDownload) (*FileDownload, error)
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
func (s *service) OpenFile(req *downloadpb.FileDownloadRequest, userID string) error {
	fileInfo := &FileDownload{
		Name:               req.GetName(),
		Extension:          req.GetExtension(),
		Path:               req.GetPath(),
		FromPersonalFolder: req.GetFromPersonalFolder(),
	}

	var fullPath string

	if req.GetFromPersonalFolder() {
		fileInfo.Owner = userID
	} else {
		fileInfo.Owner = req.GetExternalFolderId()
	}

	if fileInfo.FromPersonalFolder {
		fullPath = userID + "/" + fileInfo.Path
	} else {
		fullPath = fileInfo.Owner + "/" + fileInfo.Path
	}

	fullPath = "files/" + fullPath + "/" + fileInfo.Name + "." + fileInfo.Extension

	readFile, err := os.OpenFile(fullPath, os.O_RDONLY, 0644)
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

// //RecordDownloadFile records file download
// func (s *service) RecordDownloadFile(userID string) {
// 	_ = s.r.UpdateAsDownloaded(s.fileData, userID)
// }
