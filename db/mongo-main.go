package main

import (
    "mustard/internal/gopkg.in/mgo.v2"
    "mustard/internal/gopkg.in/mgo.v2/bson"
    "fmt"
    "log"
)


type Person struct {
    Name string
    Phone string
}

func main() {
    session, err := mgo.Dial("127.0.0.1")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

    c := session.DB("test").C("people")
    err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
            &Person{"Cla", "+55 53 8402 8510"})
    if err != nil {
        log.Fatal(err)
    }

    result := Person{}
    err = c.Find(bson.M{"name": "Ale"}).One(&result)
    c.Find(nil).Batch(4).All()
    i := c.Find(nil).Iter()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Phone:", result.Phone)
}

