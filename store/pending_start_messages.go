package store

import (
	"github.com/cloudfoundry/hm9000/models"
	"reflect"
)

func (store *RealStore) SavePendingStartMessages(messages ...models.PendingStartMessage) error {
	return store.save(messages, "/start", 0)
}

func (store *RealStore) GetPendingStartMessages() ([]models.PendingStartMessage, error) {
	slice, err := store.get("/start", reflect.TypeOf([]models.PendingStartMessage{}), reflect.ValueOf(models.NewPendingStartMessageFromJSON))
	return slice.Interface().([]models.PendingStartMessage), err
}

func (store *RealStore) DeletePendingStartMessages(messages ...models.PendingStartMessage) error {
	return store.delete(messages, "/start")
}