package main

import (
	"strconv"
	"fmt"
)
type Stringer struct{
}
func (s*Stringer)String() string {
	return "Hello world"
}
func do(v interface{}) string {
	switch u := v.(type) {
	case int:
		return strconv.Itoa(u*2) // u has type int

	case string:
		mid := len(u) / 2 // split - u has type string
		return u[mid:] + u[:mid] // join

	case Stringer: //*another (non-empty) interface*
		return u.String() //call via method dispatch
	}
	return "unknown"
}
func main() {
	fmt.Println(do(1))
	fmt.Println(do("xx"))
}
