package cache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	items map[string]storageItem
	mu    sync.RWMutex
}

type storageItem struct {
	value      interface{}
	expireDate time.Time
}

func isOverdue(sItem storageItem) bool {
	return time.Now().After(sItem.expireDate)
}

func NewCache() *Cache {
	i := make(map[string]storageItem)
	c := Cache{items: i}
	go c.cleanerRun()
	return &c
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	//add validations to key exist, value exist ttl exist
	item := storageItem{value: value, expireDate: time.Now().Add(ttl)}
	c.items[key] = item
	return nil
}

func (c *Cache) Get(key string) (interface{}, error) {
	value, ok := c.items[key]
	if !ok {
		return nil, errors.New("key doesn't exist")
	}
	if isOverdue(value) {
		return nil, errors.New("value expired")
	}
	return value, nil
}

func (c *Cache) Delete(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
	return len(c.items)
}

func (c *Cache) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.items)
}
