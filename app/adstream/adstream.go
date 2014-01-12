package adstream

import (
	"fmt"
	"code.google.com/p/go.net/websocket"
	"encoding/base64"
	"encoding/xml"
	"github.com/mati1979/go-revel-mobile-cars-adstream/app/xmlcodec"
)

type AdEvent struct {
	XMLName xml.Name `xml:"http://services.mobile.de/schema/search event"`
	typ string `xml:"http://services.mobile.de/schema/search event-type"`
	ad Ad `xml:"http://services.mobile.de/schema/ad ad"`
}

type Ad struct {
	XMLName xml.Name `xml:"http://services.mobile.de/schema/ad ad"`
	key string `xml:"ad-key,attr"`
}

func Connect() {
	url := fmt.Sprintf("ws://%s/mobile-ad-stream/websocket/events", "adstream.mobile.de:80")
	config, err := websocket.NewConfig(url, "http://localhost/")
	fmt.Println("connect")
	if err != nil {
		fmt.Println(fmt.Sprintf("error1:%s", err))
		return
	}
	login := "xxx" + ":" + "xxx"
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(login))
	config.Header.Add("Authorization", auth)

	ws, err := websocket.DialConfig(config)
	if err != nil {
		fmt.Println(fmt.Sprintf("error2:%s", err))
		return
	}

	for {
		var event AdEvent
		err := xmlcodec.XMLCodec.Receive(ws, &event)
		if err != nil {
			fmt.Println(fmt.Sprintf("error4:%s", err))
			return
		}
		fmt.Println(fmt.Sprintf("event:%s", event))
	}

}

func Init() {
	go Connect()
}
