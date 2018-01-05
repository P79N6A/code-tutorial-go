package main

import (
        "fmt"
        "reflect"
)

// Definition of Filter function
type filterFunc func(interface{}) bool

// Filter filters the array based on the predicate
func Filter(fn filterFunc, array interface{}) []interface{} {
        val := reflect.ValueOf(array)
        var outputArray []interface{}
        for i := 0; i < val.Len(); i++ {
                if fn(val.Index(i).Interface()) {
                        outputArray = append(outputArray, val.Index(i).Interface())
                }
        }
        return outputArray
}
func main() {
        inputList := []int{1, 2, 3}

        modulotwo := func(x interface{}) bool {
                return x.(int)%2 == 0
        }

        expected := []int{2}
        actual := Filter(modulotwo, inputList)

        for i, e := range expected {
                if e != actual[i] {
                        fmt.Println("xxxxx")
                }
        }
}

