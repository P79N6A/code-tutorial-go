package main

import (
    "github.com/hashicorp/memberlist"
    "fmt"
)

func main() {
    list, err := memberlist.Create(memberlist.DefaultLocalConfig())
    if err != nil {
        panic("Failed to create memberlist: " + err.Error())
    }

    // Join an existing cluster by specifying at least one known member.
    n, err := list.Join([]string{"1.2.3.4"})
    if err != nil {
        panic("Failed to join cluster: " + err.Error())
    }
    fmt.Println(n)

    // Ask for members of the cluster
    for _, member := range list.Members() {
        fmt.Printf("Member: %s %s\n", member.Name, member.Addr)
    }
    
}
