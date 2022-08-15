package stores

import (
	"errors"
	"fmt"

	badger "github.com/dgraph-io/badger/v3"
)

type Item struct {
	UID         string
	Category    string
	Description string
	Price       float64
	Quantity    uint64
}

type ItemStore struct {
	DB *badger.DB
}

func (itemStore *ItemStore) InitStore() {
	var err error
	itemStore.DB, err = badger.Open(badger.DefaultOptions("data/items"))
	if err != nil {
		fmt.Println(err.Error())
	}
	defer itemStore.DB.Close()
}

func (itemStore *ItemStore) checkIfItemUIDExists(uid string) bool {
	err := itemStore.DB.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(uid))
		return err
	})
	return err != nil
}

func (itemStore *ItemStore) AddItemToDB(uid, category, description string, price float64, quantity uint64) error {
	res := itemStore.checkIfItemUIDExists(uid)
	if res {
		return errors.New("uid already exists")
	}
	item := Item{
		UID:         uid,
		Category:    category,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}

	err := itemStore.DB.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(uid), []byte(EncodeToJSON(item)))
		err := txn.SetEntry(e)
		return err
	})
	return err
}

func (itemStore *ItemStore) GetItemFromDB(uid string) (Item, error) {
	item := Item{}
	err := itemStore.DB.View(func(txn *badger.Txn) error {
		value, err := txn.Get([]byte(uid))
		if err != nil {
			return err
		}
		err = value.Value(func(val []byte) error {
			DecodeFromJSON(val, item)
			return nil
		})
		return err
	})
	return item, err
}

func (itemStore *ItemStore) GetItemsFromDB() ([]Item, error) {
	itemList := []Item{}
	var item Item
	err := itemStore.DB.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			value := it.Item()
			err := value.Value(func(v []byte) error {
				DecodeFromJSON(v, item)
				itemList = append(itemList, item)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	return itemList, err
}
