package main

import (
	"encoding/json"

	"github.com/boltdb/bolt"
)

const (
	flagUninit  = 0 //flag when download/upload has not been initialized yet
	flagFail    = 1 //flag when download/upload failed
	flagInit    = 2 //flag when download/upload was initialized
	flagSuccess = 3 //flag when download/upload was successful
	flagNone    = 4 //flag when flag is not uploded/downloaed but downloaded/uploaded
)

type fileModel struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	Extension    string `json:"extension"`
	OnServerID   string `json:"onServerId"`
	FullPath     string `json:"fullPath"`
	ParentFolder string `json:"parentFolder"`
	Owner        string `json:"owner"`
	Size         int64  `json:"size"`
	UploadFlag   int    `json:"uploadFlag"`
	DownloadFlag int    `json:"downloadFlag"`
}

type fileStorage struct{ db *bolt.DB }

// inserts if id == 0, updates if id != 0
func (s *fileStorage) insertOrUpdate(file *fileModel, id uint64) error {
	if id == 0 {
		if _, err := s.findOneByString("fullPath", file.FullPath); err == nil {
			return errFileAllreadyExists
		}

		if _, err := s.findOneByString("onServerId", file.OnServerID); err == nil {
			return errFileAllreadyExists
		}
	}

	return s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("files"))

		if id == 0 {
			id, _ = bucket.NextSequence()
		}

		file.ID = id

		fileJSON, err := json.Marshal(file)

		if err != nil {
			return err
		}

		if err := bucket.Put(itob(id), []byte(fileJSON)); err != nil {
			return err
		}

		return nil
	})
}

func (s *fileStorage) getByID(id uint64) (*fileModel, error) {
	var model fileModel

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("files"))

		v := b.Get(itob(id))

		if err := json.Unmarshal(v, &model); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (s *fileStorage) findOneByString(filed, value string) (*fileModel, error) {
	var model fileModel

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("files"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var obj map[string]json.RawMessage
			if err := json.Unmarshal(v, &obj); err != nil {
				return err
			}

			var retrivedValue string
			if err := json.Unmarshal(obj[filed], &retrivedValue); err != nil {
				return err
			}

			if retrivedValue == value {
				if err := json.Unmarshal(v, &model); err != nil {
					return err
				}
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	//all files in database must have at least FullPath
	if model.FullPath == "" {
		return nil, errFileNotFound
	}

	return &model, nil
}

func (s *fileStorage) findManyByInt(filed string, compare func(retrivedValue int) bool) ([]fileModel, error) {
	var models []fileModel

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("files"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var obj map[string]json.RawMessage
			if err := json.Unmarshal(v, &obj); err != nil {
				return err
			}

			var retrivedValue int
			if err := json.Unmarshal(obj[filed], &retrivedValue); err != nil {
				return err
			}

			if compare(retrivedValue) {
				var model fileModel
				if err := json.Unmarshal(v, &model); err != nil {
					return err
				}

				models = append(models, model)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	if len(models) < 1 {
		return nil, errFileNotFound
	}

	return models, nil
}
