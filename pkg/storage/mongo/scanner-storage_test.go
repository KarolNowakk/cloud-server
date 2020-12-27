package storage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetUpdatedAfterShouldReturnAllFilesThatHasBeenUpdateAffterSpecifiedDate(t *testing.T) {
	db := dbMock()
	defer clearDatabase(db)

	scanner := NewScannerStorageService(db)

	user := getSampleInsertedUser(db)

	howManyFilesShouldCreate := 5

	for i := 0; i < howManyFilesShouldCreate; i++ {
		getSampleInsertedFile(db, user.ID)
	}

	time.Sleep(time.Second)

	files, err := scanner.GetUpdatedAfter(time.Now())

	require.Nil(t, err)
	require.Equal(t, len(files), howManyFilesShouldCreate)
}
