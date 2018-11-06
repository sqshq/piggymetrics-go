package account

import (
	"encoding/json"
	"errors"
	"github.com/labstack/gommon/log"
	"github.com/sqshq/piggymetrics/internal/model/user"
	"github.com/sqshq/piggymetrics/internal/store"
	"go.etcd.io/bbolt"
	"time"
)

type Account struct {
	Name     string    `json:"name"`
	LastSeen time.Time `json:"lastSeen"`
	Incomes  []Item    `json:"incomes"`
	Expenses []Item    `json:"expenses"`
	Saving   Saving    `json:"saving"`
	Note     string    `json:"note"`
}

type Item struct {
	Title      string     `json:"title"`
	Amount     string     `json:"amount"`
	Currency   Currency   `json:"currency"`
	TimePeriod TimePeriod `json:"period"`
	Icon       string     `json:"icon"`
}

type Saving struct {
	Amount         int64    `json:"amount"`
	Currency       Currency `json:"currency"`
	Interest       float64  `json:"interest"`
	Deposit        bool     `json:"deposit"`
	Capitalization bool     `json:"capitalization"`
}

type Currency string

type TimePeriod string

const (
	USD Currency = "USD"
	EUR Currency = "EUR"
	RUB Currency = "RUB"
)

const (
	YEAR    TimePeriod = "YEAR"
	QUARTER TimePeriod = "QUARTER"
	MONTH   TimePeriod = "MONTH"
	DAY     TimePeriod = "DAY"
	HOUR    TimePeriod = "HOUR"
)

func Create(str *store.Store, u *user.User) *Account {

	svg := Saving{Amount: 0, Currency: USD, Interest: 0, Deposit: false, Capitalization: false}
	acc := Account{Name: u.Username, LastSeen: time.Now(), Saving: svg}

	err := str.Db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(store.AccountBucket))

		// check for duplicates
		if b.Get([]byte(acc.Name)) != nil {
			return errors.New("Account already exists: " + acc.Name) // TODO return err
		}

		// serialize
		encoded, err := json.Marshal(acc)
		if err != nil {
			return err
		}

		// save
		if e := b.Put([]byte(acc.Name), encoded); e != nil {
			return errors.New("Failed to save account in the store: " + acc.Name) // TODO return err
		}

		return nil
	})

	if err != nil {
		log.Error("Error during account creation: ", err) // TODO return err
	}

	return &acc
}

func Update(str *store.Store, acc *Account) error {

	acc.LastSeen = time.Now()

	err := str.Db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(store.AccountBucket))

		// serialize
		encoded, err := json.Marshal(acc)
		if err != nil {
			return err
		}

		// save
		if e := b.Put([]byte(acc.Name), encoded); e != nil {
			return errors.New("Failed to update account in the store: " + acc.Name)
		}

		return nil
	})

	return err
}

func FindByName(str *store.Store, name string) *Account {

	acc := new(Account)

	str.Db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(store.AccountBucket))
		encoded := b.Get([]byte(name))

		if err := json.Unmarshal(encoded, &acc); err != nil {
			return errors.New("Can't find account by name: " + name)
		}

		return nil
	})

	return acc
}
