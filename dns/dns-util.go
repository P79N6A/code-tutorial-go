package main
import (
        "gcodebase/hash"
        "path/filepath"
        "fmt"
)
func GenWriteFileTemporaryName(name string) string {
        dir := filepath.Dir(name)
        base := filepath.Base(name)
        subfix := hash.GenerateKey(5)
        newBase := "." + base + "." + subfix
        return filepath.Join(dir,newBase)
}
func main() {
        str := GenWriteFileTemporaryName("./../xxx.log")
        fmt.Println(str)
}
