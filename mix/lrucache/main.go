package main

import "fmt"

// https://medium.com/@fazlulkabir94/lru-cache-golang-implementation-92b7bafb76f0

// LRUCache holds the cache information
type LRUCache struct {
	capacity int
	size     int
	pageList Queue
	pageMap  map[int]*Qnode
}

// Queue holds information about a queue
type Queue struct {
	front *Qnode
	rear  *Qnode
}

// Qnode holds information about a node in a queue
type Qnode struct {
	key   int
	value int
	prev  *Qnode
	next  *Qnode
}

// addQnode creates a node with the given values
func addQnode(key int, value int) *Qnode {
	return &Qnode{
		key:   key,
		value: value,
		prev:  nil,
		next:  nil,
	}
}

// isEmpty returns true if the queue is empty
func (q *Queue) isEmpty() bool {
	return q.rear == nil
}

// addFrontPage add a new page to the front of the queue
func (q *Queue) addFrontPage(key int, value int) *Qnode {
	page := addQnode(key, value)

	// queue is empty
	if q.front == nil && q.rear == nil {
		q.front = page
		q.rear = page
	} else {
		// Otherwise, we move the current front to the next one
		// and put at the front the added page.
		page.next = q.front
		q.front.prev = page
		q.front = page
	}
	return page
}

// moveToFront moves to the front an existing page
func (q *Queue) moveToFront(page *Qnode) {
	if page == q.front {
		// we do nothing
		return
	}
	// given page is at the end of the queue
	if page == q.rear {
		q.rear = q.rear.prev
		q.rear.next = nil
	} else {
		page.prev.next = page.next
		page.next.prev = page.prev
	}

	// put the page to the front of the queue
	page.next = q.front
	q.front.prev = page
	q.front = page
}

// removeRear removes the rear from the queue
func (q *Queue) removeRear() {
	if q.isEmpty() {
		return
	}
	// if front and rear are equals, let's just set them to 0
	if q.front == q.rear {
		q.front = nil
		q.rear = nil
	} else {
		// case they are not equal
		// the current rear is set to the previous one and the next one is nil
		q.rear = q.rear.prev
		q.rear.next = nil
	}
}

// getRear returns the given rear
func (q *Queue) getRear() *Qnode {
	return q.rear
}

// Constructor initializes the LRUCache with the given capacity
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		pageMap:  make(map[int]*Qnode),
	}
}

// Get returns the value corresponding to the given key
// if not found, return -1
// else move the page to the front and return the value
func (lru *LRUCache) Get(key int) int {
	if _, found := lru.pageMap[key]; !found {
		// the value is not found
		return -1
	}

	// retrieve the value
	val := lru.pageMap[key].value
	// move to front the given page
	lru.pageList.moveToFront(lru.pageMap[key])

	return val
}

// Put inserts a new page in the queue
func (lru *LRUCache) Put(key int, value int) {
	if _, found := lru.pageMap[key]; found {
		// update the key value with the given one
		lru.pageMap[key].value = value
	}

	// case the size and the capacity are the same
	if lru.size == lru.capacity {
		// retrieve the key of the rear
		key := lru.pageList.getRear().key
		// remove the rears
		lru.pageList.removeRear()
		// decrease the size and delete the key
		lru.size--
		delete(lru.pageMap, key)
	}

	page := lru.pageList.addFrontPage(key, value)
	fmt.Println(lru.pageMap)
	lru.size++
	lru.pageMap[key] = page
}

func main() {
	cache := Constructor(2)
	cache.Put(2,2) // 2:2
	fmt.Println(cache.Get(2)) // 2
	fmt.Println(cache.Get(1)) // -1
	cache.Put(1,1) // 1:1
	cache.Put(1,5) // 1:5
	fmt.Println(cache.Get(1)) // 5
	fmt.Println(cache.Get(2)) // 2
	cache.Put(8, 8) // 8
	fmt.Println(cache.Get(1)) // 5
	fmt.Println(cache.Get(8)) // 8
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
