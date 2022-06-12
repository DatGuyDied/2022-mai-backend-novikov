package cache

import "container/list"

type keyValue struct {
	key, value string
}

type LRU struct {
	size, cap int

	list  *list.List
	cache map[string]*list.Element
}

func NewLRU(cap int) *LRU {
	return &LRU{
		size:  0,
		cap:   cap,
		list:  list.New(),
		cache: make(map[string]*list.Element, cap),
	}
}

func (c *LRU) Get(key string) (string, bool) {
	e, ok := c.cache[key]
	if !ok {
		return "", false
	}

	c.list.MoveToFront(e)

	return e.Value.(keyValue).value, true
}

func (c *LRU) Set(key, value string) {
	e, ok := c.cache[key]
	if ok {
		e.Value = keyValue{key, value}
		c.list.MoveToFront(e)
		return
	}

	c.cache[key] = c.list.PushFront(keyValue{key, value})
	c.size++

	if c.size > c.cap {
		last := c.list.Back().Value.(keyValue)
		c.Remove(last.key)
	}
}

func (c *LRU) Remove(key string) {
	e, ok := c.cache[key]
	if !ok {
		return
	}

	c.list.Remove(e)
	delete(c.cache, key)
	c.size--
}
