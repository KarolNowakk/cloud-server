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

	files, err := scanner.GetUpdatedAfter(time.Now().Add(-2*time.Minute), user.ID.Hex())

	require.Nil(t, err)
	require.Equal(t, howManyFilesShouldCreate, len(files))
}

func TestGetUpdatedAfterShouldReturnEmptySliceIfNothingShowedUpOnServerAferSpecifiedDate(t *testing.T) {
	db := dbMock()
	defer clearDatabase(db)

	scanner := NewScannerStorageService(db)

	user := getSampleInsertedUser(db)

	for i := 0; i < 5; i++ {
		getSampleInsertedFile(db, user.ID)
	}

	files, err := scanner.GetUpdatedAfter(time.Now().Add(2*time.Minute), user.ID.Hex())

	require.Nil(t, err)
	require.Equal(t, 0, len(files))
}

func TestGetUpdatedAfterShouldReturnZeroFilesIfUserRequestingFilesIsNotOwnerOfTheFiles(t *testing.T) {
	db := dbMock()
	defer clearDatabase(db)

	scanner := NewScannerStorageService(db)

	user1 := getSampleInsertedUser(db)
	user2 := getSampleInsertedUser(db)

	howManyFilesShouldCreate := 5

	for i := 0; i < howManyFilesShouldCreate; i++ {
		getSampleInsertedFile(db, user1.ID)
	}

	time.Sleep(time.Second)

	files, err := scanner.GetUpdatedAfter(time.Now(), user2.ID.Hex())

	require.Nil(t, err)
	require.Equal(t, len(files), 0)
}
