package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type CacheItem struct {
	Key   Key
	Value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	item, exists := c.items[key]
	if exists {
		item.Value = CacheItem{Key: key, Value: value}
		c.queue.MoveToFront(item)
		return true
	}
	newItem := c.queue.PushFront(CacheItem{Key: key, Value: value})
	c.items[key] = newItem
	if c.queue.Len() > c.capacity {
		lastItem := c.queue.Back()
		c.queue.Remove(lastItem)
		delete(c.items, lastItem.Value.(CacheItem).Key)
	}
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	item, exists := c.items[key]
	if exists {
		c.queue.MoveToFront(item)
		cacheItem := item.Value.(CacheItem)
		return cacheItem.Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
