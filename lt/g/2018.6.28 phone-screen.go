package main

import "fmt"

/*
6.29日给我通知， 说coding有较为严重的bug，不给通过。。。 fail了，准备了两个月的g就这样解释了。。。
*/
/*
2018.6.28  3:30pm-4:15pm
Please use this Google doc to code during your interview. To free your hands for coding, we recommend that you use a headset or a phone with speaker option.

https://docs.google.com/document/d/1LlYtpwjase1rM4ycE0uVV3NGRUdSyLyRvsALTpS3ewc/edit?dac=1

flxdns.com.

A Weighted Item is a pair of int "id" and float "weight".
Weighted Item Array: An array with no repeated id, and the weights are in descending order.
Input: N Weighted Item array.
Output: A new Weighted Item array.

A: (id: 1, w: 2) (id: 5, w: 1)
B: (id: 3, w: 3) (id: 6, w: 1.5)

output: (id 3, w: 3), (id: 1, w: 2) (id: 6, w: 1.5) (id 5, w: 1)


A: (id: 1, w: 5) (id: 5, w: 4)
B: (id: 5, w: 5) (id: 1, w: 1.5)

output: (id 1, w: 3) (id: 6, w: 1.5) (id 5, w: 1)

O(len(A)+len(B))

Merge sort O(MlogN)  N is arrary num, M array length sum
*/

type item struct {
    id int
    w  float32
}

func main() {
    mm := mergeN([][]*item{
        []*item{&item{1, 5}, &item{5, 4}},
        []*item{&item{5, 5}, &item{1, 1.5}},
        []*item{&item{2, 5}, &item{3, 1.5}},
    })
    print(mm)
    mm = mergeN(nil)
    print(mm)
}
func mergeN(items [][]*item) []*item {
    nitem := mergeNR(items)
    return nitem[0]
}
func mergeNR(items [][]*item) [][]*item {
    if len(items) <= 0 {
        return [][]*item{nil}
    }
    if len(items) <= 1 {
        return items
    }
    nitems := make([][]*item, 0)
    if len(items)%2 == 1 {
        items = append(items, nil)
    }
    // 0,1,2,3
    for j := 0; j+1 < len(items); j += 2 {
        ni := mergeTwo(items[j], items[j+1])
        nitems = append(nitems, ni)
    }
    return mergeNR(nitems)
}
func print(it []*item) {
    str := ""
    for _, i := range it {
        str += fmt.Sprintf("%d,%f;", i.id, i.w)
    }
    fmt.Println(str)
}
func mergeTwo(items1, items2 []*item) []*item {
    // items1,items2 could be nil.
    ret := make([]*item, 0)
    ids := make(map[int]bool)
    i1, i2 := 0, 0
    for i1 < len(items1) && i2 < len(items2) {
        if items1[i1].id == items2[i2].id {
            if ids[items1[i1].id] == false {
                if items1[i1].w > items2[i2].w {
                    ret = append(ret, items1[i1])
                } else {
                    ret = append(ret, items2[i2])
                }
                ids[items1[i1].id] = true
            }
            i1, i2 = i1+1, i2+1
        } else {
            if items1[i1].w > items2[i2].w {
                if ids[items1[i1].id] == false {
                    ret = append(ret, items1[i1])
                    ids[items1[i1].id] = true
                }
                i1 += 1
            } else {
                if ids[items2[i2].id] == false {
                    ret = append(ret, items2[i2])
                    ids[items2[i2].id] = true
                }
                i2 += 1
            }
        }
    }
    for i1 < len(items1) {
        if ids[items1[i1].id] == false {
            ret = append(ret, items1[i1])
            ids[items1[i1].id] = true
        }
        i1 += 1
    }
    /*
    // 错误的代码，循环的出口没看仔细，导致无法走出循环。。。。
    // 并且在ids设置i2之前，i2已经变化了。
    // i2++ 需要在任何时候都进行。
    for i2 < len(items2) {
        if ids[items2[i2].id] == false {
            ret = append(ret,items2[i2])
            i2 += 1
            ids[items2[i2].id]=true
        }
    }
    */
    for i2 < len(items2) {
        if ids[items2[i2].id] == false {
            ret = append(ret, items2[i2])
            ids[items2[i2].id] = true
        }
        i2 += 1
    }
    return ret
}
