package main
import (
	"log"
	"fmt"
	"io/ioutil"
	//"os"
	"mustard/base/string_util"
	"mustard/utils/page_analysis"
	"mustard/internal/github.com/PuerkitoBio/goquery"
	"strings"
	"mustard/internal/gopkg.in/mgo.v2"
)
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
func check(e error) {
	if e != nil {
		panic(e)
	}
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
func parse2(s string) Fang {
	var fa Fang
	path := strings.Split(s,"/")
	raw_url := path[len(path)-1]
	raw_url = strings.Replace(raw_url, "___","://", 1)
	raw_url = strings.Replace(raw_url, "_","/", -1)
	fa.Url = raw_url
	dat, err := ioutil.ReadFile(s)
	check(err)
	var parser page_analysis.HtmlParser
	parser.RegisterSelector("title", func(i int, s *goquery.Selection){
		fa.Title = strings.TrimSpace(s.Text())
	})
	parser.RegisterSelector("div.xiaoqu", func(i int, s *goquery.Selection){
		str := string_util.Purify(s.Text(), "\n","\t"," ")
		fmt.Println(str)
		fa.Community = str
	})
	parser.RegisterSelector("span.pay-method", func(i int, s *goquery.Selection){
		str := string_util.Purify(s.Text(), "\n","\t"," ")
		fmt.Println(str)
		fa.PayMethod = str
	})
	parser.RegisterSelector("em.house-price", func(i int, s *goquery.Selection){
		str := string_util.Purify(s.Text(), "\n","\t"," ")
		fmt.Println(str)
		fa.Cost = str
	})
	parser.RegisterSelector("div.house-type", func(i int, s *goquery.Selection){
		str := string_util.Purify(s.Text(), "\n","\t"," ")
		fmt.Println(str)
		fa.HouseType = str
	})
	parser.RegisterSelectorWithTextKeyWord("span.pl10", "更新时间", func(i int, s *goquery.Selection){
		str := string_util.Purify(s.Text(), "\n","\t"," ")
		fmt.Println(str)
		fa.UpdateTime = str
	})
	parser.RegisterSelectorWithTextKeyWord("li.house-primary-content-li", "房屋", func(i int, s *goquery.Selection) {
		fa.HouseType2 = string_util.Purify(s.Text(), "\n", "\t", " ")
	})
	parser.RegisterSelectorWithTextKeyWord("li.house-primary-content-li", "地址", func(i int, s *goquery.Selection) {
		fa.Address = string_util.Purify(s.Text(), "\n", "\t", " ")
	})
	parser.RegisterSelectorWithTextKeyWord("li.house-primary-content-li", "配置", func(i int, s *goquery.Selection) {
		fa.Configuration = string_util.Purify(s.Text(), "\n", "\t", " ")
	})
	parser.RegisterRegex(`baidulon:'(\d+\.\d+)'`, func(i int, r[]string) {
		fa.BaiduLongitude = r[1]
	})
	parser.RegisterRegex(`baidulat:'(\d+\.\d+)'`, func(i int, r[]string) {
		fa.BaiduLatitude = r[1]
	})
	parser.RegisterRegex(`,lon:'(\d+\.\d+)'`, func(i int, r[]string) {
		fa.Longitude = r[1]
	})
	parser.RegisterRegex(`,lat:'(\d+\.\d+)'`, func(i int, r[]string) {
		fa.Latitude = r[1]
	})
	parser.Parse(raw_url,string(dat))
	return fa
}
func main() {
	//filename := os.Args[1]
	filename := "/Users/glc/workspace/test/http___bj.58.com_zufang_24725387601214x.shtml"
	fa := parse2(filename)
	save2Mongo(&fa)
}
