package main

import (
    "fmt"
    "strconv"
)


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
    lfu := Constructor(10)
    /*
    lfu.Put(1, 1)
    lfu.Put(2, 2)
    fmt.Println(lfu.Get(1))
    lfu.Put(3, 3)
    fmt.Println(lfu.Get(2))
    fmt.Println(lfu.Get(3))
    lfu.Put(4, 4)
    fmt.Println(lfu.Get(1))
    fmt.Println(lfu.Get(3))
    fmt.Println(lfu.Get(4))
    */
    //["LFUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]
    //[[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]
    a := []string{"LFUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"}
    b := [][]int{{10},{10,13},{3,17},{6,11},{10,5},{9,10},{13},{2,19},{2},{3},{5,25},{8},{9,22},{5,5},{1,30},{11},{9,12},{7},{5},{8},{9},{4,30},{9,3},{9},{10},{10},{6,14},{3,1},{3},{10,11},{8},{2,14},{1},{5},{4},{11,4},{12,24},{5,18},{13},{7,23},{8},{12},{3,27},{2,12},{5},{2,9},{13,4},{8,18},{1,7},{6},{9,29},{8,21},{5},{6,30},{1,12},{10},{4,15},{7,22},{11,26},{8,17},{9,29},{5},{3,4},{11,30},{12},{4,29},{3},{9},{6},{3,4},{1},{10},{3,29},{10,28},{1,20},{11,13},{3},{3,12},{3,8},{10,9},{3,26},{8},{7},{5},{13,17},{2,27},{11,15},{12},{9,19},{2,15},{3,16},{1},{12,17},{9,1},{6,19},{4},{5},{5},{8,1},{11,7},{5,2},{9,28},{1},{2,2},{7,4},{4,22},{7,24},{9,26},{13,28},{11,26}}
    ans := make([]int,0)
    for i:=0;i<len(a);i++ {
        fmt.Println(a[i],b[i])
        if a[i]=="put" {
            lfu.Put(b[i][0],b[i][1])
            //lfu.Debug("put",b[i][0],b[i][1])
        } else if a[i]=="get" {
            ans = append(ans,lfu.Get(b[i][0]))
            //lfu.Debug("get",b[i][0],-1)
        }
    }
    fmt.Println(ans)

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
//////////////////////////
type LinkList struct {
    head, tail *listnode
    lens       int
}
func (l *LinkList)String() string {
    str := strconv.Itoa(l.lens)
    str += "/"
    if l.head != nil {
        x := l.head
        for x != nil {
            str += strconv.Itoa(x.key) + "->"
            x = x.next
        }
    }
    str += "/"
    if l.tail != nil {
        str += strconv.Itoa(l.tail.key) + "ex"
        if l.tail.next == nil {
            str += "NIL"
        }
    }
    return str
}

func (l *LinkList)AddTail(k int) {
    n := &listnode{k, nil}
    if l.tail == nil {
        l.tail = n
        l.head = n
    } else {
        l.tail.next = n
        l.tail = l.tail.next
    }
    l.lens += 1
}
func (l *LinkList)DelByKey(k int) {
    l.head = deletek(l.head, k)
    if l.head == nil {
        l.tail = nil
    }
    // 曾经出现大bug，比如删掉的key正好是最后一个，那tail是无法被修复的
    if l.tail!=nil&&l.tail.key == k {
        l.tail = l.head
        for l.tail.next != nil {
            l.tail=l.tail.next
        }
    }
    l.lens -= 1
}
func deletek(h *listnode, k int) *listnode {
    if h == nil {
        return nil
    }
    if h.key == k {
        return h.next
    }
    h.next = deletek(h.next, k)
    return h
}
func (l *LinkList)DelFront() int {
    x := l.head.key
    l.head = l.head.next
    if l.head == nil {
        l.tail = nil
    }
    l.lens -= 1
    return x
}
func (l *LinkList)Lens() int {
    if l == nil {
        return 0
    }
    return l.lens
}

type LFUCache struct {
    // key is key ,value is value
    value map[int]int
    // key is key ,value is freq
    count map[int]int  // 可以快速定位指定key的freq,从这个来在freq里面索引key-counter
    // key is freq,value is list, for lru
    // LinkList需要支持的操作: delFront[LRU,最近未使用的],delbykey[按照key删除],在队尾增加
    freq  map[int]*LinkList
    min   int
    cap   int
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

func (this *LFUCache) Debug(op string,k,v int)  {
    fmt.Println(op,"-----------",k,v)
    fmt.Println("Value:",this.value)
    fmt.Println("Count:",this.count)
    for k,v := range this.freq {
        fmt.Println("%%%%%",k,"%%%%%%%",v.String())
    }
}
func (this *LFUCache) Get(key int) int {
    if _, ok := this.value[key]; !ok {
        return -1
    }
    freq := this.count[key] // 获取key当前频次
    this.count[key] += 1  // 频次+1
    this.freq[freq].DelByKey(key) // 从lru的list中删掉，挪到另外的freq队列，加到队尾，删除方便。
    if this.min == freq&&this.freq[freq].lens == 0 {
        this.min += 1 // freq的删除key都需要触发是否更新min的操作
        delete(this.freq,freq)
    }
    if this.freq[freq + 1] == nil {
        this.freq[freq + 1] = &LinkList{}
    }
    this.freq[freq + 1].AddTail(key)
    return this.value[key]
}

func (this *LFUCache) Put(key int, value int) {
    if this.cap <= 0 {return}
    if _, ok := this.value[key]; ok {
        //如果存在就是更新key对应的value就ok了
        this.value[key] = value
        this.Get(key)
        return
    }
    // 看下是否超过cap需要删掉最小的min对应的key，value
    if len(this.value) >= this.cap {
        evit := this.freq[this.min].DelFront()
        delete(this.value, evit)
        if this.freq[this.min].Lens() <= 0 {
            delete(this.freq, this.min)
        }
    }
    // 插入新增元素,freq为1
    this.value[key] = value
    this.count[key] = 1
    this.min = 1
    if this.freq[1] == nil {
        this.freq[1] = &LinkList{}
    }
    this.freq[1].AddTail(key)
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */