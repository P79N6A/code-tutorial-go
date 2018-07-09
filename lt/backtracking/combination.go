package main


import (
    "fmt"
    "sort"
)


// 连续的长度大于1的所有子集
func bk3(data []int, pos int,visit *[]bool, res *[]int, ret *[][]int) {
    sort.Ints(data)
    if len(*res) > 1 {
        n := make([]int,len(*res))
        copy(n,*res)
        *ret = append(*ret,n)
    }
    for i:=pos;i<len(data);i++ {
        if pos == 0 || (*visit)[i-1]==true {
            (*visit)[i] = true
            *res = append(*res, data[i])
            bk3(data,i+1,visit, res, ret)
            *res = (*res)[:len(*res)-1]
            (*visit)[i] = false
        }
    }
}
// 从集合中拿cnt个数的子集
func bk2(data []int, cnt int, pos int, res *[]int, ret *[][]int) {
    if cnt == 0 {
        n := make([]int,len(*res))
        copy(n,*res)
        *ret = append(*ret,n)
        return
    }
    for i:=pos;i<len(data);i++ {
        *res = append(*res,data[i])
        // 和perm区别在于下次传入的是i+1.
        //这层递归按照顺序取i，不取i取i+1, 不取i,i+1,取i+2 。。。。
        bk2(data,cnt-1,i+1,res,ret)
        *res = (*res)[:len(*res)-1]
    }
}
// 从集合中拿cnt个数的子集；另外一种写法
func bk4(data []int, cnt int, pos int, res *[]int, ret *[][]int) {
    if cnt == 0 {
        n := make([]int,len(*res))
        copy(n,*res)
        *ret = append(*ret,n)
        return
    }
    if pos >= len(data) {
        return
    }
    // 本层将pos的决策做完。取或者不取
    // 通过控制取或者不取可以得到不同的结果顺序
    bk4(data,cnt,pos+1,res,ret) // 不取决策
    *res = append(*res,data[pos])
    bk4(data,cnt-1,pos+1,res,ret)// 取的决策
    *res = (*res)[:len(*res)-1]
}


func main() {
    data := []int{16,8,4,2,1}
    ret := make([][]int,0)
    res := make([]int,0)
    visit := make([]bool,len(data))
    bk3(data,0,&visit,&res,&ret)
    fmt.Println(ret)

    ret = make([][]int,0)
    res = make([]int,0)
    bk2(data,2,0,&res,&ret)
    fmt.Println(ret)

    ret = make([][]int,0)
    res = make([]int,0)
    bk4(data,2,0,&res,&ret)
    fmt.Println(ret)

    data1 := []int{1,2,3,3,4}
    ret = make([][]int,0)
    res = make([]int,0)
    bk4(data1,2,0,&res,&ret)
    fmt.Println(ret)

}
