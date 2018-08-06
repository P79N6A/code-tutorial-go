/*
	Go supports pointers, allowing you to pass references to values
	and records within your program.
	Go pointer like C/C++

	Go’s structs are typed collections of fields.
	They’re useful for grouping data together to form records.
	
	Struct pointer could use dots!!!
	
*/
package bsample

import (
	"fmt"
)

/*
	We’ll show how pointers work in contrast to values with 2 functions: 
	zeroval and zeroptr. zeroval has an int parameter,
	so arguments will be passed to it by value.
	zeroval will get a copy of ival distinct from the one in the calling function.
*/
func zeroval(ival int) {
    ival = 0
}
// zeroptr in contrast has an *int parameter, meaning that it takes an int pointer.
// The *iptr code in the function body then dereferences the pointer from its memory
// address to the current value at that address. Assigning a value to
// a dereferenced pointer changes the value at the referenced address.
func zeroptr(iptr *int) {
    *iptr = 0
}



/////////struct support///////////////////

type person struct {
    name string
    age  int
}


	
func PointAndStruct() {
    fmt.Println("=================== pointer ===================")
    i := 1
    fmt.Println("initial:", i)
    zeroval(i)
    fmt.Println("zeroval:", i)
	// The &i syntax gives the memory address of i, i.e. a pointer to i.
    zeroptr(&i)
    // zeroval doesn’t change the i in main,
    // but zeroptr does because it has a reference to the memory address for that variable.
    fmt.Println("zeroptr:", i)
	// Pointers can be printed too.
    fmt.Println("pointer:", &i)


    fmt.Println("=================== pointer ===================")
    // This syntax creates a new struct.
    fmt.Println(person{"Bob", 20})
	// You can name the fields when initializing a struct.
    fmt.Println(person{name: "Alice", age: 30})
    // Omitted fields will be zero-valued.
    fmt.Println(person{name: "Fred"})
    // An & prefix yields a pointer to the struct.
    fmt.Println(&person{name: "Ann", age: 40})
	// Access struct fields with a dot.
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)
	// You can also use dots with struct pointers - the pointers are automatically dereferenced.
    sp := &s // sp is a reference
    sa := s // sa is a copy
    fmt.Println(sp.age)
	// Structs are mutable.
    sp.age = 51
    fmt.Println(sp.age)
    fmt.Println(sa.age)
}
