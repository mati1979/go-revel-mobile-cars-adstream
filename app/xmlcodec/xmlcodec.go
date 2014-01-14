package xmlcodec

import (
	"encoding/xml"
	"code.google.com/p/go.net/websocket"
)

type AdEvent struct {
	XMLName xml.Name `xml:"http://services.mobile.de/schema/search event,omitempty"`
	EventType string `xml:"event-type,omitempty"`
	Ad *Ad `xml:"ad,omitempty"`
}

type Ad struct {
	XMLName xml.Name `xml:"http://services.mobile.de/schema/ad ad,omitempty"`
	AdKey string `xml:"key,attr,omitempty"`
	Seller *Seller `xml:"seller,omitempty"`
}

type Seller struct {
	XMLName xml.Name `xml:"http://services.mobile.de/schema/seller seller,omitempty"`
	SellerKey string `xml:"key,attr,omitempty"`
	SellerAddress *SellerAddress `xml:"address,omitempty"`
	SellerCoords *SellerCoords `xml:"coordinates,omitempty"`
}

type SellerAddress struct {
	XMLName xml.Name `xml:"http://services.mobile.de/schema/seller address,omitempty"`
	SellerZipCode *SellerZipCode `xml:"zipcode,omitempty"`
	SellerCountryCode *SellerCountryCode `xml:"country-code,omitempty"`
}

type SellerCoords struct {
	XMLName xml.Name `xml:"http://services.mobile.de/schema/seller coordinates,omitempty"`
	Latitude string `xml:"latitude,omitempty"`
	Longitude string `xml:"longitude,omitempty"`
}

type SellerZipCode struct {
	XMLName xml.Name `xml:"http://services.mobile.de/schema/seller zipcode,omitempty"`
	ZipCode string `xml:"value,attr,omitempty"`
}

type SellerCountryCode struct {
	XMLName xml.Name `xml:"http://services.mobile.de/schema/seller country-code,omitempty"`
	CountryCode string `xml:"value,attr,omitempty"`
}

func xmlMarshal(v interface{}) (msg []byte, payloadType byte, err error) {
	//buff := &bytes.Buffer{}
	msg, err = xml.Marshal(v)
	//msgRet := buff.Bytes()
	return msg, websocket.TextFrame, nil
}

func xmlUnmarshal(msg []byte, payloadType byte, v interface{}) (err error) {
	// r := bytes.NewBuffer(msg)
	//fmt.Println("eventXml:" + string(msg[:len(msg)]))
	err = xml.Unmarshal(msg, v)
	return err
}

var XMLCodec = websocket.Codec{xmlMarshal, xmlUnmarshal}
