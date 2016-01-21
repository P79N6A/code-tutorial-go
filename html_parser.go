package main
import (
	"fmt"
	"log"
	"io/ioutil"
	"os"
//	"strings"
	"regexp"
	"mustard/utils/page_analysis"
	"mustard/internal/github.com/PuerkitoBio/goquery"
	"strings"
	"mustard/internal/gopkg.in/mgo.v2"
)
func check(e error) {
	if e != nil {
		panic(e)
	}
}
type Fang struct {
	// 标题，小区，租金，户型，地址，配置，经纬度
	Url	string
	Title string
	Community string
	Cost	string
	PayMethod	string
	HouseType	string
	HouseType2	string
	Address	string
	Configuration	string
	BaiduLongitude	string
	BaiduLatitude	string
	Longitude	string
	Latitude	string
	UpdateTime	string

}
func parse(s string) Fang {
	var fa Fang
	path := strings.Split(s,"/")
	raw_url := path[len(path)-1]
	raw_url = strings.Replace(raw_url, "___","://", 1)
	raw_url = strings.Replace(raw_url, "_","/", -1)
	fa.Url = raw_url
	dat, err := ioutil.ReadFile(s)
	check(err)
	var parser page_analysis.HtmlParser
	parser.Parse(raw_url,string(dat))
	doc := parser.GetDocument()
	doc.Find("title").Each(func(i int, s *goquery.Selection){
		fa.Title = strings.TrimSpace(s.Text())
	})
	doc.Find("div.xiaoqu").Each(func(i int, s *goquery.Selection){
		str := strings.Replace(s.Text(),"\n","",-1)
		str = strings.Replace(str,"\t","",-1)
		str = strings.Replace(str," ","",-1)
		fmt.Println(str)
		fa.Community = str
	})
	doc.Find("span.pay-method").Each(func(i int, s *goquery.Selection){
		str := strings.Replace(s.Text(),"\n","",-1)
		str = strings.Replace(str,"\t","",-1)
		str = strings.Replace(str," ","",-1)
		fmt.Println(str)
		fa.PayMethod = str
	})
	doc.Find("em.house-price").Each(func(i int, s *goquery.Selection){
		str := strings.Replace(s.Text(),"\n","",-1)
		str = strings.Replace(str,"\t","",-1)
		str = strings.Replace(str," ","",-1)
		fmt.Println(str)
		fa.Cost = str
	})
	doc.Find("div.house-type").Each(func(i int, s *goquery.Selection){
		str := strings.Replace(s.Text(),"\n","",-1)
		str = strings.Replace(str,"\t","",-1)
		str = strings.Replace(str," ","",-1)
		fmt.Println(str)
		fa.HouseType = str
	})
	doc.Find("span.pl10").Each(func(i int, s *goquery.Selection){
		str := strings.Replace(s.Text(),"\n","",-1)
		str = strings.Replace(str,"\t","",-1)
		str = strings.Replace(str," ","",-1)
		fmt.Println(str)
		if strings.Contains(str, "更新时间") {
			fa.UpdateTime = str
		}
	})
	doc.Find("li.house-primary-content-li").Each(func(i int, s *goquery.Selection){
		str := strings.Replace(s.Text(),"\n","",-1)
		str = strings.Replace(str,"\t","",-1)
		str = strings.Replace(str," ","",-1)
		fmt.Println(str)
		if strings.Contains(str, "租金") {
	//		fa.CostFull = str
		} else if strings.Contains(str, "房屋") {
			fa.HouseType2 = str
		} else if strings.Contains(str, "地址") {
			fa.Address = str
		} else if strings.Contains(str, "配置") {
			fa.Configuration = str
		}
	})
	l, _ := regexp.Compile(`baidulon:'(\d+\.\d+)'`)
	baiduLong := l.FindStringSubmatch(string(dat))
	if len(baiduLong) > 1 {
		fa.BaiduLongitude = baiduLong[1]
	}
	r, _ := regexp.Compile(`baidulat:'(\d+\.\d+)'`)
	baiduLat := r.FindStringSubmatch(string(dat))
	if len(baiduLat) > 1 {
		fa.BaiduLatitude = baiduLat[1]
	}
	ll, _ := regexp.Compile(`,lon:'(\d+\.\d+)'`)
	longs := ll.FindStringSubmatch(string(dat))
	if len(longs) > 1 {
		fa.Longitude = longs[1]
	}
	rr, _ := regexp.Compile(`,lat:'(\d+\.\d+)'`)
	lat := rr.FindStringSubmatch(string(dat))
	if len(lat) > 1 {
		fa.Latitude = lat[1]
	}
	return fa
}
func save2Mongo(fa *Fang) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("crawl").C("crawl_58")
	err = c.Insert(fa)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	filename := os.Args[1]
	fa := parse(filename)
	save2Mongo(&fa)
}
