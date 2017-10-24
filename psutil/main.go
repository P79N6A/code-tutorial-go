package main

import (
        "github.com/shirou/gopsutil/cpu"
        "github.com/shirou/gopsutil/mem"
        "github.com/shirou/gopsutil/load"
        "github.com/shirou/gopsutil/process"
        "fmt"
        "os"
)
func main() {
        info,_ := cpu.Info()
        fmt.Println(info)
        vm,_ := mem.VirtualMemory()
        fmt.Println(vm)

        avg,_ :=load.Avg()
        fmt.Println(avg)

        pro,_:=process.NewProcess(int32(os.Getpid()))
        mi,_ := pro.MemoryInfo()
        fmt.Println(mi)
        fmt.Println(pro.MemoryPercent())
        tm,_ := pro.Times()
        fmt.Println(tm)
}
