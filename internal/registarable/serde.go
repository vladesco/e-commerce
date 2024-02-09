package registarable

type Serde interface {
	Register(value Registrable, options ...BuildOption) error
	RegisterKey(key string, value any, options ...BuildOption) error
	RegisterFactory(key string, factory Factory, options ...BuildOption) error
}
