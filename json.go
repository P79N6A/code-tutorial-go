package main

import (
        "encoding/json"
        "fmt"
)
type IPA struct {
        IP  string `json:"ip"`
}
func main() {
        a := []IPA{
                IPA{
                        IP:"1.1.1.1",
                },
                IPA{
                        IP:"1.1.1.2",
                },
        }
        b,_ := json.Marshal(&a)
        fmt.Println(string(b))
}
