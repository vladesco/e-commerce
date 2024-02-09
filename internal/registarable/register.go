package registarable

import (
	"fmt"
	"reflect"
)

func Register(registry Registry, value Registrable, serializer Serializer, deserializer Deserializer, options []BuildOption) error {
	valueType := reflect.TypeOf(value)
	var valueKey string

	if valueType.Kind() == reflect.Pointer && reflect.ValueOf(value).IsNil() {
		valueKey = reflect.New(valueType).Interface().(Registrable).GetKey()
	} else {
		valueKey = value.GetKey()
	}

	return RegisterKey(registry, valueKey, value, serializer, deserializer, options)
}

func RegisterKey(registry Registry, key string, value any, serializer Serializer, deserializer Deserializer, options []BuildOption) error {
	valueType := reflect.TypeOf(value)

	if valueType.Kind() == reflect.Pointer {
		value = valueType.Elem()
	}

	return registry.register(key, func() any {
		return reflect.New(valueType)
	}, serializer, deserializer, options)
}

func RegisterFactory(registry Registry, key string, factory Factory, serializer Serializer, deserializer Deserializer, options []BuildOption) error {
	if factory() == nil {
		return fmt.Errorf("facotory for entity %s can't return nil", key)
	}

	if reflect.TypeOf(factory()).Kind() != reflect.Pointer {
		return fmt.Errorf("facotory for entity %s should return only pointer", key)
	}

	return registry.register(key, factory, serializer, deserializer, options)
}
