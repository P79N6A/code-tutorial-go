package main
/*
Given an array of strings, group anagrams together.

Example:

Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
Output:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
Note:

All inputs will be in lowercase.
The order of your output does not matter.

*/
// 问题的关键在于如何将一个字符串hash到一个数字上.哈希函数定义至关重要.
// 这个case给了我们一个提示. 我们将每个字母映射到一个prime 质数上, 则质数相乘的结果肯定是唯一不碰撞的.
// 质数的乘法作为不同位置出现的序列方式很有建设性~~~~~很有创意
// 这种应用场景很多...

/*
map的key,也就是哈希函数定义有很多
1. 将str排序归一化,这个问题在于需要排序
2. 进制转换,但问题在于不确定a出现多少次
3. 使用prime相乘,这也是唯一的.但问题在于可能会越界
4. 如果只出现一次,是0/1的问题,则使用bitset
*/

func groupAnagrams(strs []string) [][]string {
        // a-z 映射到质数
        prime := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103};//最多10609个z
        res := make(map[int][]string)
        for _,str := range strs {
                key := 1
                for _,s := range str {
                        key *= prime[int(s-'a')]
                }
                if _,ok := res[key];!ok {
                        res[key] = make([]string,0)
                }
                res[key] = append(res[key],str)
        }
        ret := make([][]string,0)
        for _,r := range res {
                ret = append(ret,r)
        }
        return ret
}
func main() {
}
