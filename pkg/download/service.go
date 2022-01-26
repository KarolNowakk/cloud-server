package download

import (
	"context"
	"os"
)

//Service provides fiel uploading operations
type Service interface {
	ReadChunk(file *os.File, chunkSize int, off int64, whence int) ([]byte, error)
	OpenFile(path string) (*os.File, error)
	FindFile(ctx context.Context, fileID string) (FileDownload, error)
	DeleteFile(ctx context.Context, fileID string) error
}

//Repository is interface that plugged in repo service must satisfy
type Repository interface {
	FindFile(ctx context.Context, fileID string) (FileDownload, error)
	DeleteFile(ctx context.Context, fileID string) error
}

//NewService returns new upload handler instance
func NewService(r Repository) Service {
	return &service{r: r}
}

//Handler handles file upload
type service struct {
	r Repository
}

func (s service) FindFile(ctx context.Context, fileID string) (FileDownload, error) {
	return s.r.FindFile(ctx, fileID)
}

func (s service) OpenFile(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_RDONLY, 0644)
}

func (service) ReadChunk(file *os.File, chunkSize int, off int64, whence int) ([]byte, error) {
	bytes := make([]byte, chunkSize)

	file.Seek(off, whence)
	_, err := file.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (s service) DeleteFile(ctx context.Context, fileID string) error {
	return s.r.DeleteFile(ctx, fileID)
}
