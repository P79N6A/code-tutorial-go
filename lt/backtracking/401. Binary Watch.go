package main

import "fmt"

/*

A binary watch has 4 LEDs on the top which represent the hours (0-11), and the 6 LEDs on the bottom represent the minutes (0-59).

Each LED represents a zero or one, with the least significant bit on the right.


For example, the above binary watch reads "3:25".

Given a non-negative integer n which represents the number of LEDs that are currently on, return all possible times the watch could represent.

Example:

Input: n = 1
Return: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]
Note:
The order of output does not matter.
The hour must not contain a leading zero, for example "01:00" is not valid, it should be "1:00".
The minute must be consist of two digits and may contain a leading zero, for example "10:2" is not valid, it should be "10:02".

*/
func readBinaryWatch(num int) []string {
    ret := make([]string,0)
    for i:=0;i<=num;i++ {
        h := gh(i)
        m := gm(num-i)
        for _,hh := range h {
            if hh > 11 {
                continue
            }
            for _,mm := range m {
                if mm > 59 {
                    continue
                }
                ret = append(ret,fmt.Sprintf("%d:%02d",hh,mm))
            }
        }
    }
    return ret
}
func gm(num int) []int {
    ret := make([]int,0)
    bk([]int{32,16,8,4,2,1},num,0,0,&ret)
    return ret
}
func gh(num int) []int {
    ret := make([]int,0)
    bk([]int{8,4,2,1},num,0,0,&ret)
    return ret
}
func bk(data []int, cnt int, pos int, sum int, ret *[]int) {
    if cnt == 0 {
        *ret = append(*ret,sum)
        return
    }
    for i:=pos;i<len(data);i++ {
        bk(data,cnt-1,i+1,sum+data[i],ret)
    }
}


func main() {
    //ret := make([]int,0)
    //bk([]int{8,4,2,1},2,0,0,&ret)
    //fmt.Print(ret)
    fmt.Println(readBinaryWatch(2))
    fmt.Println(readBinaryWatch(0))

}
