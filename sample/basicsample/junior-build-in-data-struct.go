/*
	array:In Go, an array is a numbered sequence of elements of a specific length.
	
	slice:Slices are a key data type in Go, giving a more powerful interface 
			to sequences than arrays.
    slice like vector... 可变数组长度
	map:
		Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).
	range:
		use for travel
		range iterates over of elements in a variety of data structures. 
		Let’s see how to use range with some of the data structures we’ve already learned.
		
*/
package main

import "fmt"
// ******* function first letter should be Upper
func Juniordata() {
	// ============== array ==============
	fmt.Println("=================== array ===================")
    var a [5]int
    fmt.Println("emp:", a)
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])
    fmt.Println("len:", len(a))
    // array int with init
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
    // =================== slice ===================
    fmt.Println("=================== slice ===================")
    /*
     Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
     To create an empty slice with non-zero length, use the builtin make.
     Here we make a slice of strings of length 3 (initially zero-valued).
     3 是初始长度, slice是可边长度的array
    */
    s := make([]string, 3)
    fmt.Println("emp:", s)
	// We can set and get just like with arrays.
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])
    fmt.Println("len:", len(s))
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

	// Slices can also be copy’d. Here we create an empty slice c
	// of the same length as s and copy into c from s.
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)
	// Slices support a “slice” operator with the syntax slice[low:high]. For example, this gets a slice of the elements s[2], s[3], and s[4].
    l := s[2:5]
    fmt.Println("sl1:", l)
	// This slices up to (but excluding) s[5].
    l = s[:5]
    fmt.Println("sl2:", l)
	// And this slices up from (and including) s[2].
    l = s[2:]
    fmt.Println("sl3:", l)
	// We can declare and initialize a variable for slice in a single line as well.
    ///////////////////////////
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)
	// Slices can be composed into multi-dimensional data structures.
	// The length of the inner slices can vary, unlike with multi-dimensional arrays.
    twoSD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoSD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoSD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoSD)
    fmt.Println("=================== map ===================")

	//To create an empty map, use the builtin make: make(map[key-type]val-type).
    m := make(map[string]int)
	//Set key/value pairs using typical name[key] = val syntax.
    m["k1"] = 7
    m["k2"] = 13
	//Printing a map with e.g. Println will show all of its key/value pairs.
    fmt.Println("map:", m)
	//Get a value for a key with name[key].
    v1 := m["k1"]
    fmt.Println("v1: ", v1)
	//The builtin len returns the number of key/value pairs when called on a map.
    fmt.Println("len:", len(m))
	//The builtin delete removes key/value pairs from a map.
    delete(m, "k2")
    fmt.Println("map:", m)
	//The optional second return value when getting a value from a map indicates if the key was present in the map. 
	//This can be used to disambiguate between missing keys and keys with zero values like 0 or "". 
	//Here we didn’t need the value itself, so we ignored it with the blank identifier _.
    _, prs := m["k2"]
    fmt.Println("prs:", prs)
	//You can also declare and initialize a new map in the same line with this syntax.
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)


    fmt.Println("=================== range ===================")

	// Here we use range to sum the numbers in a slice. Arrays work like this too.
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)
	// range on arrays and slices provides both the index and value for each entry. Above we didn’t need the index, so we ignored it with the blank identifier _. Sometimes we actually want the indexes though.
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
	// range on map iterates over key/value pairs.
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
	// range on strings iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself.
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}

func main() {
    Juniordata()
}