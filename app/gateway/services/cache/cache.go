package cache

//go:generate moq -fmt goimports -out cache_mock.gen.go . Cache:CacheMock
type Cache interface {
	Get(key string) (interface{}, error)
	Save(key string, value []byte) error
}
