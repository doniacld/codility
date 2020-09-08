package main

import "fmt"

// LRUCache
type LRUCache struct {
	capacity int
	size     int
	head     *Node
	tail     *Node
	mapPage  map[int]*Node
}

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

// Constructor initializes the LRUCache object with the provided capacity
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		mapPage:  make(map[int]*Node),
	}
}

func moveToEnd(cache *LRUCache, node *Node) {
	// already at tail
	if node == cache.tail {
		return
	}

	oldPrev := node.prev
	oldNext := node.next

	// case node is the head
	if node == cache.head {
		// remove from head of chain
		cache.head = oldNext
		oldNext.prev = nil
	} else {
		// remove from middle of chain
		oldPrev.next = oldNext
		oldNext.prev = oldPrev
	}

	// add at the end
	oldTail := cache.tail
	oldTail.next = node

	cache.tail = node
	node.prev = oldTail
	node.next = nil
}

// Get returns the value of a key
// and updates the head
func (cache *LRUCache) Get(key int) int {
	if pointer, ok := cache.mapPage[key]; ok {
		// case found, update to tail
		moveToEnd(cache, pointer)
		return pointer.value
	} else {
		// not found
		return -1
	}
}

func (cache *LRUCache) Put(key, value int) {
	if pointer, ok := cache.mapPage[key]; ok {
		// value is found, move it to the end
		pointer.value = value
		moveToEnd(cache, pointer)
	} else {
		if cache.size < cache.capacity {
			// increase the size
			cache.size++
		} else {
			// remove the one at the end
			delete(cache.mapPage, cache.head.key)
			if cache.head != nil {
				cache.head.prev = nil
			}
		}

	}
	n := Node{
		key:   key,
		value: value,
		prev:  cache.tail,
		next:  nil,
	}
	cache.mapPage[key] = &n

	// update old tail
	if cache.tail != nil {
		cache.tail.next = &n
	}

	cache.tail = &n

	// update head if needed
	if cache.head == nil {
		cache.head = &n

	}
}

func main() {
	cache := Constructor(2)
	cache.Put(2, 2)           // 2:2
	fmt.Println(cache.Get(2)) // 2
	fmt.Println(cache.Get(1)) // -1
	cache.Put(1, 1)           // 1:1
	cache.Put(1, 5)           // 1:5
	fmt.Println(cache.Get(1)) // 5
	fmt.Println(cache.Get(2)) // 2
	cache.Put(8, 8)           // 8:8
	fmt.Println(cache.Get(1)) // -1
	fmt.Println(cache.Get(8)) // 8
}
