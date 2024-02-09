package registarable

import (
	"fmt"
	"sync"

	"github.com/stackus/errors"
)

type (
	Registrable interface {
		GetKey() string
	}

	Factory      func() any
	Serializer   func(value any) ([]byte, error)
	Deserializer func(data []byte, value any) error

	Registry interface {
		Serialize(key string, value any) ([]byte, error)
		Deserialize(key string, data []byte, options ...BuildOption) (any, error)
		Build(key string, options ...BuildOption) (any, error)
		register(key string, factory Factory, serializer Serializer, deserialzer Deserializer, options []BuildOption) error
	}
)

type (
	registered struct {
		factory      Factory
		serializer   Serializer
		deserializer Deserializer
		options      []BuildOption
	}

	registry struct {
		sync.Mutex
		registered map[string]registered
	}
)

func NewRegistry() *registry {
	return &registry{
		registered: make(map[string]registered),
	}
}

func (registry *registry) Serialize(key string, value any) ([]byte, error) {
	registered, exists := registry.registered[key]

	if !exists {
		return nil, fmt.Errorf("not found any registered entities for key %s", key)
	}

	return registered.serializer(value)
}

func (registry *registry) Deserialize(key string, data []byte, options ...BuildOption) (any, error) {
	value, err := registry.Build(key, options...)

	if err != nil {
		return nil, errors.Wrap(err, "error while deserializing entity")
	}

	err = registry.registered[key].deserializer(data, value)

	if err != nil {
		return nil, errors.Wrap(err, "error while populating deserialized entity")
	}

	return value, nil
}

func (registry *registry) Build(key string, options ...BuildOption) (any, error) {
	registered, exists := registry.registered[key]

	if !exists {
		return nil, fmt.Errorf("not found any registered entities for key %s", key)
	}

	value := registered.factory()
	fullListOfOptions := append(registered.options, options...)

	for _, option := range fullListOfOptions {
		err := option(value)

		if err != nil {
			return nil, errors.Wrap(err, "error while building entity")
		}
	}

	return value, nil
}

func (registry *registry) register(key string, factory Factory, serializer Serializer, deserializer Deserializer, options []BuildOption) error {
	registry.Lock()
	defer registry.Unlock()

	if _, exists := registry.registered[key]; exists {
		return fmt.Errorf("entity with key %s has been already registered", key)
	}

	registry.registered[key] = registered{
		factory:      factory,
		serializer:   serializer,
		deserializer: deserializer,
		options:      options,
	}

	return nil
}

// CHECKS
var _ Registry = (*registry)(nil)
