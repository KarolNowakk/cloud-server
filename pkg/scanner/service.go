package scanner

import (
	"time"
)

//Service provides fiel uploading operations
type Service interface {
	Scan(date time.Time, userID string) ([]ScannedFile, error)
}

//Repository is interface that plugged in repo service must satisfy
type Repository interface {
	GetUpdatedAfter(date time.Time, userID string) ([]ScannedFile, error)
}

//NewService returns new upload handler instance
func NewService(r Repository) Service {
	return &service{r: r}
}

//Handler handles file upload
type service struct {
	r Repository
}

func (s *service) Scan(date time.Time, userID string) ([]ScannedFile, error) {
	return s.r.GetUpdatedAfter(date, userID)
}
