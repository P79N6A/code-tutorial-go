package main
import (
	"encoding/gob"
	"encoding/binary"
	"bytes"
	"fmt"
)
type Base struct {
	message string
}

type Items map[string]int

type Obj struct {
	name    string
	info    string
	keys    Items
	base    Base
}
func main() {
	m := Obj{"xx","22",Items{"x":1},Base{message:"wef"}}
	b := new(bytes.Buffer)
	fmt.Println(b.Len())
	e := gob.NewEncoder(b)

	fmt.Println(b.Len())
	// Encoding the map
	err := e.Encode(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(b.Len())

	var decodedMap Obj
	d := gob.NewDecoder(b)

	// Decoding the serialized data
	err = d.Decode(&decodedMap)
	if err != nil {
		panic(err)
	}

	// Ta da! It is a map!
	fmt.Printf("%#v\n", decodedMap)
}
