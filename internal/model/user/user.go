package user

import (
	"errors"
	"github.com/labstack/gommon/log"
	"github.com/sqshq/piggymetrics/internal/store"
	"go.etcd.io/bbolt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Create(str *store.Store, u *User) *User {

	err := str.Db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(store.UserBucket))

		// check for duplicates
		if b.Get([]byte(u.Username)) != nil {
			return errors.New("User already exists: " + u.Username)
		}

		// hash given password
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return errors.New("Failed to generate hash for user: " + u.Username) // TODO return err
		}

		// save
		if e := b.Put([]byte(u.Username), hash); e != nil {
			return errors.New("Failed to save user in the store: " + u.Username) // TODO return err
		}

		return err
	})

	if err != nil {
		log.Error("Error during account creation: ", err) // TODO return err
	}

	return u
}

func Authenticate(str *store.Store, u *User) bool {

	var hash []byte

	err := str.Db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(store.UserBucket))
		hash = b.Get([]byte(u.Username))
		return nil
	})

	if err != nil {
		log.Error("Error during authentication: ", err)
		return false
	}

	err = bcrypt.CompareHashAndPassword(hash, []byte(u.Password))

	if err != nil {
		return false
	}

	return true
}
