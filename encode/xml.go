package main

import (
	"encoding/xml"
	"fmt"
	"time"
)

const atomFeedString = `
<?xml version="1.0" encoding="utf-8"?>
<feed xmlns="http://www.w3.org/2005/Atom" xml:lang="en-us" updated="2009-10-04T01:35:58+00:00">
	<title>Code Review - My issues</title>
	<link href="http://codereview.appspot.com/" rel="alternate"></link>
	<link href="http://codereview.appspot.com/rss/mine/rsc" rel="self"></link>
	<id>http://codereview.appspot.com/</id>
	<author><name>rietveld&lt;&gt;</name></author>
	<entry>
		<title>rietveld: an attempt at pubsubhubbub</title>
		<link href="http://codereview.appspot.com/126085" rel="alternate"></link>
		<updated>2009-10-04T01:35:58+00:00</updated>
		<author><name>email-address-removed</name></author>
		<id>urn:md5:134d9179c41f806be79b3a5f7877d19a</id>
		<summary type="html">
			the top of feeds.py marked NOTE(rsc).
		</summary>
	</entry>
	<entry>
		<title>rietveld: correct tab handling</title>
		<link href="http://codereview.appspot.com/124106" rel="alternate"></link>
		<updated>2009-10-03T23:02:17+00:00</updated>
		<author><name>email-address-removed</name></author>
		<id>urn:md5:0a2a4f19bb815101f0ba2904aed7c35a</id>
		<summary type="html">
			This fixes the buggy tab rendering that can be seen at
		</summary>
	</entry>
</feed>`

type Feed struct {
	XMLName xml.Name      `xml:"http://www.w3.org/2005/Atom feed"`
	Title   string    `xml:"title"`
	Id      string    `xml:"id"`
	Link    []Link    `xml:"link"`
	Updated time.Time `xml:"updated,attr"`
	Author  Person    `xml:"author"`
	Entry   []Entry   `xml:"entry"`
}

type Entry struct {
	Title   string    `xml:"title"`
	Id      string    `xml:"id"`
	Link    []Link    `xml:"link"`
	Updated time.Time `xml:"updated"`
	Author  Person    `xml:"author"`
	Summary Text      `xml:"summary"`
}

type Link struct {
	Rel  string `xml:"rel,attr,omitempty"`
	Href string `xml:"href,attr"`
}

type Person struct {
	Name     string `xml:"name"`
	URI      string `xml:"uri"`
	Email    string `xml:"email"`
	InnerXML string `xml:",innerxml"`
}

type Text struct {
	Type string `xml:"type,attr,omitempty"`
	Body string `xml:",chardata"`
}

func main() {
	var f Feed
	if err := xml.Unmarshal([]byte(atomFeedString), &f); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n",f)

}
