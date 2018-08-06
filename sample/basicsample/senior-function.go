/*

	1.Go supports anonymous functions, which can form closures.
		Anonymous functions are useful when you want to define a function inline 
		without having to name it.
	2.Go supports recursive functions. Here’s a classic factorial example.
	
*/
package bsample

import (
	"fmt"
)

//	This function intSeq returns another function, 
//	which we define anonymously in the body of intSeq. 
//	The returned function closes over the variable i to form a closure.
func intSeq() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

// This fact function calls itself until it reaches the base case of fact(0).
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func Seniorfunc() {
	// We call intSeq, assigning the result (a function) to nextInt.
	// This function value captures its own i value,
	// which will be updated each time we call nextInt.
    nextInt := intSeq()
	// See the effect of the closure by calling nextInt a few times.
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())
	// To confirm that the state is unique to that particular function,
	// create and test a new one.
    newInts := intSeq()
    fmt.Println(newInts())
    
    fmt.Println(fact(7))
}