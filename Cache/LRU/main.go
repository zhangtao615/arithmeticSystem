package LRU

import "container/list"

type Cache interface {
	Set(key string, value interface{})

	Get(key string) interface{}

	Del(key string)

	DelOldest()

	UseBytes()
}

type Value interface {
	Len() int
}

type entry struct {
	key   string
	value interface{}
}

type lru struct {
	maxBytes int
	useBytes int
	ll       *list.List
	cache    map[string]*list.Element
}

func (lru *lru) Set(key string, value interface{}) {
	if element, ok := lru.cache[key]; ok {
		lru.ll.MoveToFront(element)
	} else {
		if lru.useBytes > lru.maxBytes {

		}
	}
}
