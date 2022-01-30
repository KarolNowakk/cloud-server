package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/boltdb/bolt"
	"github.com/stretchr/testify/require"
)

func mockTestDB() *bolt.DB {
	os.Mkdir("mockDB", 0777)

	db, err := bolt.Open("./mockDB/bolt", 0777, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	// Start writable transaction.
	tx, err := db.Begin(true)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	if _, err := tx.CreateBucketIfNotExists([]byte("config")); err != nil {
		log.Fatal(err)
	}
	if _, err := tx.CreateBucketIfNotExists([]byte("files")); err != nil {
		log.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	return db
}

func clearMockDB() {
	os.RemoveAll("./mockDB/bolt")
}

func getFileMock() *fileModel {
	return &fileModel{
		Name:         "testName",
		FullPath:     "test/path/testName.exe",
		OnServerID:   "testServerID",
		DownloadFlag: 0,
		UploadFlag:   4,
	}
}

func TestInsertOrUpdateShouldInsertFiles(t *testing.T) {
	db := mockTestDB()
	defer clearMockDB()

	file := getFileMock()

	fileStorage := fileStorage{db: db}

	err := fileStorage.insertOrUpdate(file, 0)

	require.Nil(t, err)

	i := 0

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("files"))
		c := b.Cursor()

		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			i++
		}

		return nil
	})

	require.Equal(t, 1, i)
}

func TestInsertOrUpdateShouldReturnErrWhenFullPathAllreadyExists(t *testing.T) {
	db := mockTestDB()
	defer clearMockDB()

	file := getFileMock()

	fileStorage := fileStorage{db: db}

	fileStorage.insertOrUpdate(file, 0)
	file.OnServerID = "differentServerID"

	err := fileStorage.insertOrUpdate(file, 0)

	require.IsType(t, errFileAllreadyExists, err)
}

func TestInsertOrUpdateShouldReturnErrWhenOnServerIDAllreadyExists(t *testing.T) {
	db := mockTestDB()
	defer clearMockDB()

	file := getFileMock()

	fileStorage := fileStorage{db: db}

	fileStorage.insertOrUpdate(file, 0)
	file.FullPath = "different/test/path"

	err := fileStorage.insertOrUpdate(file, 0)

	require.IsType(t, errFileAllreadyExists, err)
}

func TestInsertOrUpdateShouldUpdateFiles(t *testing.T) {
	db := mockTestDB()
	defer clearMockDB()

	file := getFileMock()

	fileStorage := fileStorage{db: db}

	fileStorage.insertOrUpdate(file, 0)

	file.Name = "testUpdate"
	file.FullPath = "different/full/path"
	file.OnServerID = "differentOnServerId"

	err := fileStorage.insertOrUpdate(file, file.ID)

	require.Nil(t, err)

	model := &fileModel{}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("files"))

		v := b.Get(itob(file.ID))
		fmt.Println(v, file.ID)
		json.Unmarshal(v, model)

		return nil
	})

	require.Equal(t, "testUpdate", model.Name)
}

func TestGetByIDShouldReturnFileModelAssignedToGivenID(t *testing.T) {
	db := mockTestDB()
	defer clearMockDB()

	file := getFileMock()

	fileStorage := fileStorage{db: db}

	fileStorage.insertOrUpdate(file, 0)

	retrivedFile, err := fileStorage.getByID(file.ID)

	require.Nil(t, err)

	require.Equal(t, "testName", retrivedFile.Name)
}

func TestFindOneByStringShouldReturnFileModelMatchingGivenFiledAndValue(t *testing.T) {
	db := mockTestDB()
	defer clearMockDB()

	file := getFileMock()

	fileStorage := fileStorage{db: db}

	fileStorage.insertOrUpdate(file, 0)

	retrivedFile, err := fileStorage.findOneByString("fullPath", file.FullPath)

	require.Nil(t, err)
	require.Equal(t, file.OnServerID, retrivedFile.OnServerID)
}

func TestFindOneByStringShouldReturnErrFileNotFoundIfFileNotFound(t *testing.T) {
	db := mockTestDB()
	defer clearMockDB()

	file := getFileMock()

	fileStorage := fileStorage{db: db}

	fileStorage.insertOrUpdate(file, 0)

	_, err := fileStorage.findOneByString("fullPath", "adadfadcd")

	require.IsType(t, err, errFileNotFound)
}

func TestFindManyByIntShouldReturnAsManyFilesAsMatchesProvidedCallback(t *testing.T) {
	db := mockTestDB()
	defer clearMockDB()

	fileStorage := fileStorage{db: db}

	for i := 0; i < 10; i++ {
		file := getFileMock()

		file.FullPath += fmt.Sprint(i)
		file.OnServerID += fmt.Sprint(i)
		file.DownloadFlag = flagUninit

		fileStorage.insertOrUpdate(file, 0)
	}

	files, err := fileStorage.findManyByInt("downloadFlag", func(retrived int) bool {
		return retrived <= flagUninit
	})

	require.Nil(t, err)
	require.Equal(t, 10, len(files))
}
