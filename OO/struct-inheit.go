package main
import (
	"fmt"
	"unsafe"
)

// struct embedding
type Person struct {
	name    string
	age     int
}
func (p *Person)ShowCard() {
	fmt.Printf("%s, %i",p.name,p.age)
}
type Man struct {
	Person
	some    string
}
type WoMan struct {
	Person
}
func show(p *Person) {
	p.ShowCard()
}
func main() {
	w := WoMan{Person{"alice",29}}
	show((*Person)(unsafe.Pointer(&w)))
}
