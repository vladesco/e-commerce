package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/stackus/errors"
)

type StoreRepository interface {
	Save(ctx context.Context, store *Store) error
	Update(ctx context.Context, store *Store) error
	Delete(ctx context.Context, storeId string) error
	Find(ctx context.Context, storeId string) (*Store, error)
	FindAll(ctx context.Context) ([]*Store, error)
}

type Store struct {
	Id            string
	Name          string
	Participating bool
}

func CreateStore(name string) (*Store, error) {
	if name == "" {
		return nil, ErrorStoreNameIsMissed
	}

	store := &Store{
		uuid.NewString(),
		name,
		false,
	}

	return store, nil
}

func (store *Store) EnableParticipation() (err error) {
	if store.Participating {
		return ErrorStoreIsAlreadyParticipating
	}

	store.Participating = true

	return
}

func (store *Store) DisableParticipation() (err error) {
	if !store.Participating {
		return ErrorStoreIsAlreadyNotParticipating
	}

	store.Participating = false
	return
}

var (
	ErrorStoreNameIsMissed              = errors.Wrap(errors.ErrBadRequest, "[STORE]: store name is missed")
	ErrorStoreIsAlreadyParticipating    = errors.Wrap(errors.ErrBadRequest, "[STORE]: store is already participating")
	ErrorStoreIsAlreadyNotParticipating = errors.Wrap(errors.ErrBadRequest, "[STORE]: store is already not participating")
)
