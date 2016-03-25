package main

import "fmt"
/*
Reference type
	slices
	maps
	channels
	pointers
	functions
Not Reference type:
	array
}:

*/
func modifyArray(array [5]int) {
	for i,_ := range array {
		array[i] += 1
	}
}
func modifySlice(slice []int) {
	for i,_ := range slice {
		slice[i] += 1
	}
}
func modifyMap(m map[int]int) {
	for k,_ := range m {
		m[k] += 1
	}
}
func main() {
	// array is not reference
	array := [5]int{1,2,3,4,5}
	modifyArray(array)
	fmt.Println(array)
	// slice
	slice := []int{1,2,3,4,5}
	modifySlice(slice)
	fmt.Println(slice)
	// map
	m := map[int]int{1:1,2:3,3:4}
	modifyMap(m)
	fmt.Println(m)
}
