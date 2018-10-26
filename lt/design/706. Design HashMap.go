package main

import "fmt"

/*
Design a HashMap without using any built-in hash table libraries.

To be specific, your design should include these functions:

put(key, value) : Insert a (key, value) pair into the HashMap. If the value already exists in the HashMap, update the value.
get(key): Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
remove(key) : Remove the mapping for the value key if this map contains the mapping for the key.

Example:

MyHashMap hashMap = new MyHashMap();
hashMap.put(1, 1);
hashMap.put(2, 2);
hashMap.get(1);            // returns 1
hashMap.get(3);            // returns -1 (not found)
hashMap.put(2, 1);          // update the existing value
hashMap.get(2);            // returns 1
hashMap.remove(2);          // remove the mapping for 2
hashMap.get(2);            // returns -1 (not found)

Note:

All keys and values will be in the range of [0, 1000000].
The number of operations will be in the range of [1, 10000].
Please do not use the built-in HashMap library.
*/

func main() {
    hashmap := Constructor()
    hashmap.Put(1,1)
    hashmap.Put(2,2)
    fmt.Println(hashmap.Get(1))
    fmt.Println(hashmap.Get(3))
    hashmap.Put(2,1)
    fmt.Println(hashmap.Get(2))
    hashmap.Remove(2)
    fmt.Println(hashmap.Get(2))
}

type ListBucket struct {
    Val  *Bucket
    Next *ListBucket
}
func get(l *ListBucket,k int) int {
    if l == nil {return -1}
    if l.Val.key == k {
        return l.Val.value
    }
    return get(l.Next,k)
}
func remove(l *ListBucket,k int) *ListBucket {
    if l == nil {return nil}
    if l.Val.key == k {
        return l.Next
    }
    l.Next = remove(l.Next,k)
    return l
}
func put(l *ListBucket,k,v int) *ListBucket {
    return &ListBucket{&Bucket{k,v},l}
}
func update(l *ListBucket,k,v int) {
    if l == nil {return}
    if l.Val.key == k {
        l.Val.value = v
    } else {
        update(l.Next,k,v)
    }
}

type Bucket struct {
    key,value int
}

type MyHashMap struct {
    bucket []*ListBucket
    mod int
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
    return MyHashMap{
        bucket:make([]*ListBucket,1000001),
        mod:1000001,
    }
}

func (this *MyHashMap) idx(key int) int {
    return key % this.mod
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
    if this.Get(key) == -1 {
        this.bucket[this.idx(key)]=put(this.bucket[this.idx(key)],key,value)
    } else {
        update(this.bucket[this.idx(key)],key,value)
    }
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
    return get(this.bucket[this.idx(key)],key)
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
    if this.Get(key) == -1 {return}
    this.bucket[this.idx(key)]=remove(this.bucket[this.idx(key)],key)
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */