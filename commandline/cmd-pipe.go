package main

import (
        "bytes"
        "io"
        "os/exec"
        "time"
        "fmt"
        "syscall"
)

func Execute(output_buffer *bytes.Buffer,timeout time.Duration, stack ...*exec.Cmd) (err error) {
        var error_buffer bytes.Buffer
        pipe_stack := make([]*io.PipeWriter, len(stack)-1)
        i := 0
        for ; i < len(stack)-1; i++ {
                stdin_pipe, stdout_pipe := io.Pipe()
                stack[i].Stdout = stdout_pipe
                stack[i].Stderr = &error_buffer
                stack[i+1].Stdin = stdin_pipe
                pipe_stack[i] = stdout_pipe
        }
        stack[i].Stdout = output_buffer
        stack[i].Stderr = &error_buffer

        if err := call(stack, pipe_stack,timeout); err != nil {
                fmt.Printf("%s,%v",string(error_buffer.Bytes()), err)
        }
        return err
}

func call(stack []*exec.Cmd, pipes []*io.PipeWriter,timeout time.Duration) (err error) {
        if stack[0].Process == nil {
                if err = stack[0].Start(); err != nil {
                        return err
                }
        }
        if len(stack) > 1 {
                if err = stack[1].Start(); err != nil {
                        return err
                }
                defer func() {
                        if err == nil {
                                pipes[0].Close()
                                err = call(stack[1:], pipes[1:],timeout)
                        }
                }()
        }
        done := make(chan error)
        go func() {
                done <- stack[0].Wait()
        }()

        select {
        case <-time.After(timeout):
        // timeout
                if errr := stack[0].Process.Kill(); errr != nil {
                        err =  fmt.Errorf("failed to kill: %s, error: %s", stack[0].Path, err)
                        return
                }
                go func() {
                        <-done // allow goroutine to exit
                }()
                err = fmt.Errorf("TimeOut %s",stack[0].Path)
                return
        case err = <-done:
                if exiterr, ok := err.(*exec.ExitError); ok {
                        if status, ok := exiterr.ProcessState.Sys().(syscall.WaitStatus); ok {
                                fmt.Printf("%d",status.ExitStatus())
                        }
                }
        }
        return
}
func RunCommandsWithTimeOut(timeout time.Duration,cmds ...[]string) (err error, stdout []byte) {
        var b bytes.Buffer
        execs := make([]*exec.Cmd,0)
        for _,cmd := range cmds {
                execs = append(execs, exec.Command(cmd[0],cmd[1:]...))
        }
        err = Execute(&b, timeout, execs...)
        if err != nil {
                return
        }
        return nil, b.Bytes()
}
func main() {
        err,out := RunCommandsWithTimeOut(time.Second,[]string{"ls","/Users/gaolichuang/workspace/go/src/glory_dns"},
        []string{"grep","M"},
        []string{"sort","-r"})
        fmt.Println(err)
        fmt.Println(string(out))
}