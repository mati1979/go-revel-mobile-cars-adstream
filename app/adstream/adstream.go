package adstream

import (
	"fmt"
	"encoding/base64"
	"code.google.com/p/go.net/websocket"
	"github.com/mati1979/go-revel-mobile-cars-adstream/app/xmlcodec"
	"github.com/robfig/revel"
)

type AdData struct {
	AdId string
	ZipCode string
}

func Connect() {
	user, foundLogin := revel.Config.String("adstream.user")
	pass, foundPass := revel.Config.String("adstream.pass")

	if !foundLogin || !foundPass {
	}

	url := fmt.Sprintf("ws://%s/mobile-ad-stream/websocket/events", "adstream.mobile.de:80")
	config, err := websocket.NewConfig(url, "http://localhost/")
	fmt.Println("connect")
	if err != nil {
		fmt.Println(fmt.Sprintf("error1:%s", err))
		return
	}
	login := user + ":" + pass
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(login))
	config.Header.Add("Authorization", auth)

	ws, err := websocket.DialConfig(config)
	if err != nil {
		fmt.Println(fmt.Sprintf("error2:%s", err))
		return
	}

	for {
		var event xmlcodec.AdEvent
		err := xmlcodec.XMLCodec.Receive(ws, &event)
		if err != nil {
			fmt.Println(fmt.Sprintf("error4:%s", err))
			return
		}
		if event.EventType == "AD_CREATE_OR_UPDATE" {
			adData := AdData{event.Ad.AdKey, event.Ad.Seller.SellerAddress.SellerZipCode.ZipCode}
			fmt.Println(fmt.Sprintf("ad-data:%s", adData))
		}
	}
}

func Init() {
	go Connect()
}
