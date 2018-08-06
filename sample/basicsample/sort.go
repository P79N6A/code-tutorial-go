/*
	Sorting:
		Go’s sort package implements sorting for builtins and user-defined types.
		We’ll look at sorting for builtins first.
	Sorting by Functions:
		Sometimes we’ll want to sort a collection by something other than its natural order.
		For example, suppose we wanted to sort strings by their length instead of alphabetically.
		Here’s an example of custom sorts in Go.
		
*/
package bsample


import "fmt"
import "sort"
/*
$ go run sorting.go
Strings: [a b c]
Ints:    [2 4 7]
Sorted:  true
*/
func Sorting() {
	// Sort methods are specific to the builtin type;
	// here’s an example for strings.
	// Note that sorting is in-place,
	// so it changes the given slice and doesn’t return a new one.
    strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)
	// An example of sorting ints.
    ints := []int{7, 2, 4}
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)
	// We can also use sort to check if a slice is already in sorted order.
    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)
}





// In order to sort by a custom function in Go, we need a corresponding type.
// Here we’ve created a ByLength type that is just an alias for the builtin []string type.
type ByLength []string
// We implement sort.Interface - Len, Less, and Swap - on our type
// so we can use the sort package’s generic Sort function.
// Len and Swap will usually be similar across types and
// Less will hold the actual custom sorting logic.
// In our case we want to sort in order of increasing string length,
// so we use len(s[i]) and len(s[j]) here.
func (s ByLength) Len() int {
    return len(s)
}
func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}
//With all of this in place,
// we can now implement our custom sort by casting the original fruits slice to ByLength,
// and then use sort.Sort on that typed slice.

/*
Running our program shows a list sorted by string length, as desired.
$ go run sorting-by-functions.go 
[kiwi peach banana]
By following this same pattern of creating a custom type,
 implementing the three Interface methods on that type, and then calling sort.
 Sort on a collection of that custom type,
 we can sort Go slices by arbitrary functions.
*/
func SortByFunc() {
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(ByLength(fruits))
    fmt.Println(fruits)
}
