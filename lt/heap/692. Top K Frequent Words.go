package main

import (
    "container/heap"
    "fmt"
)

/*
Given a non-empty list of words, return the k most frequent elements.

Your answer should be sorted by frequency from highest to lowest. If two words have the same frequency, then the word with the lower alphabetical order comes first.

Example 1:
Input: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
Output: ["i", "love"]
Explanation: "i" and "love" are the two most frequent words.
    Note that "i" comes before "love" due to a lower alphabetical order.
Example 2:
Input: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
Output: ["the", "is", "sunny", "day"]
Explanation: "the", "is", "sunny" and "day" are the four most frequent words,
    with the number of occurrence being 4, 3, 2 and 1 respectively.
Note:
You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
Input words contain only lowercase letters.
Follow up:
Try to solve it in O(n log k) time and O(n) extra space.
*/

func main() {
    fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"},2))
    fmt.Println(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4))

}

type item struct {
    key string
    num int
}
type ItemHeap []item
func (h ItemHeap) Len() int      { return len(h) }
func (h ItemHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h ItemHeap) Less(i, j int) bool {
    if h[i].num == h[j].num {
        return h[i].key > h[j].key // 如果个数一样，则字典序排在前边，反过来了
    }
    return h[i].num < h[j].num // 小顶堆
}
func (h *ItemHeap) Push(x interface{}) {
    *h = append(*h, x.(item))
}
func (h *ItemHeap) Pop() interface{} {
    old := *h
    n := len(old)
    item := old[n-1]
    *h = old[0 : n-1]
    return item
}

func topKFrequent(words []string, K int) []string {
    set := make(map[string]int)
    for _, w := range words {
        set[w] += 1
    }
    h := &ItemHeap{}
    heap.Init(h)
    for word, num := range set {
        heap.Push(h, item{word,num})
        if h.Len() > K {
            heap.Pop(h)
        }
    }
    ret := make([]string, h.Len())
    for i:=len(ret)-1;i>=0;i-- {
        // 注意这个地方是heap.Pop(),这样才是取的队头，并且将0和n-1的元素交换
        // 如果是h.Pop() 则是没有经过swap的队尾了。
        ret[i]=heap.Pop(h).(item).key
    }
    return ret
}
