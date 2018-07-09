package main

import "fmt"

/*
Initially, there is a Robot at position (0, 0).
Given a sequence of its moves, judge if this robot makes a circle,
which means it moves back to the original place.

The move sequence is represented by a string.
And each move is represent by a character. The valid robot moves are
R (Right), L (Left), U (Up) and D (down).
The output should be true or false representing whether the robot makes a circle.

Example 1:
Input: "UD"
Output: true
Example 2:
Input: "LL"
Output: false
*/
// 出现circle的判断，并非是回到原点的判断
func judgeCircle(moves string) bool {
    ws := split(moves)
    path := make(map[string]bool)
    x,y := 0,0
    path[convert(x,y)]=true
    for _,w := range ws {
        u,d,l,r := 0,0,0,0
        for i:=0;i<len(w);i++ {
            if w[i]=='U' {u += 1}
            if w[i]=='D' {d += 1}
            if w[i]=='L' {l += 1}
            if w[i]=='R' {r += 1}
            if l + r > 0 {
                if l > r {
                    for j:=0;j<(l-r);j++ {
                        x -= 1
                        if path[convert(x,y)] == true {
                            return true
                        }
                        path[convert(x,y)]=true
                    }
                } else if l < r {
                    for j:=0;j<(r-l);j++ {
                        x += 1
                        if path[convert(x,y)] == true {
                            return true
                        }
                        path[convert(x,y)]=true
                    }
                } else {
                    if len(ws) == 1 {
                        return true
                    }
                }
            } else {
                if u > d {
                    for j:=0;j<(u-d);j++ {
                        y += 1
                        if path[convert(x,y)] == true {
                            return true
                        }
                        path[convert(x,y)]=true
                    }
                } else if u < d {
                    for j:=0;j<(d-u);j++ {
                        y -= 1
                        if path[convert(x,y)] == true {
                            return true
                        }
                        path[convert(x,y)]=true
                    }
                } else {
                    if len(ws) == 1 {
                        return true
                    }
                }
            }
        }
    }
    return false
}
func convert(a, b int) string {
    return fmt.Sprintf("%d:%d", a, b)
}
func split(str string)[]string {
    if len(str) <= 0{return nil}
    ret := make([]string,0)
    word,ve := "",false
    if str[0]=='U' || str[0]=='D' {
        ve = true
    }
    for i:=0;i<len(str);i++ {
        if str[i]=='U' || str[i]=='D' {
            if ve {
                if word != "" {
                    ret = append(ret,word)
                }
                ve,word = false,""
            }
        } else {
            if !ve {
                if word != "" {
                    ret = append(ret,word)
                }
                ve,word = true,""
            }
        }
        word += string(str[i])
    }
    if word != "" {
        ret = append(ret,word)
    }
    return ret
}
func judgeCircle1(moves string) bool {
    path := make(map[string]bool)
    startx, starty, endx, endy := 0, 0, 0, 0
    start, end := 0, 0
    for end < len(moves) {
        if moves[start] == 'U' || moves[start] == 'D' {
            for end < len(moves) {
                if moves[end] == 'U' {
                    endx, endy,end = endx, endy+1,end+1
                } else if moves[start] == 'D' {
                    endx, endy,end = endx, endy-1,end+1
                } else {
                    break
                }
            }
            l, h := starty, endy
            if starty > endy {
                l, h = endy, starty

            }
            for i := l; i <= h; i++ {
                if path[convert(endx, i)] == true {
                    return false
                }
                path[convert(endx, i)] = true
            }
            start = end
            startx, starty = endx, endy
        } else {
            for end < len(moves) {
                if moves[end] == 'L' {
                    endx, endy,end = endx-1, endy,end+1
                } else if moves[start] == 'R' {
                    endx, endy,end = endx+1, endy,end+1
                } else {
                    break
                }
            }
            l, h := startx, endx
            if startx > endx {
                l, h = endx, startx
            }
            for i := l; i <= h; i++ {
                if path[convert(i, endy)] == true {
                    return false
                }
                path[convert(i, endy)] = true
            }
            start = end
            startx, starty = endx, endy
        }
    }

    return false
}

func main() {
    //fmt.Println(judgeCircle("LDRRLRUULR"))
    //fmt.Println(judgeCircle("LDRU"))
    fmt.Println(judgeCircle("LR"))

}
