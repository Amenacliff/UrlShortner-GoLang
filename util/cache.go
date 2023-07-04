package util

import (
	"encoding/json"
	"log"
)

type Node struct {
	key  string
	Data []byte
	prev *Node
	next *Node
}

type LRUCache struct {
	head      *Node
	tail      *Node
	maxLength int
	length    int
	cacheMap  map[string]*Node
}

func (cache *LRUCache) InitializeCache(maxSize int) *LRUCache {
	newLruCache := LRUCache{
		maxLength: maxSize,
		length:    0,
		cacheMap:  make(map[string]*Node),
	}

	return &newLruCache
}

func (cache *LRUCache) IncreaseCacheSize(addedSpace int) {
	cache.maxLength += addedSpace
}

func (cache *LRUCache) Put(key string, data any) {

	byteData, err := json.Marshal(data)

	if err != nil {
		log.Println(err.Error())
	}

	node := cache.Get(key)

	if node != nil {
		cache.RemoveItem(node)
	}

	newCacheItem := &Node{
		key:  key,
		Data: byteData,
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

func (cache *LRUCache) RemoveItem(cacheItem *Node) {

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

func (cache *LRUCache) RemoveLastCacheItem() {
	lastItem := cache.tail
	if lastItem != nil {
		previousItem := lastItem.prev
		cache.tail = previousItem
		previousItem.next = nil
		delete(cache.cacheMap, lastItem.key)
	}
	cache.length = cache.length - 1
}

func (cache *LRUCache) AddItemToFront(cacheItem *Node) *Node {
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

func (cache *LRUCache) Get(key string) *Node {
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
