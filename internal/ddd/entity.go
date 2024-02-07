package ddd

type Entity struct {
	id   string
	name string
}

func NewEntity(id, name string) Entity {
	return Entity{
		id:   id,
		name: name,
	}
}

func (entity *Entity) GetId() string          { return entity.id }
func (entity *Entity) GetName() string        { return entity.name }
func (entity *Entity) Equals(other Ider) bool { return entity.id == other.GetId() }

func (entity *Entity) setId(id string)     { entity.id = id }
func (entity *Entity) setName(name string) { entity.name = name }

// CHECKS
var _ interface {
	Ider
	Namer
} = (*Entity)(nil)
