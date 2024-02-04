package domain

type StoreCreated struct {
	Store *Store
}

func (StoreCreated) Name() string {
	return StoreCreatedEventName
}

type StoreParticipationEnabled struct {
	Store *Store
}

func (StoreParticipationEnabled) Name() string {
	return StoreParticipationEnabledEventName
}

type StoreParticipationDisabled struct {
	Store *Store
}

func (StoreParticipationDisabled) Name() string {
	return StoreParticipationDisabledEventName
}

var (
	StoreCreatedEventName               = "store.StoreCreated"
	StoreParticipationEnabledEventName  = "store.StoreParticipationEnabled"
	StoreParticipationDisabledEventName = "store.StoreParticipationDisabled"
)
