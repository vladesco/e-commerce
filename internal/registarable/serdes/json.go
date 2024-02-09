package serdes

import (
	"encoding/json"

	"github.com/vladesco/e-commerce/internal/registarable"
)

type JsonSerde struct {
	registry registarable.Registry
}

func NewJsonSerde(registry registarable.Registry) *JsonSerde {
	return &JsonSerde{
		registry: registry,
	}
}

func (serde *JsonSerde) Register(value registarable.Registrable, options ...registarable.BuildOption) error {
	return registarable.Register(serde.registry, value, serde.serialize, serde.deserialize, options)
}

func (serde *JsonSerde) RegisterKey(key string, value any, options ...registarable.BuildOption) error {
	return registarable.RegisterKey(serde.registry, key, value, serde.serialize, serde.deserialize, options)
}

func (serde *JsonSerde) RegisterFactory(key string, facotory registarable.Factory, options ...registarable.BuildOption) error {
	return registarable.RegisterFactory(serde.registry, key, facotory, serde.serialize, serde.deserialize, options)
}

func (serde *JsonSerde) serialize(value any) ([]byte, error) {
	return json.Marshal(value)
}

func (serde *JsonSerde) deserialize(data []byte, value any) error {
	return json.Unmarshal(data, value)
}

// CHECKS
var _ registarable.Serde = (*JsonSerde)(nil)
