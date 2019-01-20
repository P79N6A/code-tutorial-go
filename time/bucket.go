package main

import (
    "fmt"
    "time"
)

func getBucketKeyByTs(ts int64,bucket time.Duration) string {
    cur := time.Unix(ts,0)
    hour,minute := cur.Hour(),cur.Minute()
    return genBucketKey(hour,minute,bucket)
}
func genBucketKey(hour,minute int,bucket time.Duration) string {
    key := hour*60 + minute
    step := bucket / time.Minute
    key -= key%int(step)
    return fmt.Sprintf("%02d%02d",key / 60,key % 60)
}
func getAllBucketKeys(bucket time.Duration) []string {
    ans := make([]string,0)
    step := bucket / time.Minute
    hour,minute := 0,0
    for hour < 24 {
        for minute < 60 {
            ans = append(ans,fmt.Sprintf("%02d%02d",hour,minute))
            minute += int(step)
        }
        minute = 0
        hour += 1
    }
    return ans
}
func main() {
    fmt.Println(genBucketKey(2,34,time.Minute*30))
    fmt.Println(genBucketKey(2,24,time.Minute*30))
    fmt.Println(genBucketKey(0,59,time.Minute*30))
    fmt.Println(genBucketKey(0,59,time.Minute*10))
    fmt.Println(getBucketKeyByTs(time.Now().Unix(),time.Minute*10))
    fmt.Println(getBucketKeyByTs(time.Now().Unix(),time.Minute*20))
    fmt.Println(getBucketKeyByTs(time.Now().Unix(),time.Minute*30))
    fmt.Println(getBucketKeyByTs(time.Now().Unix(),time.Hour*24))
    var _all_bucket_keys []string = func() []string {
        ans := make([]string,0)
        for i:=0;i<24;i++ {
            for j:=0;j<60;j+=5 {
                ans = append(ans,fmt.Sprintf("%02d%02d",i,j))
            }
        }
        return ans
    }()
    fmt.Println(_all_bucket_keys)
    fmt.Println(len(_all_bucket_keys))
    fmt.Println(getAllBucketKeys(time.Minute*5))
    fmt.Println(len(getAllBucketKeys(time.Minute*20)))
}
