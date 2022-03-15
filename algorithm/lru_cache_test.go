package algorithm

import "testing"

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	i := cache.Get(1)
	cache.Put(3, 3)
	i = cache.Get(2)
	cache.Put(4, 4)
	i = cache.Get(1)
	i = cache.Get(3)
	i = cache.Get(4)
	t.Errorf("i is %v", i)
}
