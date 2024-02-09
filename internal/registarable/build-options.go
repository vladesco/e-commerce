package registarable

import (
	"fmt"
	"reflect"
)

type BuildOption func(value any) error

func ValidateImplementation(checkedValue any) BuildOption {
	checkedType := reflect.TypeOf(checkedValue)

	if checkedType.Kind() == reflect.Pointer {
		checkedType = checkedType.Elem()
	}

	if checkedType.Kind() != reflect.Interface {
		panic(fmt.Sprintf("%T is not an interface", checkedValue))
	}

	return func(value any) error {
		valueType := reflect.TypeOf(value)

		if !valueType.Implements(checkedType) {
			return fmt.Errorf("%T does not implement %T", value, checkedType)
		}

		return nil
	}
}
