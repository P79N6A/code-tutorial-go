package main

import (
        "encoding/json"
        "fmt"
)

type A struct {
        Name string
        Sex int32
}
type AA struct {
        A1      A      `json:"a1,omitempty"`
        A2      A      `json:"a2,omitempty"`
        R1       string
        R2       string
}

func main() {
        aa := AA{
                A1:A{
                        Name:"Alice",
                        Sex:1,
                },
                R1:"RRR",
        }
        jsonstr,_ := json.Marshal(&aa)
        fmt.Println(string(jsonstr))
}
