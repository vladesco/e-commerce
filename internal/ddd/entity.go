package ddd

type Entity interface {
	GetId() string
}

type EntityBase struct {
	id string
}

func (entity *EntityBase) GetId() string {
	return entity.id
}
