package main
/*
Given a non-empty array of unique positive integers A, consider the following graph:

There are A.length nodes, labelled A[0] to A[A.length - 1];
There is an edge between A[i] and A[j] if and only if A[i] and A[j] share a common factor greater than 1.
Return the size of the largest connected component in the graph.



Example 1:

Input: [4,6,15,35]
Output: 4

Example 2:

Input: [20,50,9,63]
Output: 2

Example 3:

Input: [2,3,6,7,4,12,21,39]
Output: 8

Note:

1 <= A.length <= 20000
1 <= A[i] <= 100000
*/
import (
    "fmt"
    "gcodebase/time_util"
    "math"
)

func main() {
    //fmt.Println(largestComponentSize([]int{4,6,15,35}))
    //fmt.Println(largestComponentSize([]int{20,50,9,63}))
    //fmt.Println(largestComponentSize([]int{2,3,6,7,4,12,21,39}))
    //fmt.Println(largestComponentSize([]int{99,100,69,39,14,56,91,60}))
    // 下面的case有问题
    fmt.Println(largestComponentSize([]int{5803,6153,13,2062,6161,2068,7172,8219,6174,2080,36,4138,6188,8237,46,8240,8242,4151,6202,8253,8269,2126,6226,2135,4187,97,102,9233,6263,126,3776,2178,4233,8330,9581,8342,152,6297,5487,4253,8350,2208,6308,4262,4263,6314,1053,8373,184,4281,2242,8388,6346,6352,2258,6355,2261,2084,4815,6365,2270,225,4330,4333,6525,4341,6390,3455,4355,262,8456,6410,2318,6873,4379,6435,297,2439,302,2364,2372,334,4431,2387,340,8590,345,2400,8548,6508,367,4466,371,6516,2429,2662,391,2442,8589,398,8595,2452,4506,415,2467,8612,8615,4525,6578,8627,4535,2489,445,4542,447,8644,4558,629,6611,4567,6564,6620,4582,6638,496,4596,2549,4605,513,6660,8714,4620,2579,2585,6685,544,4644,557,8816,6717,8767,6723,2628,582,4679,3853,4689,2643,599,4700,614,8807,618,4882,6768,626,4724,2677,2689,642,4740,6790,4749,158,8848,5913,4769,2724,2730,4779,8897,717,2206,719,2769,8916,6871,4825,2788,8934,6890,751,6898,4855,2811,8960,4915,8967,4877,786,2835,2843,2848,6945,2851,8997,2856,6958,6961,9011,9012,9014,6976,4931,4932,842,7309,4947,9046,7001,2910,9055,2917,1169,7017,874,7655,9072,2931,9078,2935,9024,7044,9103,912,9105,923,5021,7071,1520,5029,2982,5033,2987,7089,7667,2997,950,7533,7096,954,6033,5059,970,9719,973,3023,1669,5075,3034,2895,5084,9181,7141,7142,3050,9196,7150,9201,1015,5116,9220,9223,3083,1036,1041,175,7197,1055,1057,3108,3109,7212,3123,5173,7223,1081,7007,1085,3143,9293,5424,5205,1893,7266,1128,9325,1146,1157,5258,9357,1166,9361,7325,9385,9392,7346,3254,1208,5309,9412,9416,5321,1229,7374,1238,5339,9436,5344,1254,9451,5368,3326,5376,1282,2945,3341,3343,7440,7441,1298,3348,5397,5398,3351,3633,5416,7465,3975,1325,3376,1335,3385,9536,7492,5445,3399,3402,3408,5462,9559,9934,5474,1383,3437,7535,9593,9597,1406,5503,7555,5511,3464,1423,7568,7570,3651,7573,3481,1437,7589,3496,9649,5559,9665,1474,7627,1487,7632,1495,5592,9126,3558,593,7657,9708,3568,5618,595,3575,5625,3583,1536,1537,5634,9131,7685,5646,3601,1556,7705,9757,1572,9777,3635,7736,5692,9795,3656,7773,6758,7783,7792,9842,9847,1667,7813,8813,9873,1682,1683,9880,1648,9899,3756,1710,5811,3769,9916,7114,4725,1728,9924,5835,3790,4387,740,7906,5859,1778,3828,3830,1784,9983,1799,9994,5901,1807,1810,5909,8836,1818,5919,5921,1827,7978,5931,3900,1853,7998,3907,8005,8012,3925,8030,3935,3940,5989,8038,3730,5998,3953,8054,1914,6020,6023,1931,1934,6031,8080,1688,8516,1955,4006,4007,6059,6064,1981,4037,6087,1999,2002,4054,8159,4072,6124,2032,4083,2037,8186}))  //439
    t1 := time_util.GetTimeInMs()
    fmt.Println("cost ms",time_util.GetTimeInMs()-t1)
}
/*

联通图的算法相对简单,UnionFind 或者DFS
问题的关键在于数据量太大,两两判断是否有common fatcor(GCD)会超时. O(N2)
从factorgraph关系,反向的构建邻接表,复杂度N2,也会很耗时
问题的耗时点还是在判断一个数字有多少factor
*/
func largestComponentSize(A []int) int {
    if len(A) <= 0 {return 0}
    maxA := A[0]
    for i:=0;i<len(A);i++ {
        if maxA < A[i] {maxA=A[i]}
    }
    primes := getPrime(maxA)
    factorGraph := make(map[int][]int,0)
    dsu := NewDSU(len(A))
    for i:=0;i<len(A);i++ {
        a := A[i]
        // 这两个for循环耗时； tongguo  a%p==0 => a/=p 将prime都消掉
        for _,p := range primes {
            if a%p == 0 {
                if len(factorGraph)<=0 {
                    factorGraph[p]=[]int{i}
                } else {
                    factorGraph[p]=append(factorGraph[p],i)
                    dsu.Union(i,factorGraph[p][0])
                }
                /*
                优化计算的;把这个factor消耗完
                */
                for a%p == 0 {
                    a/=p
                }
            }
            if a == 1 {
                break
            }
        }
    }
    /*
    并查集 计算最大集合的方法！！！
    */
    ans := 0
    for _,g := range dsu.GroupSize() {
        if ans < g {ans = g}
    }
    return ans
}
type DSU struct {
    parents []int //并查集的索引
    counts []int
    maxCnt int
}
func NewDSU(size int) *DSU {
    d := &DSU{
        parents: make([]int, size),
        counts: make([]int, size),
    }
    for i := 0; i < len(d.parents); i++ {
        d.parents[i] = i // 初始化DSU makeUnion
        d.counts[i] = 1
    }
    return d
}
func (d *DSU) GroupSize() map[int]int {
    // 获取当前并查集有几个集合！
    // 并且记录每个集合有多少元素
    ans := make(map[int]int)
    for i:=0;i<len(d.parents);i++ {
        ans[d.Find(i)]+=1
    }
    return ans
}
func (d *DSU) Find(x int) int {
    if x != d.parents[x] {
        // 这个递归过程会优化parent指向
        d.parents[x] = d.Find(d.parents[x])
    }
    return d.parents[x]
}
func (d *DSU) Union(x, y int) {
    //需要考虑x,y的rank，如果能够保证x,y的关系，则不需要了
    x = d.Find(x)
    y = d.Find(y)
    if x == y {
        return
    }
    cnt := d.counts[x] + d.counts[y]
    if cnt > d.maxCnt {
        d.maxCnt = cnt
    }
    // x,y 的rank，让小的当parent
    if x < y {
        // y的parent是x；记录的count也是需要记录parent x的
        d.parents[y] = x
        d.counts[x]=cnt
    } else {
        d.parents[x] = y
        d.counts[y]=cnt
    }
}
func getPrime(cap int) []int {
    ans := make([]int,0)
    for i:=2;i<=cap;i++ {
        if isPrime2(i) {
            ans = append(ans,i)
        }
    }
    return ans
}
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


func largestComponentSize1(A []int) int {
    t1 := time_util.GetTimeInMs()
    primes := getPrime(len(A))
    fmt.Println(len(primes))
    t1 = time_util.GetTimeInMs()
    factorGraph := make(map[int][]int,0)
    for i:=0;i<len(A);i++ {
        for _,p := range primes {
            if A[i]%p == 0 {
                if len(factorGraph)<=0 {
                    factorGraph[p]=[]int{i}
                } else {
                    factorGraph[p]=append(factorGraph[p],i)
                }
            }
        }
    }
    fmt.Println("cost ms ffff",time_util.GetTimeInMs()-t1)
    graph := make([][]int,len(A))
    for _,ll := range factorGraph {
        for i:=0;i<len(ll);i++ {
            for j:=i;j<len(ll);j++ {
                if len(graph[ll[i]]) <= 0 {
                    graph[ll[i]]=[]int{ll[j]}
                } else {
                    graph[ll[i]] = append(graph[ll[i]],ll[j])
                }
                if len(graph[ll[j]]) <= 0 {
                    graph[ll[j]]=[]int{ll[i]}
                } else {
                    graph[ll[j]] = append(graph[ll[j]],ll[i])
                }
            }
        }
    }
    fmt.Println("cost ms xxx",time_util.GetTimeInMs()-t1)
    ans := 0
    visit := make([]bool,len(A))
    for i:=0;i<len(A);i++ {
        if visit[i]==false {
            one := 0
            dfs(i,graph,&visit,&one)
            if one > ans {ans = one}
        }
    }
    return ans
}
func dfs(v int,graph [][]int,visit *[]bool, num *int) {
    (*visit)[v] = true
    *num += 1
    for _,u := range graph[v] {
        if (*visit)[u] {continue}
        dfs(u,graph,visit,num)
    }
}

func hasFactor(a,b int) bool {
    return gcd(a,b) != 1
}

func gcd(A, B int) int {
    if (A < B) {A, B = B, A}
    if (A%B == 0) {
        return B
    } else {
        return gcd(B, A%B)
    }
}
