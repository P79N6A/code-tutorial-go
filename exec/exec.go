package main

import (
        "os/exec"
        "time"
        "fmt"
        "syscall"
        "bytes"
)

func RunCombineCommandWithTimeOut(cmds []string, timeout time.Duration) (err error,returnCode int, stderrout []byte) {
        if len(cmds) <= 0 {
                err = fmt.Errorf("Invalid Command %v",cmds)
                return
        }
        cmd := exec.Command(cmds[0],cmds[1:]...)
        done := make(chan error)
        go func() {
                stderrout,err = cmd.CombinedOutput()
                done <- err
        }()

        select {
        case <-time.After(timeout):
        // timeout
                if errr := cmd.Process.Kill(); errr != nil {
                        err =  fmt.Errorf("failed to kill: %s, error: %s", cmd.Path, err)
                        return
                }
                go func() {
                        <-done // allow goroutine to exit
                }()
                err = fmt.Errorf("TimeOut %s",cmd.Path)
                return
        case err = <-done:
                if exiterr, ok := err.(*exec.ExitError); ok {
                        if status, ok := exiterr.ProcessState.Sys().(syscall.WaitStatus); ok {
                                returnCode = status.ExitStatus()

                        }
                }
        }
        return
}
func RunCommandWithTimeOut(cmds []string, timeout time.Duration) (err error,returnCode int, stdout []byte,stderr []byte) {
        if len(cmds) <= 0 {
                err = fmt.Errorf("Invalid Command %v",cmds)
                return
        }
        cmd := exec.Command(cmds[0],cmds[1:]...)
        cmdStdOut := &bytes.Buffer{}
        cmdStdErr := &bytes.Buffer{}
        cmd.Stdout = cmdStdOut
        cmd.Stderr = cmdStdErr

        cmd.Start()
        done := make(chan error)
        go func() {
                done <- cmd.Wait()
        }()

        select {
        case <-time.After(timeout):
        // timeout
                if errr := cmd.Process.Kill(); errr != nil {
                        err =  fmt.Errorf("failed to kill: %s, error: %s", cmd.Path, err)
                        return
                }
                go func() {
                        <-done // allow goroutine to exit
                }()
                err = fmt.Errorf("TimeOut %s",cmd.Path)
                return
        case err = <-done:
                stdout = cmdStdOut.Bytes()
                stderr = cmdStdErr.Bytes()
                if exiterr, ok := err.(*exec.ExitError); ok {
                        if status, ok := exiterr.ProcessState.Sys().(syscall.WaitStatus); ok {
                                returnCode = status.ExitStatus()

                        }
                }
        }
        return
}
func RunTwoCommandWithTimeOut(cmds1 []string, cmds2 []string, timeout time.Duration) (err error,returnCode int, stdout []byte,stderr []byte) {
        if len(cmds1) <= 0 || len(cmds2) <= 0 {
                err = fmt.Errorf("Invalid Command %v, %v",cmds1,cmds2)
                return
        }


        cmdStdOut := &bytes.Buffer{}
        cmdStdErr := &bytes.Buffer{}
        cmd1 := exec.Command(cmds1[0],cmds1[1:]...)
        cmd2 := exec.Command(cmds2[0],cmds2[1:]...)
        cmd2.Stdin,_ = cmd1.StdoutPipe()
        cmd2.Stdout = cmdStdOut
        cmd2.Stderr = cmdStdErr

        cmd2.Start()
        cmd1.Start()
        done := make(chan error)
        go func() {
                done <- cmd2.Wait()
        }()

        select {
        case <-time.After(timeout):
        // timeout
                if errr := cmd2.Process.Kill(); errr != nil {
                        err =  fmt.Errorf("failed to kill: %s, error: %s", cmd2.Path, err)
                        return
                }
                go func() {
                        <-done // allow goroutine to exit
                }()
                err = fmt.Errorf("TimeOut %s",cmd2.Path)
                return
        case err = <-done:
                stdout = cmdStdOut.Bytes()
                stderr = cmdStdErr.Bytes()
                if exiterr, ok := err.(*exec.ExitError); ok {
                        if status, ok := exiterr.ProcessState.Sys().(syscall.WaitStatus); ok {
                                returnCode = status.ExitStatus()

                        }
                }
        }
        return
}
func main(){
        /*
        cmd := exec.Command("ls", "-l","/")
        e,code,so,se := RunCommandWithTimeOut(cmd,time.Second*3)

        fmt.Println("XXXXXX")
        fmt.Println(e)
        fmt.Println("XXXXXX")
        fmt.Println(code)
        fmt.Println("XXXXXX")
        fmt.Println(string(so))
        fmt.Println("XXXXXX")
        fmt.Println(string(se))
        fmt.Println("XXXXXX")
        cmd1 := exec.Command("ls", "-l","/")
        e,code,so = RunCombineCommandWithTimeOut(cmd1,time.Second*3)
        fmt.Println("XXXXXX")
        fmt.Println(e)
        fmt.Println("XXXXXX")
        fmt.Println(code)
        fmt.Println("XXXXXX")
        fmt.Println(string(so))
        fmt.Println("XXXXXX")
        */
        a,b,c,d := RunTwoCommandWithTimeOut([]string{"ls"},[]string{"sort"},time.Second)
        fmt.Println(a)
        fmt.Println(b)
        fmt.Println(string(c))
        fmt.Println(d)
}


