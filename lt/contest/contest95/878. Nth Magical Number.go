package main

import (
        "fmt"
        "math"
)

func main() {
        //fmt.Println(nthMagicalNumber(1, 2, 3))//2
        //fmt.Println(nthMagicalNumber(4, 2, 3))//6
        fmt.Println(nthMagicalNumber(5, 2, 3))//10
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
func nthMagicalNumber(N,A,B int) int {
        var a,b,t,l,r float64 = float64(A),float64(B),0,2,math.Pow(10,14)
        for b!= 0 {
                t=a
                a=b
                b=float64(int(t)%int(b))
        }
        for l<r {
                var m float64 = (l+r)/2
                if int(m)/(A)+int(m)/(B)-int(m)/(A*B/gcd(A,B)) < (N) {
                        l=m+1
                } else {
                        r=m
                }
        }
        return int(r)
}


func nthMagicalNumber1(NN int, AA int, BB int) int {
        var A, B, N float64 = float64(AA), float64(BB), float64(NN)
        if A > B {
                A, B = B, A
        }
        var lcm float64 = float64(AA*BB/gcd(AA,BB))
        fmt.Println(AA,BB,lcm)
        var l, r float64 = A, B*N*A
        //fmt.Println(l,r)
        for l < r {
                var m float64 = l + (r - l) / 2
                counter := m / A + m / B - m /lcm
                fmt.Println("idx",m,l,r,counter)
                if counter >= N {
                        r = m
                } else {
                        l = m + 1
                }
        }
        return int(r)
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


