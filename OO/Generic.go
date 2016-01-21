package main

import "fmt"
// http://www.weberc2.com/posts/2014/12/12/living-without-generics.txt
// http://cholerae.com/2015/01/02/%E3%80%90%E7%BF%BB%E8%AF%91%E3%80%91%E7%94%9F%E6%B4%BB%E5%9C%A8%E6%B2%A1%E6%9C%89%E6%B3%9B%E5%9E%8B%E7%9A%84Go%E8%AF%AD%E8%A8%80%E4%B8%96%E7%95%8C%E9%87%8C/
type Stack struct {
	items     []interface{}
	lastIndex int
}

func NewStack() *Stack {
	return &Stack{items: make([]interface{}, 0)}
}

// push `v` onto the stack
func (s *Stack) Push(v interface{}) {
	// if we're below capacity, add the new element to `s.lastIndex`
	// otherwise, append onto the end
	if s.lastIndex < len(s.items) {
		s.items[s.lastIndex] = v
	} else {
		s.items = append(s.items, v)
	}
	s.lastIndex++
}

// pull the top element off the stack
func (s *Stack) Pull() interface{} {
	s.lastIndex--
	return s.items[s.lastIndex]
}

func (s *Stack) Size() int {
	return s.lastIndex
}

////////////////////////
type IntStack struct {
	*Stack
}

func NewIntStack() *IntStack {
	return &IntStack{NewStack()}
}

func (s *IntStack) Push(i int) {
	s.Stack.Push(i)
}

func (s *IntStack) Pull() int {
	// pull the top element from the stack and type assert
	// it back into an int
	return s.Stack.Pull().(int)
}

/////////////
