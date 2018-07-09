package main

import (
        "crypto/md5"
        "fmt"
        "encoding/hex"
)
func GetMD5Hash(text string) string {
        hasher := md5.New()
        hasher.Write([]byte(text))
        return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
        str := "fweihgorwj"
        fmt.Printf("%x", md5.Sum([]byte(str)))

        fmt.Sprintf(GetMD5Hash(str))
}
