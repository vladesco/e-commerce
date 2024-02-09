package ddd

import (
	"fmt"

	"github.com/vladesco/e-commerce/internal/registarable"
)

type EventSetter interface {
	setEvents([]Event)
}

func SetEvents(events ...Event) registarable.BuildOption {
	return func(value any) error {
		if aggregate, ok := value.(EventSetter); ok {
			aggregate.setEvents(events)

			return nil
		} else {
			return fmt.Errorf("%T doesnt have the method setEvents([]Event)", value)
		}
	}
}
