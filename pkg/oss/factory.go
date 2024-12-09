package oss

type Factory interface {
	Create() (Client, error)
}

var factoryMap = map[string]Factory{}

func registerFactory(provider string, factory Factory) {
	factoryMap[provider] = factory
}

func getFactory(provider string) (Factory, bool) {
	factory, ok := factoryMap[provider]
	return factory, ok
}
