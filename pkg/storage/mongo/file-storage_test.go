package storage

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func TestInsertShouldGoOk(t *testing.T) {
	coll := dbMock().Collection("folders")
	defer clearDatabase(coll.Database())

	folder := &folderModel{
		Name:     "test",
		FullPath: "testing/testosteron/test",
	}

	user := getSampleInsertedUser(coll.Database())

	insert(coll, folder, user.ID)

	cursor, err := coll.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	for cursor.Next(context.Background()) {
		i++
	}

	require.Equal(t, 1, i)
}

func TestUpdateShouldGoOk(t *testing.T) {
	coll := dbMock().Collection("folders")
	defer clearDatabase(coll.Database())

	folder := folderModel{
		Name:     "test",
		FullPath: "testing/testosteron/test",
	}

	user := getSampleInsertedUser(coll.Database())

	insert(coll, &folder, user.ID)

	_ = coll.FindOne(context.Background(), bson.M{"fullPath": folder.FullPath}).Decode(&folder)

	time.Sleep(1 * time.Second)

	update(coll, &folder, user.ID)
	_ = coll.FindOne(context.Background(), bson.M{"fullPath": folder.FullPath}).Decode(&folder)

	require.False(t, folder.CreatedAt.Equal(folder.UpdatedAt))
}

func TestDoFolderActionShouldInsertAsMAnyFolderAsIsFoldersInPathExcludingDotsAndFile(t *testing.T) {
	coll := dbMock().Collection("folders")
	defer clearDatabase(coll.Database())

	s := FileStorageService{folderColl: coll}

	user := getSampleInsertedUser(coll.Database())
	file := getSampleFile(coll.Database(), user.ID)
	file.FullPath = "testing/tester/teste.r/././././testosteron/file.pdf"

	s.doFolderAction(file, user.ID)

	cursor, _ := coll.Find(context.Background(), bson.M{})

	i := 0
	for cursor.Next(context.Background()) {
		i++
	}

	require.Equal(t, 3, i)
}

func TestDoFolderActionShouldUpdateFoldersIfTheyAllreadyExists(t *testing.T) {
	coll := dbMock().Collection("folders")
	defer clearDatabase(coll.Database())

	s := FileStorageService{folderColl: coll}

	user := getSampleInsertedUser(coll.Database())
	file := getSampleFile(coll.Database(), user.ID)
	file.FullPath = "testing/tester/teeest/file.pdf"

	s.doFolderAction(file, user.ID)

	cursor, _ := coll.Find(context.Background(), bson.M{})

	var inserted []folderModel
	for cursor.Next(context.Background()) {
		var model folderModel

		cursor.Decode(&model)

		inserted = append(inserted, model)
	}

	// time.Sleep(2 * time.Second)

	file.FullPath = "testing/tester/teeest/file.pdf"

	s.doFolderAction(file, user.ID)

	cursor, _ = coll.Find(context.Background(), bson.M{})

	var updated []folderModel
	for cursor.Next(context.Background()) {
		var model folderModel

		cursor.Decode(&model)

		updated = append(updated, model)
	}

	for i, value := range inserted {
		fmt.Println(value.UpdatedAt, updated[i].UpdatedAt)
		require.False(t, value.UpdatedAt.Equal(updated[i].UpdatedAt))
	}
}
func TestDoFolderActionShouldNotInsertAnyFolder(t *testing.T) {
	coll := dbMock().Collection("folders")
	defer clearDatabase(coll.Database())

	s := FileStorageService{folderColl: coll}

	user := getSampleInsertedUser(coll.Database())
	file := getSampleFile(coll.Database(), user.ID)
	file.FullPath = "/file.pdf"

	s.doFolderAction(file, user.ID)

	cursor, _ := coll.Find(context.Background(), bson.M{})

	i := 0
	for cursor.Next(context.Background()) {
		i++
	}

	require.Equal(t, 0, i)
}

func TestUpdateOrCreateShouldInserFileIfItDoNotAllreadyExists(t *testing.T) {
	filesColl := dbMock().Collection("files")
	folderColl := dbMock().Collection("folders")
	defer clearDatabase(filesColl.Database())

	s := FileStorageService{fileColl: filesColl, folderColl: folderColl}

	user := getSampleInsertedUser(filesColl.Database())
	file := getSampleFileAsUploadFile(filesColl.Database(), user.ID)

	s.UpdateOrCreate(file, user.ID.Hex())

	cursor, _ := filesColl.Find(context.Background(), bson.M{})

	i := 0
	for cursor.Next(context.Background()) {
		i++
	}

	require.Equal(t, 1, i)
}

func TestCheckFilePathShoulReturnErrorWhenFileDoesNotExistsOnDatabase(t *testing.T) {
	coll := dbMock().Collection("folders")
	defer clearDatabase(coll.Database())

	s := FileStorageService{folderColl: coll}

	user := getSampleInsertedUser(coll.Database())
	fullPath := "testing/files"

	var folder folderModel
	err := s.checkFilePath(coll, &folder, fullPath, user.ID)

	require.NotNil(t, err)
}

func TestCheckFilePathShoulReturnNilWhenFileDoExistsOnDatabase(t *testing.T) {
	coll := dbMock().Collection("folders")
	defer clearDatabase(coll.Database())

	s := FileStorageService{folderColl: coll}

	user := getSampleInsertedUser(coll.Database())

	folder := folderModel{
		Name:     "test",
		FullPath: "testing/testosteron/test",
		Owner:    user.ID,
	}

	insert(coll, &folder, user.ID)

	err := s.checkFilePath(coll, &folder, folder.FullPath, user.ID)

	require.Nil(t, err)
}
