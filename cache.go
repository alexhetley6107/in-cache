package cache

type Cache map[string]int

func New() *Cache {
	cache := Cache(make(map[string]int))
	return &cache
}

func (c *Cache) Set(key string, value int) {
	(*c)[key] = value
}

func (c *Cache) Get(key string) int {
	return (*c)[key]
}

func (c *Cache) Delete(key string) {
	delete(*c, key)
}