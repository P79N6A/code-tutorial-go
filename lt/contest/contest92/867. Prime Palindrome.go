package main

import (
        "math"
        "fmt"
)

func primePalindrome(N int) int {
        /*
        1xxx1
        3xxx3
        5xxx5
        7xxx7
        9xxx9
        */
        num,mN := 0,N
        for mN > 0 {
                mN = mN/10
                num += 1
        }
        for i:=num;;i++ {
                sub := genPalindrome(i)
                for _,s := range sub {
                        if s >= N && isPrime(s) {
                                return s
                        }
                }
        }
        return -1
}
func genPalindrome(n int) []int {
        // TODO.
        if n == 0 {return nil}
        if n == 1 {return []int{0,1,2,3,4,5,6,7,8,9}}
        if n == 2 {return []int{0,11,22,33,44,55,66,77,88,99}}
        ret := make([]int,0)
        sub := genPalindrome(n-2)
        //ret = append(ret,0)
        for i:=0;i<=9;i++ {
                for _,s := range sub{
                        r := i
                        for j:=1;j<n;j++ {
                              r *=10
                        }
                        rt := r+s*10+i
                        ret = append(ret, rt)
                }
        }
        return ret

}
func isPrime(num int) bool {
        if num == 1 {return false}
        if num == 2 || num == 3 {return true}
        if num %6 !=1 && num%6 != 5 {return false}
        t := int(math.Sqrt(float64(num)))
        for i:=5;i<=t;i+=6 {
                if num%i==0 || num%(i+2)==0 {
                        return false
                }
        }
        return true
}
func main() {
        fmt.Println(primePalindrome(930))
        fmt.Println(genPalindrome(5))
        fmt.Println(isPrime(1))
}
