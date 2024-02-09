package ddd

import (
	"fmt"

	"github.com/vladesco/e-commerce/internal/registarable"
)

type IdSetter interface {
	setId(string)
}

func SetId(id string) registarable.BuildOption {
	return func(value any) error {
		if aggregate, ok := value.(IdSetter); ok {
			aggregate.setId(id)

			return nil
		} else {
			return fmt.Errorf("%T doesnt have the method setId(string)", value)
		}
	}
}

type NameSetter interface {
	setName(string)
}

func SetName(name string) registarable.BuildOption {
	return func(value any) error {
		if aggregate, ok := value.(NameSetter); ok {
			aggregate.setName(name)

			return nil
		} else {
			return fmt.Errorf("%T doesnt have the method setName(string)", value)
		}
	}
}
