package domain

import (
	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/internal/ddd"
)

type Store struct {
	ddd.AggregateBase
	name          string
	participating bool
}

func CreateStore(name string) (*Store, error) {
	if name == "" {
		return nil, ErrorStoreNameIsMissed
	}

	store := &Store{
		ddd.CrateAggregate(),
		name,
		false,
	}

	store.AddEvent(&StoreCreated{
		store,
	})

	return store, nil
}

func (store *Store) EnableParticipation() (err error) {
	if store.participating {
		return ErrorStoreIsAlreadyParticipating
	}

	store.participating = true

	store.AddEvent(&StoreParticipationEnabled{
		store,
	})

	return
}

func (store *Store) DisableParticipation() (err error) {
	if !store.participating {
		return ErrorStoreIsAlreadyNotParticipating
	}

	store.participating = false

	store.AddEvent(&StoreParticipationDisabled{
		store,
	})

	return
}

var (
	ErrorStoreNameIsMissed              = errors.Wrap(errors.ErrBadRequest, "[STORE]: store name is missed")
	ErrorStoreIsAlreadyParticipating    = errors.Wrap(errors.ErrBadRequest, "[STORE]: store is already participating")
	ErrorStoreIsAlreadyNotParticipating = errors.Wrap(errors.ErrBadRequest, "[STORE]: store is already not participating")
)
