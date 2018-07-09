package main

import (
    "math"
    "fmt"
)

/*
基础计算逻辑。 验证从2到开方之间数字能否被他整除。
超过这个不可能有结果了
一个数若可以进行因数分解，那么分解时得到的两个数一定是一个小于等于sqrt(n)，一个大于等于sqrt(n)
O(sqrt(n))
*/
func isPrime(n int) bool {
    if n < 2 {return false}
    for i:=2;i*i<=n;i++ {
        if n%i == 0 {return false}
    }
    return true
}
/*
首先看一个关于质数分布的规律：大于等于5的质数一定和6的倍数相邻。例如5和7，11和13,17和19等等；

证明：令x≥1，将大于等于5的自然数表示如下：
······ 6x-1，6x，6x+1，6x+2，6x+3，6x+4，6x+5，6(x+1），6(x+1)+1 ······
可以看到，不在6的倍数两侧，即6x两侧的数为6x+2，6x+3，6x+4，由于2(3x+1)，3(2x+1)，2(3x+2)，
所以它们一定不是素数，再除去6x本身，显然，素数要出现只可能出现在6x的相邻两侧。
*/
func isPrime2(num int) bool {
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
    fmt.Println(isPrime(1000))
    
}
