package main

import (
	"fmt"
)

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	// Create dummy head and tail
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	node, isExist := this.cache[key]
	if isExist {
		this.moveToFront(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// If key exists, update and move to front
	if node, isExist := this.cache[key]; isExist {
		node.value = value
		this.moveToFront(node)
		return
	}

	// If at capacity, remove least recently used
	if len(this.cache) == this.capacity {
		lru := this.tail.prev
		this.removeNode(lru)
		delete(this.cache, lru.key)
	}

	// Add new node
	newNode := &Node{
		key:   key,
		value: value,
	}
	this.addToFront(newNode)
	this.cache[key] = newNode
}

func (this *LRUCache) moveToFront(node *Node) {
	this.removeNode(node)
	this.addToFront(node)
}

func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) addToFront(node *Node) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

func main() {
	cache := Constructor(3)

	fmt.Println(cache.Get(3)) // -1

	cache.Put(2, 1)
	cache.Put(3, 2)

	fmt.Println(cache.Get(2)) // 1
	fmt.Println(cache.Get(3)) // 2

	cache.Put(4, 3)
	cache.Put(5, 4)

	fmt.Println(cache.Get(2)) // -1 (evicted)
	fmt.Println(cache.Get(4)) // 3
}
