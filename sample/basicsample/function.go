/*
	Functions are central in Go. 
	1.In function, Go requires explicit returns
	2.Function allow multi return value
		Go has built-in support for multiple return values. 
		This feature is used often in idiomatic Go, for example to return 
		both result and error values from a function.-
	3.Variadic Functions
		Variadic functions can be called with any number of trailing arguments. 
		For example, fmt.Println is a common variadic function.
*/
package bsample

import (
	"fmt"
)
//Here’s a function that takes two ints and returns their sum as an int.
func plus(a int, b int) int {
	// Go requires explicit returns,
	// i.e. it won’t automatically return the value of the last expression.
    return a + b
}
// When you have multiple consecutive parameters of the same type,
// you may omit the type name for the like-typed parameters up to
// the final parameter that declares the type.
func plusPlus(a, b, c int) int {
    return a + b + c
}

func vals() (int, int) {
    return 3, 7
}

// Here’s a function that will take an arbitrary number of ints as arguments.
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}
func Function() {
    res := plus(1, 2)
    fmt.Println("1+2 =", res)
    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
    // multi return
    fmt.Println("=================== Multi return value ===================")
    a, b := vals()
    fmt.Println(a,b)
	// If you only want a subset of the returned values, use the blank identifier _.
    _, c := vals()
    fmt.Println(c)
    fmt.Println("=================== Variadic Functions ===================")
    sum(1, 2)
    sum(1, 2, 3)
	// If you already have multiple args in a slice, 
	// apply them to a variadic function using func(slice...) like this.
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}