package main

import (
	"encoding/binary"

	"github.com/boltdb/bolt"
)

const (
	keyToken             = "currentToken"
	keyUserID            = "userID"
	keyTokenExpiringTime = "tokenExpiry"
	keyLastServerCheck   = "lastServerCheck"
	keyMainDirPath       = "mainDirPath"
)

func saveToken(db *bolt.DB, token string, userID string, expTime int64) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("config"))

		if err := b.Put([]byte(keyToken), []byte(token)); err != nil {
			return err
		}

		if err := b.Put([]byte(keyUserID), []byte(userID)); err != nil {
			return err
		}

		bytesTime := make([]byte, 8)
		binary.LittleEndian.PutUint64(bytesTime, uint64(expTime))

		if err := b.Put([]byte(keyTokenExpiringTime), bytesTime); err != nil {
			return err
		}

		return nil
	})
}

func findConfigString(db *bolt.DB, key string) (string, error) {
	var value string

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("config"))
		v := b.Get([]byte(key))
		value = string(v)

		return nil
	})

	if err != nil {
		return "", err
	}

	return value, nil
}
