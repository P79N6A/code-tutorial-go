package main

import (
        "github.com/fsnotify/fsnotify"
        "fmt"
)

func ExampleNewWatcher() {
        watcher, err := fsnotify.NewWatcher()
        if err != nil {
                fmt.Errorf("%v",err)
                return
        }
        //defer watcher.Close()

        //done := make(chan bool)
        go func() {
                for {
                        select {
                        case event := <-watcher.Events:
                                fmt.Println("event:", event)
                                if event.Op&fsnotify.Write == fsnotify.Write {
                                        fmt.Println("modified file:", event.Name)
                                }
                        case err := <-watcher.Errors:
                                fmt.Println("error:", err)
                        }
                }
        }()

        err = watcher.Add("/Users/gaolichuang/workspace/go/src/groot_keeper/logs")
        if err != nil {
                fmt.Errorf("%v",err)
        }
        //<-done

}
func main() {
        ExampleNewWatcher()
        for {

        }
}
