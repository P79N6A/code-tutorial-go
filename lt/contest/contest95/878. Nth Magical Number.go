package main

import (
        "fmt"
)

func main() {
        fmt.Println(nthMagicalNumber(1, 2, 3))//2
        fmt.Println(nthMagicalNumber(4, 2, 3))//6
        fmt.Println(nthMagicalNumber(5, 2, 4))//10
        fmt.Println(nthMagicalNumber(3,6,4))//10
        fmt.Println(nthMagicalNumber(1000000000,40000, 40000))//10
        fmt.Println(nthMagicalNumber(1000000000,40000, 40000))//10
}
/*
  int nthMagicalNumber(int N, int A, int B) {
        long a = A, b = B, tmp, l = 2, r = pow(10, 14);
        while (b) tmp = a, a = b, b = tmp % b;
        while (l < r) {
            long m = (l + r) / 2;
            if (m / A + m / B - m / (A * B / a) < N) l = m + 1;
            else r = m;
        }
        return l % (long)(pow(10, 9) +7);
    }
*/


func nthMagicalNumber(N int, A int, B int) int {
        if A>B {A,B=B,A}
        lcm := A*B/gcd(A,B)
        l,r := A,B*N
        for l<r {
                m := (l+r)/2
                c := m/A+m/B-m/lcm
                if c <N {
                        l=m+1
                } else {
                        r=m
                }
        }
        return r%1000000007
}

func gcd(A, B int) int {
        if (A < B) {
                A, B = B, A
        }
        if (A % B == 0) {
                return B
        } else {
                return gcd(B, A % B)
        }
}
/**/


