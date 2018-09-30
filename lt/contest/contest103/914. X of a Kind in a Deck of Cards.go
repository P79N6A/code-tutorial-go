package contest103

import (
    "fmt"
    "sort"
)

func main() {
    a := []string{"a","aa","b"}
    sort.Strings(a)
    fmt.Println(a)
    sort.Slice(a, func(i, j int) bool {
        return a[i]<a[j]
    })

    fmt.Println(gcd(0,5))
}

func gcd(A, B int) int {
    if (A < B) {A, B = B, A}
    if (A%B == 0) {
        return B
    } else {
        return gcd(B, A%B)
    }
}
