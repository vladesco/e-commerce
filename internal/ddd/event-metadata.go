package ddd

type Metadata map[string]any

func (metadata Metadata) Set(key string, value any) {
	metadata[key] = value
}

func (metadata Metadata) Get(key string) any {
	value, ok := metadata[key]

	if !ok {
		return nil
	}

	return value
}

func (metadata Metadata) Del(key string) {
	delete(metadata, key)
}

func (metadata Metadata) configureEvent(e *event) {
	for key, value := range metadata {
		e.metadata[key] = value
	}
}

// CHECKS

var _ interface {
	EventOption
} = (*Metadata)(nil)
