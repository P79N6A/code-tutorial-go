package main

import (
        "fmt"
        "path/filepath"
)
func main() {
        fmt.Println(filepath.Abs("."))
        fmt.Println(filepath.Abs("/etc"))
        fmt.Println(filepath.Abs("../../"))
        fmt.Println(filepath.Join("","bb"))
}
