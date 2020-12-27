package scanner

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

type repoMock struct {
	files []ScannedFile
}

func (s *repoMock) insertManyFiles(userID string, howMany int) {
	for i := 0; i < howMany; i++ {
		file := ScannedFile{
			ID:        randomHexID(),
			Path:      "path0/path1/path2",
			Owner:     userID,
			Name:      "file",
			Extension: "pdf",
		}

		s.files = append(s.files, file)
	}
}

func (s *repoMock) GetUpdatedAfter(date time.Time, userID string) ([]ScannedFile, error) {
	files := []ScannedFile{}

	for _, file := range s.files {
		if userID == file.Owner {
			files = append(files, file)
		}
	}

	return files, nil
}

func randomHexID() string {
	bytes := make([]byte, 16)
	_, _ = rand.Read(bytes)

	return hex.EncodeToString(bytes)
}
