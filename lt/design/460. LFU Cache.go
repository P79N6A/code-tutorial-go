package main


//Design and implement a data structure for Least Frequently Used (LFU) cache. It should support the following operations: get and put.
//
//get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
//put(key, value) - Set or insert the value if the key is not already present. When the cache reaches its capacity, it should invalidate the least frequently used item before inserting a new item. For the purpose of this problem, when there is a tie (i.e., two or more keys that have the same frequency), the least recently used key would be evicted.
//
//Follow up:
//Could you do both operations in O(1) time complexity?
//
//Example:
//
//LFUCache cache = new LFUCache( 2 /* capacity */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // returns 1
//cache.put(3, 3);    // evicts key 2
//cache.get(2);       // returns -1 (not found)
//cache.get(3);       // returns 3.
//cache.put(4, 4);    // evicts key 1.
//cache.get(1);       // returns -1 (not found)
//cache.get(3);       // returns 3
//cache.get(4);       // returns 4

func main() {
    
}
/*
问题：按照频次更换缓存
思路：
1.第一个map记录key,value；解决get
2.第二个map记录count, 需要看都出现了哪些freq的counter，min如果变了，往上加找到另外的最小min freq
3.第三个map key是 freq出现频次，value是个list，list支持操作delete，add，维持lru顺序，那就insert到head，delete在tail
看上去linkedHashMap才能满足需求
解决维护lru顺序删除数据问题。
3.min 记录最小的freq
*/


type listnode struct {
    key  int
    next *listnode
}
func get(head *listnode, idx int) int {
    if head == nil {
        return -1
    }
    if idx == 0 {
        return head.val
    }
    return get(head.next, idx-1)
}
func add(head *listnode, idx int, n *listnode) {
    if idx == 0 {
        n.next = head.next
        head.next = n
        return
    }
    add(head.next, idx-1, n)
}
func deletelist(head *listnode, idx int) {
    if head == nil||head.next == nil {return}
    // not first.
    if idx == 0 {
        head.next = head.next.next
        return
    }
    deletelist(head, idx-1)
}
//////////////////////////
type LinkList struct {
    head,tail *listnode
    lens int
}
func (l *LinkList)AddTail(k int) {

}
func (l *LinkList)DelByKey(k int)  {

}
func (l *LinkList)DelFront() int {

}
func (l *LinkList)Lens() int {
    return l.lens
}

type LFUCache struct {
    // key is key ,value is value
    value map[int]int
    // key is key ,value is freq
    count map[int]int
    // key is freq,value is list, for lru
    freq map[int]*LinkList
    min int
    cap int
}


func Constructor(capacity int) LFUCache {
    return LFUCache{
        value:make(map[int]int),
        count:make(map[int]int),
        freq:make(map[int]*LinkList),
        min:0,
        cap:capacity,
    }
}


func (this *LFUCache) Get(key int) int {
    if _,ok := this.value[key];!ok {return -1}
    freq := this.count[key] // 获取key当前频次
    this.count[key]+=1  // 频次+1
    this.freq[freq].DelByKey(key) // 从lru的list中删掉，挪到另外的freq队列，加到队尾，删除方便。
    if this.min==freq&&this.freq[freq].lens==0 {
        this.min+=1 // freq的删除key都需要触发是否更新min的操作
    }
    if this.freq[freq+1]==nil{
        this.freq[freq+1]=&LinkList{}
    }
    this.freq[freq+1].AddTail(key)
    return this.value[key]
}


func (this *LFUCache) Put(key int, value int)  {
    if _,ok := this.value[key];ok {
        //如果存在就是更新key对应的value就ok了
        this.value[key]=value
        this.Get(key)
        return
    }
    // 看下是否超过cap需要删掉最小的min对应的key，value
    if len(this.value) >= this.cap {
        evit := this.freq[this.min].DelFront()
        delete(this.value,evit)
        if this.freq[this.min].Lens() <= 0 {
            delete(this.freq,this.min)
        }
    }

    this.value[key]=value
    this.count[key]=1
    this.min = 1
    this.freq[1].AddTail(key)

}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */