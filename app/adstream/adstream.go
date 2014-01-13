package adstream

import (
	"fmt"
	"encoding/base64"
	"code.google.com/p/go.net/websocket"
	"github.com/mati1979/go-revel-mobile-cars-adstream/app/xmlcodec"
	"github.com/robfig/revel"
	"os"
	"container/list"
	"time"
)

type AdEvent struct {
	AdId string
	ZipCode string
	Timestamp int    // Unix timestamp (secs)
	Lat float32
	Lon float32
}

type Subscription struct {
	New <- chan AdEvent
}

var (
	subscribe = make(chan (chan <- Subscription), 10)
	unsubscribe = make(chan (<-chan AdEvent), 10)
	publish = make(chan AdEvent, 10)
)

func Subscribe() Subscription {
	resp := make(chan Subscription)
	subscribe <- resp
	return <-resp
}

// Owner of a subscription must cancel it when they stop listening to events.
func (sub Subscription) Cancel() {
	unsubscribe <- sub.New // Unsubscribe the channel.
	drain(sub.New) // Drain it, just in case there was a pending publish.
}

func AdStream() {
	subscribers := list.New()
	for {
		select {
		case ch := <-subscribe:
			subscriber := make(chan AdEvent, 10)
			subscribers.PushBack(subscriber)
			ch <- Subscription{subscriber}

		case unsub := <-unsubscribe:
			for ch := subscribers.Front(); ch != nil; ch = ch.Next() {
				if ch.Value.(chan AdEvent) == unsub {
					subscribers.Remove(ch)
					break
				}
			}
		case event := <-publish:
			for ch := subscribers.Front(); ch != nil; ch = ch.Next() {
				ch.Value.(chan AdEvent) <- event
			}
		}
	}
}

func Connect() {
	user, foundLogin := revel.Config.String("adstream.user")
	pass, foundPass := revel.Config.String("adstream.pass")

	if !foundLogin || !foundPass {
		fmt.Printf("Need credentials")
	}

	url := fmt.Sprintf("ws://%s/mobile-ad-stream/websocket/events", "adstream.mobile.de:80")
	config, err := websocket.NewConfig(url, url)
	if err != nil {
		fmt.Printf("Config failed: %s\n", err.Error())
		os.Exit(1)
	}

	login := user + ":" + pass
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(login))
	config.Header.Add("Authorization", auth)

	ws, err := websocket.DialConfig(config)
	if err != nil {
		fmt.Printf("Dial failed: %s\n", err.Error())
		os.Exit(1)
	}

	for {
		var event xmlcodec.AdEvent
		err := xmlcodec.XMLCodec.Receive(ws, &event)
		if err != nil {
			fmt.Println(fmt.Sprintf("error4:%s", err))
			os.Exit(1)
		}
		if event.EventType == "AD_CREATE_OR_UPDATE" {
			adData := AdEvent{event.Ad.AdKey, event.Ad.Seller.SellerAddress.SellerZipCode.ZipCode, int(time.Now().Unix()), 0, 0}
			publish <- adData
		}
	}
}

// Drains a given channel of any messages.
func drain(ch <-chan AdEvent) {
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return
			}
		default:
			return
		}
	}
}

func Init() {
	go Connect()
	go AdStream()
}
