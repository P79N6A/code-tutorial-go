package main

type BIT struct {
    nums []int
    bit []int
}

func (t *BIT)update(i int, val int) {
    diff := t.nums[i]-val
    for j:=i;j<len(t.bit);j+=(j&-j) {
        t.bit[j] += diff
    }
    t.nums[i]+=val
}
func (t *BIT)sumRange(i,j int) int {
    return t.getsum(j+1)-t.getsum(i)
}
func (t *BIT)getsum(i int) int {
    res := 0
    for j:=i;i>=0;j-=(j&-j) {
        res += t.bit[j]
    }
    return res
}

func main() {
    
}
