package main

import "fmt"

/*
Write a class StockSpanner which collects daily price quotes for some stock, and returns the span of that stock's price for the current day.

The span of the stock's price today is defined as the maximum number of consecutive days (starting from today and going backwards) for which the price of the stock was less than or equal to today's price.

For example, if the price of a stock over the next 7 days were [100, 80, 60, 70, 60, 75, 85], then the stock spans would be [1, 1, 1, 2, 1, 4, 6].



Example 1:

Input: ["StockSpanner","next","next","next","next","next","next","next"], [[],[100],[80],[60],[70],[60],[75],[85]]
Output: [null,1,1,1,2,1,4,6]
Explanation:
First, S = StockSpanner() is initialized.  Then:
S.next(100) is called and returns 1,
S.next(80) is called and returns 1,
S.next(60) is called and returns 1,
S.next(70) is called and returns 2,
S.next(60) is called and returns 1,
S.next(75) is called and returns 4,
S.next(85) is called and returns 6.

Note that (for example) S.next(75) returned 4, because the last 4 prices
(including today's price of 75) were less than or equal to today's price.


Note:

Calls to StockSpanner.next(int price) will have 1 <= price <= 10^5.
There will be at most 10000 calls to StockSpanner.next per test case.
There will be at most 150000 calls to StockSpanner.next across all test cases.
The total time limit for this problem has been reduced by 75% for C++, and 50% for all other languages.
*/

func main() {
    obj := Constructor();
    fmt.Println(obj.Next(100))
    fmt.Println(obj.Next(80))
    fmt.Println(obj.Next(60))
    fmt.Println(obj.Next(70))
    fmt.Println(obj.Next(60))
    fmt.Println(obj.Next(75))
    fmt.Println(obj.Next(85))
}
type StockSpanner struct {
    decStack []int
    nStack []int
}
/*
题目问，新来的一个price，之前有多少连续天数的price都小于等于他
是个stack问题，可以减少计算量。那就把比当前栈顶小的个数记录着。
递减栈。
*/

func Constructor() StockSpanner {
    return StockSpanner{
        decStack:make([]int,0),
        nStack:make([]int,0),
    }
}


func (this *StockSpanner) Next(price int) int {
    span := 1
    // 栈的处理都是先搞出栈，在处理入栈，这样逻辑判断代码会少一些
    for len(this.decStack) > 0 && price >= this.decStack[len(this.decStack)-1] { // 栈不空，并且栈顶加个比当前加个低
        span += this.nStack[len(this.nStack)-1] // span叠加
        this.nStack = this.nStack[:(len(this.nStack)-1)] // 出栈
        this.decStack=this.decStack[:(len(this.decStack)-1)] //出栈
    }
    // 入栈，span计算完毕
    this.nStack = append(this.nStack,span)
    this.decStack = append(this.decStack,price)
    return span
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */