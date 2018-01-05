package main

import (
        "fmt"
        "reflect"
)

// Definition of Map function
type mapFunc func(interface{}) interface{}

// Map maps the function onto the array
func Map(fn mapFunc, array interface{}) []interface{} {
        val := reflect.ValueOf(array)
        outputArray := make([]interface{}, val.Len())
        for i := 0; i < val.Len(); i++ {
                outputArray[i] = fn(val.Index(i).Interface())
        }
        return outputArray
}

func main() {
        inputList := []int{1, 2, 3}

        square := func(x interface{}) interface{} {
                return x.(int) * x.(int)
        }

        expected := []int{1, 4, 9}
        actual := Map(square, inputList)

        for i, e := range expected {
                if e != actual[i] {
                        fmt.Println("XXX")
                }
        }
}
