package main

import (
        "bytes"
        "io"
        "os"
        "os/exec"
)


func f1() {
        c1 := exec.Command("ls")
        c2 := exec.Command("wc", "-l")

        r, w := io.Pipe()
        c1.Stdout = w
        c2.Stdin = r

        var b2 bytes.Buffer
        c2.Stdout = &b2

        c1.Start()
        c2.Start()
        c1.Wait()
        w.Close()
        c2.Wait()
        io.Copy(os.Stdout, &b2)
}
func main() {
        c1 := exec.Command("ls")
        c2 := exec.Command("wc", "-l")
        c2.Stdin, _ = c1.StdoutPipe()
        c2.Stdout = os.Stdout
        _ = c2.Start()
        _ = c1.Run()
        _ = c2.Wait()
}