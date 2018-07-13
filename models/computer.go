package models

import (
	"fmt"

	bolt "github.com/coreos/bbolt"
	"github.com/lflxp/devops/pkg/db"
)

type Data struct {
	Time string
	Info string
}

type Test struct {
	Table string
	Key   string
	Value string
}

func (this *Test) AddTest() error {
	err := db.Db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(this.Table))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		err = b.Put([]byte(this.Key), []byte(this.Value))
		return err
	})
	return err
}

func (this *Test) GetTest() ([]*Test, error) {
	tmp := []*Test{}
	err := db.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(this.Table))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			t := &Test{Table: this.Table}
			t.Key = fmt.Sprintf("%s", k)
			t.Value = fmt.Sprintf("%s", v)
			tmp = append(tmp, t)
		}
		return nil
	})
	return tmp, err
}
