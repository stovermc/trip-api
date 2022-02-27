package domain

type Trip interface {
	ID() string
	Name() string
}

type trip struct {
	id string
	name string
} 

func NewTrip(id, name string) Trip{
	return &trip{
		id: id,
		name: name,
	}
}

func (t *trip) ID() string {
	return t.id
}

func (t *trip) Name() string {
	return t.name
}