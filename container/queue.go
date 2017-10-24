package main

import "fmt"

// NewStack returns a new stack.
func NewStack() *Stack {
        return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
        nodes []interface{}
        count int
}

// Push adds a node to the stack.
func (s *Stack) Push(i interface{}) {
        s.nodes = append(s.nodes[:s.count], i)
        s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() interface{} {
        if s.count == 0 {
                return nil
        }
        s.count--
        return s.nodes[s.count]
}

// NewQueue returns a new queue with the given initial size.
func NewQueue(size int) *Queue {
        return &Queue{
                nodes: make([]interface{}, size),
                size:  size,
        }
}

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue struct {
        nodes []interface{}
        size  int
        head  int
        tail  int
        count int
}

// Push adds a node to the queue.
func (q *Queue) Push(i interface{}) {
        if q.head == q.tail && q.count > 0 {
                nodes := make([]interface{}, len(q.nodes)+q.size)
                copy(nodes, q.nodes[q.head:])
                copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
                q.head = 0
                q.tail = len(q.nodes)
                q.nodes = nodes
        }
        q.nodes[q.tail] = i
        q.tail = (q.tail + 1) % len(q.nodes)
        q.count++
}

// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() interface{} {
        if q.count == 0 {
                return nil
        }
        node := q.nodes[q.head]
        q.head = (q.head + 1) % len(q.nodes)
        q.count--
        return node
}
func main() {
        s := NewStack()
        s.Push(1)
        s.Push("xxx")
        s.Push("efw")
        fmt.Println(s.Pop(), s.Pop(), s.Pop())

        q := NewQueue(1)
        q.Push("xxx")
        q.Push("ww")
        q.Push(22)
        fmt.Println(q.Pop(), q.Pop(), q.Pop())
}