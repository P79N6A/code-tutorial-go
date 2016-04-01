package main

import (
	"fmt"
	"bytes"
	"encoding/xml"
)

type Return struct {
	Body string `xml:",chardata"`
	//RequestId int
	//DataCenterId string
	//DataCenterVersion int
	//StorageId string
}

type StorageReturn struct{
	Ret  Return `xml:"return"`
}

type Body struct {
	StrgRet StorageReturn `xml:"ns2:createStorageReturn"`
}

type StorageResponse struct{
	XMLName xml.Name
	RespBody Body `xml:"S:Body"`
}

/*
<?xml version='1.0' encoding='UTF-8'?>
<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
<S:Body>
	<ns2:createStorageReturn xmlns:ns2="http://ws.api.mysite.com/">
	<return>
		<requestId>16660663</requestId>
		<dataCenterId>ssrr-444tt-yy-99</dataCenterId>
		<dataCenterVersion>12</dataCenterVersion>
		<storageId>towrrt24903FR55405</storageId>
	</return>
	</ns2:createStorageReturn>
	</S:Body>
</S:Envelope>`
*/
func main(){
	s := `<?xml version='1.0' encoding='UTF-8'?>
	<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
	<S:Body>
	<ns2:createStorageReturn xmlns:ns2="http://ws.api.mysite.com/">
	<return>
	xxxxx
	</return>
	</ns2:createStorageReturn>
	</S:Body>
	</S:Envelope>`
	//<requestId>16660663</requestId><dataCenterId>ssrr-444tt-yy-99</dataCenterId><dataCenterVersion>12</dataCenterVersion><storageId>towrrt24903FR55405</storageId>

	parser := xml.NewDecoder(bytes.NewBufferString(s))
	envelope := new(StorageResponse)
	err := parser.DecodeElement(&envelope, nil)
	if err != nil {
		fmt.Printf("Error in parsing")
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("%+v", envelope)
}
