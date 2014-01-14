package adstream

import (
	"fmt"
	"encoding/base64"
	"code.google.com/p/go.net/websocket"
	"github.com/robfig/revel"
	"os"
	"container/list"
	"io"
	"github.com/mati1979/go-revel-mobile-cars-adstream/app/xmlcodec"
	"strconv"
	"time"
)

type AdEvent struct {
	AdId string
	ZipCode string
	Timestamp int    // Unix timestamp (secs)
	Lat float64
	Lon float64
}

type Subscription struct {
	New <- chan AdEvent
	Archive []AdEvent
}

const archiveSize = 100

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
	archive := list.New()
	subscribers := list.New()
	for {
		select {
		case ch := <-subscribe:
			var events []AdEvent
			for e := archive.Front(); e != nil; e = e.Next() {
				events = append(events, e.Value.(AdEvent))
			}
			subscriber := make(chan AdEvent, 10)
			subscribers.PushBack(subscriber)
			ch <- Subscription{subscriber, events}

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
			if archive.Len() >= archiveSize {
				archive.Remove(archive.Front())
			}
			archive.PushBack(event)
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
	}

	for {
		var event xmlcodec.AdEvent
		err := xmlcodec.XMLCodec.Receive(ws, &event)
		if err != nil {
			if (err == io.EOF) {
				fmt.Println(fmt.Sprintf("timeout:%s", err.Error()))
			} else {
				fmt.Println(fmt.Sprintf("error4:%s", err.Error()))
			}
		}
		if event.EventType == "AD_CREATE_OR_UPDATE" {
			Id := event.Ad.AdKey
			ZipCode := event.Ad.Seller.SellerAddress.SellerZipCode.ZipCode
			Time := int(time.Now().Unix())
			if (event.Ad.Seller.SellerCoords != nil) {
				Lat := ParseF(&event.Ad.Seller.SellerCoords.Latitude)
				Lon := ParseF(&event.Ad.Seller.SellerCoords.Longitude)
				AdEven := AdEvent{Id, ZipCode, Time, Lat, Lon}
				publish <- AdEven
			} else {
				AdEven := AdEvent{Id, ZipCode, Time, 0, 0}
				publish <- AdEven
			}
		}
	}
}

func ParseF(s *string) float64 {
	f, err := strconv.ParseFloat(*s, 12)
	if (err != nil) {
		fmt.Println(fmt.Sprintf("ParseError:%s", err.Error()))
	}

	return f
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
