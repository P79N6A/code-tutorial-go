package main

import (
    "container/list"
    "fmt"
)

/*
Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

Follow up:
Could you do both operations in O(1) time complexity?

Example:

LRUCache cache = new LRUCache( 2 /* capacity */

//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // returns 1
//cache.put(3, 3);    // evicts key 2
//cache.get(2);       // returns -1 (not found)
//cache.put(4, 4);    // evicts key 1
//cache.get(1);       // returns -1 (not found)
//cache.get(3);       // returns 3
//cache.get(4);       // returns 4

func main() {
    cache := Constructor(2)
    cache.Put(1,1)
    cache.Put(2,2)
    fmt.Println(cache.Get(1))
    cache.Put(3,3)
    fmt.Println(cache.Get(2))
    cache.Put(4,4)
    fmt.Println(cache.Get(1))
    fmt.Println(cache.Get(3))
    fmt.Println(cache.Get(4))

}

type bucket struct {
    k,v int
}

type LRUCache struct {
    list *list.List
    index map[int]*list.Element
    cap int
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
        list:list.New(),
        index:make(map[int]*list.Element),
        cap:capacity,
    }
}


func (this *LRUCache) Get(key int) int {
    if _,ok := this.index[key];!ok {
        return -1
    }
    x := this.index[key]
    this.list.MoveToFront(x)
    return x.Value.(*bucket).v
}


func (this *LRUCache) Put(key int, value int)  {
    if _,ok := this.index[key];ok {
        this.index[key].Value = &bucket{key,value}
        this.Get(key)
        return
    }
    if len(this.index)>=this.cap {
        x := this.list.Back()
        this.list.Remove(this.list.Back())
        delete(this.index,x.Value.(*bucket).k)
    }
    this.list.PushFront(&bucket{key,value})
    this.index[key]=this.list.Front()
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
