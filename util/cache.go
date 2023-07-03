package util

type Node[T any] struct {
	key  string
	data T
	prev *Node[T]
	next *Node[T]
}

type LRUCache[T any] struct {
	head      *Node[T]
	tail      *Node[T]
	maxLength int
	length    int
	cacheMap  map[string]*Node[T]
}

func (cache *LRUCache[T]) InitializeCache(maxSize int) *LRUCache[T] {
	newLruCache := LRUCache[T]{
		maxLength: maxSize,
		length:    0,
		cacheMap:  make(map[string]*Node[T]),
	}

	return &newLruCache
}

func (cache *LRUCache[T]) IncreaseCacheSize(addedSpace int) {
	cache.maxLength += addedSpace
}

func (cache *LRUCache[T]) Put(key string, data T) {
	newCacheItem := &Node[T]{
		key:  key,
		data: data,
	}
	if cache.head == nil {
		cache.head = newCacheItem
		cache.tail = newCacheItem
	} else {
		newCacheItem = cache.AddItemToFront(newCacheItem)
	}
	cacheMap := cache.cacheMap
	cacheMap[key] = newCacheItem
	cache.cacheMap = cacheMap

	if cache.length == cache.maxLength {
		cache.RemoveLastCacheItem()
	}

	cache.length = cache.length + 1

}

func (cache *LRUCache[T]) RemoveItem(cacheItem *Node[T]) {

	if cache.head == nil {
		return
	}

	if cache.head == cacheItem {
		cache.head = cacheItem.next
	}

	if cacheItem.prev != nil {
		previousItem := cacheItem.prev
		previousItem.next = cacheItem.next
		if cacheItem.next != nil {
			nextItem := cacheItem.next
			nextItem.prev = previousItem
		}
	}

}

func (cache *LRUCache[T]) RemoveLastCacheItem() {
	lastItem := cache.tail
	if lastItem != nil {
		previousItem := lastItem.prev
		cache.tail = previousItem
		previousItem.next = nil
		delete(cache.cacheMap, lastItem.key)
	}
	cache.length = cache.length - 1
}

func (cache *LRUCache[T]) AddItemToFront(cacheItem *Node[T]) *Node[T] {
	if cache.head == nil {
		cache.head = cacheItem
		cache.tail = cacheItem
		cacheItem.prev = nil
		cacheItem.next = nil
	}

	if cache.head != nil {
		currentHead := cache.head
		cacheItem.next = currentHead
		cacheItem.prev = nil
		currentHead.prev = cacheItem
		cache.head = cacheItem
	}

	return cacheItem

}

func (cache *LRUCache[T]) Get(key string) *Node[T] {
	if value, ok := cache.cacheMap[key]; ok {
		if value != cache.head {
			cache.RemoveItem(value)
			cache.AddItemToFront(value)
		}
		return value
	} else {
		return nil
	}
}
