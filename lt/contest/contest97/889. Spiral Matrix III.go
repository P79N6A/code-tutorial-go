package main

import "fmt"

func main() {
        fmt.Println(spiralMatrixIII(1,4,0,0))
        fmt.Println(spiralMatrixIII(5,6,1,4))
        ret := make([][]int,0)
        //circle(1,4,5,6,1,&ret)
        circle(0,3,5,6,3,&ret)
        fmt.Println(ret)
}

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
        ret := make([][]int,0)
        rr,cc := r0,c0
        for i:=1;;i+=2 {
                circle(rr,cc,R,C,i,&ret)
                if len(ret) >= R*C {
                        return ret
                }
                rr,cc = rr-1,cc-1
        }
        return ret
}
func circle(r,c int, R,C int, step int,ret *[][]int) {
        //fmt.Println(r,c,R*C,step,*ret)
        rr,cc := r,c
        for i:=0;i<step;i++ {
         //       fmt.Println("X",rr,cc)
                if rr >=0&&rr<R&&cc>=0&&cc<C {
                        if len(*ret) < R*C {
                                *ret = append(*ret, []int{rr, cc})
                        }
                }
                cc = cc+1
        }
        for i:=0;i<step;i++ {
          //      fmt.Println("XX",rr,cc)
                if rr >=0&&rr<R&&cc>=0&&cc<C {
                        if len(*ret) < R*C {
                                *ret = append(*ret, []int{rr, cc})
                        }
                }
                rr=rr+1
        }
        for i:=0;i<=step;i++ {
           //     fmt.Println("XX",rr,cc)
                if rr >=0&&rr<R&&cc>=0&&cc<C {
                        if len(*ret) < R*C {
                                *ret = append(*ret, []int{rr, cc})
                        }
                }
                cc = cc-1
        }
        for i := 0; i <=step; i++ {
                //        fmt.Println(rr,cc)
                if rr >= 0&&rr < R&&cc >= 0&&cc < C {
                        if len(*ret) < R*C {
                                *ret = append(*ret, []int{rr, cc})
                        }
                }
                rr = rr - 1
        }
        //fmt.Println(*ret)
}