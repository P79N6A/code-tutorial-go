package main
/*
Design a HashSet without using any built-in hash table libraries.

To be specific, your design should include these functions:

add(value): Insert a value into the HashSet.
contains(value) : Return whether the value exists in the HashSet or not.
remove(value): Remove a value in the HashSet. If the value does not exist in the HashSet, do nothing.

Example:

MyHashSet hashSet = new MyHashSet();
hashSet.add(1);
hashSet.add(2);
hashSet.contains(1);    // returns true
hashSet.contains(3);    // returns false (not found)
hashSet.add(2);
hashSet.contains(2);    // returns true
hashSet.remove(2);
hashSet.contains(2);    // returns false (already removed)

Note:

All values will be in the range of [0, 1000000].
The number of operations will be in the range of [1, 10000].
Please do not use the built-in HashSet library.
*/

func main() {
    
}

type ListBucket struct {
    Val  int
    Next *ListBucket
}
func get(l *ListBucket,k int) bool {
    if l == nil {return false}
    if l.Val == k {
        return true
    }
    return get(l.Next,k)
}
func remove(l *ListBucket,k int) *ListBucket {
    if l == nil {return nil}
    if l.Val == k {
        return l.Next
    }
    l.Next = remove(l.Next,k)
    return l
}
func put(l *ListBucket,k int) *ListBucket {
    return &ListBucket{k,l}
}

type MyHashSet struct {
    buckets []*ListBucket
    mod int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
    return MyHashSet{
        buckets:make([]*ListBucket,100001),
        mod:100001,
    }
}

func (this *MyHashSet) Add(key int)  {
    if !this.Contains(key) {
        this.buckets[this.idx(key)]=put(this.buckets[this.idx(key)],key)
    }
}

func (this *MyHashSet) Remove(key int)  {
    if this.Contains(key) {
        this.buckets[this.idx(key)]=remove(this.buckets[this.idx(key)],key)
    }
}

func (this *MyHashSet) idx(key int) int {
    return key%this.mod
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
    return get(this.buckets[this.idx(key)],key)
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
