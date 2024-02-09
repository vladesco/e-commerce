package eventsourcing

import (
	"fmt"

	"github.com/vladesco/e-commerce/internal/registarable"
)

type VersionSetter interface {
	setVersion(int)
}

func SetVersion(version int) registarable.BuildOption {
	return func(value any) error {
		if aggregate, ok := value.(VersionSetter); ok {
			aggregate.setVersion(version)

			return nil
		} else {
			return fmt.Errorf("%T doesnt have the method setVersion(int)", value)
		}
	}
}
