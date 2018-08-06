package main

func AppendSlice(a *[]string) {
    *a = append(*a, "XXX")
}

func main() {
    a := []string{"1"}
    AppendSlice(&a)
    for _,v := range a {
        println(v)
    }
}
