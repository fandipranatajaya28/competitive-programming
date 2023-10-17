package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	cache    map[int]*list.Element
	linklist *list.List
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    make(map[int]*list.Element, capacity),
		linklist: list.New(),
		capacity: capacity,
	}
}

// 0 is key and 1 is value
func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	elem := this.cache[key]
	this.linklist.MoveToFront(elem)
	return elem.Value.([]int)[1]
}

func (this *LRUCache) Put(key int, value int) {
	// if key already exist
	if elem, ok := this.cache[key]; ok {
		this.linklist.Remove(elem)
		newelem := this.linklist.PushFront([]int{key, value})
		this.cache[key] = newelem
		return
	}
	// if max capacity reached
	if len(this.cache) == this.capacity {
		elem := this.linklist.Back()
		v := this.linklist.Remove(elem)
		delete(this.cache, v.([]int)[0])
	}
	newelem := this.linklist.PushFront([]int{key, value})
	this.cache[key] = newelem
}

func main() {
	obj := Constructor(3)
	_ = obj.Get(3)
	obj.Put(2, 1)
	obj.Put(3, 2)

	for l := obj.linklist.Front(); l != nil; l = l.Next() {
		fmt.Println(l.Value)
	}
}
