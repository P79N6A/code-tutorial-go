package main

import "fmt"

/*
A second-level method is a method which DO CALL *other methods in the class/struct/interface*


This is correct behavior for interfaces because all interface-method calls are virtual so they're resolved at run time.

*/
// interface
// A golang-Interface is a class with no fields and ONLY VIRTUAL methods.
type sampler interface {
	simple()
}

func complex(i sampler) {
	fmt.Print("in complex, calling simple:")
	i.simple()
}

type parent struct {
}

func (_ *parent) simple() {
	fmt.Println("parent simple called")
}

type child struct {
	parent
}

func (_ *child) simple() {
	fmt.Println("child simple called")
}

func main() {
	var p = parent{}
	var c = child{}

	p.simple()
	complex(&p)

	c.simple()
	complex(&c)
}

