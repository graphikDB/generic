package generic

import "time"

// Closer closes a data-structure
type Closer interface {
	Close()
}

// Len returns the length of a data-structure
type LengthCheck interface {
	Len() int
}

// Cache is a general purpose, concurrency-safe, in-memory cache data-structure
type Cache interface {
	// Close closes the cache's garbage collector
	Closer
	// Len returns the length of the queue
	LengthCheck
	// Get gets an item from the cache by key
	Get(key interface{}) (interface{}, bool)
	// Exists returns true if the key exists in the cache
	Exists(key interface{}) bool
	// Set sets an item in the cache for the specified duration. If duration == 0, the item will never expire
	Set(key interface{}, value interface{}, duration time.Duration)
	// Range iterates over every item in the cache
	Range(f func(key, value interface{}) bool)
	// Delete deletes an item from the cache
	Delete(key interface{})
}

// Queue is a general purpose, in-memory queue data-structure
type Queue interface {
	// Enqueue adds a new node to the tail of the queue
	Enqueue(value interface{})
	// Dequeue removes the head node from the queue and returns it
	Dequeue() interface{}
	// Len returns the length of the queue
	LengthCheck
}

// Stack is a general purpose, in-memory stack data-structure
type Stack interface {
	Peek() interface{}
	Pop() interface{}
	Push(value interface{})
	LengthCheck
}
