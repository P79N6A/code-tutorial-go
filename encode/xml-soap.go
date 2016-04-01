package main

import (
	"encoding/xml"
	"fmt"
)

const atomFeedString = `
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		<ns2:getRoomInfoResponse xmlns:ns2="http://ws.easy.dnion.com/">
			<return>[xxxx]</return>
		</ns2:getRoomInfoResponse>
	</soap:Body>
</soap:Envelope>`
type Envelope struct {
	XMLName xml.Name
	Body Body `xml:"Body"`
}

type Body struct {
	NS  NSResponse `xml:"getRoomInfoResponse"`
}

type NSResponse struct {
	NS  string `xml:"ns2,attr,omitempty"`
	Ret Return `xml:"return"`
}

type Return struct {
	RealData string `xml:",chardata"`
}


func main() {
	var f Envelope
	if err := xml.Unmarshal([]byte(atomFeedString), &f); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n",f)
}
