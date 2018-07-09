package main

import (
        "crypto/md5"
        "fmt"
        "encoding/hex"
        "crypto/des"
        "encoding/base32"
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
        str := "this is a exampiile"
        fmt.Println(getbase32(str))
 //       fmt.Printf("%x", md5.Sum([]byte(str)))

//        fmt.Sprintf(GetMD5Hash(str))
}

func getbase32(str string) string {
        return  base32.HexEncoding.EncodeToString([]byte(str))
}
func aesEn() {
        des.BlockSize
}
