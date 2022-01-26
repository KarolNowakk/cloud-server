package search

//Service provides fiel uploading operations
type Service interface {
	Search(phrase, userID string) ([]SearchedFile, error)
}

//Repository is interface that plugged in repo service must satisfy
type Repository interface {
	SearchByPhrase(phrase, userID string) ([]SearchedFile, error)
	GetAllFilesOfUser(userID string) ([]SearchedFile, error)
}

//NewService returns new upload handler instance
func NewService(r Repository) Service {
	return &service{r: r}
}

//Handler handles file upload
type service struct {
	r Repository
}

func (s *service) Search(phrase, userID string) ([]SearchedFile, error) {
	if len(phrase) > 3 {
		return s.r.GetAllFilesOfUser(userID)
	}

	return s.r.SearchByPhrase(phrase, userID)
}
