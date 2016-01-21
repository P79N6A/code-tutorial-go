package main
/*
A first-level method is a method which DO NOT CALL *other methods in the class/struct/interface*

This is correct behavior for structs, because all struct-method calls are non-virtual and resolved at compile time

*/

import "fmt"

type parent struct {
}

func (_ parent) simple() {
	fmt.Println("parent simple called")
}
func (p parent) complex() {
	fmt.Print("in complex, calling simple:")
	p.simple()
}

type child struct {
	parent  // embeding
}
func (_ child) simple() {
	fmt.Println("child simple called")
}

func main() {
	var p = parent{}
	var c = child{}

	p.simple()
	p.complex()

	c.simple()
	// only call parent, so this can not express like polymorphism
	c.complex() // 4 in complex, calling simple: parent simple called //2nd level method was called on base-class context
}
