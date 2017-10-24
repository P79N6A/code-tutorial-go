package main

import (
        "regexp"
        "fmt"
        "strings"
)

var html = `<h1>Machine Info</h1><key>cmd</key>:<value>../bin/glory-dns_main --conf_path_prefix=../mdata --v=6 --max_arecord_num=8 --default_a_weight_slot=800:20:20:20:20:20:20:20:20:20:20 --query_log=../logs/query.log --http_port=9750 --adjust_rr_mode=ROTATE --dns_port=9053 --msg_compress=true --http_dns_port=9153 --reload_record_interval=15 --dns_host= --enable_qos=true --enable_statistic=true --always_fill_scopenetmask=true --reload_ipaddr_shift_interval=25 --enable_qos=true --nxdomain_no_return=false --badreq_no_return=true --stdout=true</value><br><key>pid</key>:<value>11045</value><br><key>uid</key>:<value>501</value><br><key>hostname</key>:<value>gaolichuangdeMacBook-Pro.local</value><br><key>StartAt</key>:<value>2016-12-13 14:32 CST</value><br><h1>Status Info</h1><h1>Healthy</h1><div>true</div><h1>Application Info</h1><br>LocalIP:192.168.3.127; Req:0<br>ShiftIPNum:0<br>Version:2<br>TotalReq:0<br>SuccessRatio:0.00000<br>NxDomainRatio:0.00000<br>NoRetRatio:0.00000<br>QPS:0<br>NXDomainPS:0<br>NoRetPS:0<br>SuccessPS:0<br>MaxProccessTime:0 ms<br>MaxProccessMsgTime:1970-01-01 08:00 CST<br>MinProccessTime:0 ms<br>AvgProccessTime:0.00 ms<br>IPSetLoad:<nil><br>IPAddrBase:<nil><br>ZoneManager:ConfVersion:2,CurrentVersion:2,LoadUse:860ms,Err:[],DomainRR:27,GeneralDomainRR:1,Zone:[dn-dns.com. flxdns.com. tom.com.],SubZone:[],Smart:[ngnop2502.dn-dns.com. ngnop2502.flxdns.com. ngnop2502.tom.com. nxnop013.dn-dns.com. nxnop013.flxdns.com. nxnop013.tom.com. nxnop061.dn-dns.com. nxnop061.flxdns.com. nxnop061.tom.com.]<br>`
func main() {
        html = strings.Replace(html,"\n","",-1)
        fmt.Println(html)
        pairs := make(map[string]string)
        r, _ := regexp.Compile("<key>.*?</key>:<value>.*?</value>")
        for _,v := range r.FindAllString(html,-1) {
                fs := strings.Split(v,"</key>:<value>")
                if len(fs) != 2 {
                        continue
                }
                fs[0] = strings.TrimLeft(fs[0],"<key>")
                fs[1] = strings.TrimRight(fs[1],"</value>")
                pairs[fs[0]] = fs[1]
        }
        // get healthy
        rh,_ :=regexp.Compile("<h1>Healthy</h1><div>.*?</div>")
        healthy := rh.FindString(html)
        healthy = strings.Replace(healthy,"<h1>Healthy</h1><div>","",-1)
        healthy = strings.TrimRight(healthy,"</div>")

        fmt.Println("XXXXXXXXXXXXXXXX")
        rk, _ := regexp.Compile("br>.*?:.*?<")
        for _,v := range rk.FindAllString(html,-1) {
                v = strings.TrimLeft(v,"br>")
                v = strings.TrimRight(v,"<")
                fs := strings.Split(v,":")
                if len(fs) != 2 {
                        continue
                }
                pairs[fs[0]] = fs[1]
        }
        pairs["healthy"] = healthy
        for k,v := range pairs {
                fmt.Printf("%s : %s\n",k,v)
        }
}

