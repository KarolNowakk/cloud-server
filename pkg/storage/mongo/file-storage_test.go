package storage

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func TestInsertFolderShouldGoOk(t *testing.T) {
	coll := dbMock().Collection("folders")
	defer clearCollection(coll)

	folder := &folderModel{
		Name:     "test",
		FullPath: "testing/testosteron/test",
	}

	insertFolder(coll, folder)

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

func TestUpdateFolderShouldGoOk(t *testing.T) {
	coll := dbMock().Collection("folders")
	defer clearCollection(coll)

	folder := folderModel{
		Name:     "test",
		FullPath: "testing/testosteron/test",
	}

	insertFolder(coll, &folder)
	_ = coll.FindOne(context.Background(), bson.M{"fullPath": folder.FullPath}).Decode(&folder)

	time.Sleep(1 * time.Second)

	updateFolder(coll, &folder)
	_ = coll.FindOne(context.Background(), bson.M{"fullPath": folder.FullPath}).Decode(&folder)

	require.False(t, folder.CreatedAt.Equal(folder.ModifiedAt))
}
