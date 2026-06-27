package main

import (
	"container/list"
	"sync"
)

type entry[T any] struct {
	key   int
	value T
}

type Cache[T any] struct {
	capacity int
	mu       sync.Mutex
	ll       *list.List
	items    map[int]*list.Element
}

func NewCache[T any](capacity int) *Cache[T] {
	if capacity < 0 {
		capacity = 0
	}
	return &Cache[T]{
		capacity: capacity,
		ll:       list.New(),
		items:    make(map[int]*list.Element),
	}
}

func (c *Cache[T]) Set(key int, value T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.capacity == 0 {
		return
	}

	if el, ok := c.items[key]; ok {
		el.Value.(*entry[T]).value = value
		c.ll.MoveToFront(el)
		return
	}

	el := c.ll.PushFront(&entry[T]{key: key, value: value})
	c.items[key] = el

	if c.ll.Len() > c.capacity {
		if back := c.ll.Back(); back != nil {
			ent := back.Value.(*entry[T])
			delete(c.items, ent.key)
			c.ll.Remove(back)
		}
	}
}

func (c *Cache[T]) Get(key int) (T, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var zero T
	el, ok := c.items[key]
	if !ok {
		return zero, false
	}

	c.ll.MoveToFront(el)
	return el.Value.(*entry[T]).value, true
}

func (c *Cache[T]) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.ll.Init()
	c.items = make(map[int]*list.Element, c.capacity)
}
