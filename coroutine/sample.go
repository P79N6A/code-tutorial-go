package main

import (
        "time"
        "fmt"
)

type OBj struct {
        Name string
}
type Container struct {
        OBJ *OBj
}

func worker(obj *OBj) {
        for {
                time.Sleep(time.Second)
                fmt.Println(obj.Name)
        }
}
func main() {
        o := Container{
                OBJ:&OBj{
                        Name:"XXX",
                },
        }
        go worker(o.OBJ)
        time.Sleep(time.Second)
        o.OBJ = nil
        time.Sleep(time.Second * 20)
}
