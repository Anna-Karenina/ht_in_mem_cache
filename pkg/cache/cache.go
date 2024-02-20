package cache

import "errors"

type Cache struct {
	items map[string]interface{}
}

func NewCache() Cache {
	i := make(map[string]interface{})
	c := Cache{items: i}
	return c
}

func (c *Cache) Set(key string, value interface{}) {
	c.items[key] = value
}

func (c *Cache) Get(key string) (interface{}, error) {
	value, ok := c.items[key]
	if !ok {
		return nil, errors.New("key doesn't exist")
	}
	return value, nil
}

func (c *Cache) Delete(key string) int {
	delete(c.items, key)
	return len(c.items)
}

func (c *Cache) Len() int {
	return len(c.items)
}
