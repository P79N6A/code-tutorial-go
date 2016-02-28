package main
import (
	"encoding/json"
	"fmt"
)

type Foo1 struct {
	number int
	title  string
}
type Foo2 struct {
	Number int
	Title  string
}
type Foo3 struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}
func main() {
	f1, _ := json.Marshal(Foo1{number: 1,title:"xx"})
	fmt.Println(string(f1))
	f2, _ := json.Marshal(Foo2{Number: 1,Title:"xx"})
	fmt.Println(string(f2))
	f3, _ := json.Marshal(Foo3{Number: 1,Title:"xx"})
	fmt.Println(string(f3))
}
