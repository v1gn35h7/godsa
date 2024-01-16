package pro

import (
	"container/list"
	"fmt"
	"sync"
)

type CacheStore interface {
	Get(string) interface{}
	Set(string) bool
}

type cacheStore struct {
	cache    map[string]*list.Element
	data     list.List
	mux      sync.RWMutex
	capacity int
}

type Item struct {
	k string
	v int
}

func NewCacheStore() *cacheStore {
	return &cacheStore{
		cache:    make(map[string]*list.Element),
		data:     *list.New(),
		mux:      sync.RWMutex{},
		capacity: 10,
	}
}

func (cs *cacheStore) Get(key string) (interface{}, error) {

	if res, ok := cs.cache[key]; !ok {
		return res, fmt.Errorf("Key not found")
	} else {
		cs.mux.Lock()
		defer cs.mux.Unlock()
		fmt.Println(cs.data.Front())
		cs.data.MoveToFront(res)
		fmt.Println(cs.data.Front())

		return res, nil
	}
}

func (cs *cacheStore) Set(key string, val interface{}) bool {
	cs.mux.Lock()
	defer cs.mux.Unlock()
	v, _ := val.(int)

	if res, ok := cs.cache[key]; !ok {
		if cs.data.Len() >= cs.capacity {
			cs.evict()
		}

		var i *list.Element

		if cs.data.Len() > 0 {
			i = cs.data.InsertBefore(Item{key, v}, cs.data.Front())
		} else {
			i = cs.data.PushFront(Item{key, v})
		}
		cs.cache[key] = i

		fmt.Println("Item -->", i)
	} else {
		i := res.Value.(Item)
		i.v = v
		cs.data.MoveToFront(res)
		fmt.Println("Item -->", i)
	}
	return true
}

func (cs *cacheStore) evict() {
	last := cs.data.Back()
	c := cs.data.Remove(last)
	i, _ := last.Value.(Item)

	if c == nil {
		fmt.Println(" Failed to removed", c, cs.data.Len(), i)
	}

	delete(cs.cache, i.k)
}

func (cs *cacheStore) Flush() {
	fmt.Println(len(cs.cache))
	fmt.Println(cs.data.Len())
}
