package main

import (
        "fmt"
)


/////////////////////////////
// fold from right
// reduce from right.
/*
f(1,f(2,f(3,z)))
f(a,b) = a+b
*/
func foldr(f func(a,b int) int,z int,list []int)int{
        if len(list) == 0{
                return z
        }
        return f(list[0],foldr(f,z,list[1:]))
}
/*
f(f(f(z,1),2),3)

*/
func foldl(f func(a,b int) int,z int,list []int)int{
        if len(list) == 0{
                return z
        }
        return f(foldl(f,z,list[:(len(list)-1)]),list[len(list)-1])
}

// +,*,-,&,|
func ff() {
        add := func(a,b int) int{
                return a + b
        }
        sub := func(a,b int) int{
                return a-b;
        }
        multi := func (a,b int) int{
                return a*b;
        }
        count := func(a,b int) int{
                return b+1
        }

        ilist := []int{1, 2, 3}

        fmt.Printf("foldr add : %d\n",foldr(add,0,ilist))
        fmt.Printf("foldr sub : %d\n",foldr(sub,0,ilist))
        fmt.Printf("foldl sub : %d\n",foldl(sub,0,ilist))
        fmt.Printf("foldl multi : %d\n",foldr(multi,1,ilist))
        fmt.Printf("foldl count : %d\n",foldr(count,0,ilist))
}

func f() {
        ilist := []int{1, 2, 3}
        sum := 0
        for _,i := range ilist {
                sum += i
        }
        fmt.Printf("f add : %d\n",sum)

        sub := 0
        for _,i := range ilist {
                sub -= i
        }
        fmt.Printf("f sub : %d\n",sub)

        multi := 1
        for _,i := range ilist {
                multi *= i
        }
        fmt.Printf("f multi : %d\n",multi)
}

func main() {
        f()
}