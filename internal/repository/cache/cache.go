package cache

import (
	"sync"
)

type Cache struct {
	Mutex sync.RWMutex
	Data  map[string]string
}

func NewCache() *Cache {
	var cache Cache
	cache.Data = make(map[string]string, 0)

	return &cache
}
