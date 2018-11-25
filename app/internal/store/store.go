package store

import (
	"github.com/labstack/gommon/log"
	"github.com/sqshq/piggymetrics-go/app/config"
	"go.etcd.io/bbolt"
	"time"
)

const AccountBucket = "account"
const UserBucket = "user"

type Store struct {
	configuration *config.Configuration
	Db            *bbolt.DB
}

func New(c *config.Configuration) *Store {

	db, err := bbolt.Open("app/db/piggymetrics.db", 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	topBuckets := []string{AccountBucket, UserBucket}

	err = db.Update(func(tx *bbolt.Tx) error {
		for _, bktName := range topBuckets {
			if _, e := tx.CreateBucketIfNotExists([]byte(bktName)); e != nil {
				return e
			}
		}
		return nil
	})

	if err != nil {
		panic("Failed to create top level buckets)")
	}

	return &Store{configuration: c, Db: db}
}
