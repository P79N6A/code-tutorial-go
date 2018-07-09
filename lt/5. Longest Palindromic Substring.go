package main

import "fmt"

func longestPalindrome(s string) string {
        slen := len(s)
        dp := make([][]int,slen)
        for i:=0;i<slen;i++ {
                dp[i]=make([]int,slen)
        }
        left,right,maxl := 0,0,0
        for i:=0;i<slen;i++ {
                for j:=0;j<i;j++ {
                        if s[i] == s[j] {
                                if i-j < 2 {
                                        dp[j][i]=1
                                } else {
                                        dp[j][i]=dp[j+1][i-1]
                                }
                        }
                        if dp[j][i] == 1 && maxl < i-j+1 {
                                //  如果是回文,看是否是最大的.
                                // 可以衍生出来 计算回文个数等.
                                maxl = i-j+1
                                left,right =j,i
                        }
                }
                dp[i][i]=1
        }
        fmt.Println(left,right)

        return s[left:right+1]
}

func main() {
        fmt.Println(longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"))
}
