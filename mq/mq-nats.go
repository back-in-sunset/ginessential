package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/nats-io/stan.go/pb"
)

var (
	clusterID string = "test-cluster"
	clientID  string = "9993"
	natsURL   string = "nats://127.0.0.1:4222"
)

// PubSub ..
func PubSub() {
	go Publisher()
	go Subscriber1()
	go Subscriber2()
}

// PubSubWithReply ..
func PubSubWithReply() {
	Subscriber()
	PublisherReply()
}

// Publisher ..
func Publisher() {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}

	sc, err := stan.Connect(clusterID, clientID, stan.NatsConn(nc))
	if err != nil {
		log.Fatal(err)
		return
	}

	// 开启一个协程，不停的生产数据
	go func() {
		m := 0
		for {
			m++
			sc.Publish("foo1", []byte("hello message "+strconv.Itoa(m)))
			time.Sleep(time.Second)
		}

	}()

}

// Subscriber1 ..
func Subscriber1() {
	// 消费数据
	sc, err := stan.Connect(clusterID, "subscriber1", stan.NatsURL(natsURL))
	if err != nil {
		log.Fatal(err)
		return
	}

	i := 0
	mcb := func(msg *stan.Msg) {
		i++
		log.Println("[INFO] Subscriber1 subscribe:", i, "---->", msg.Subject, msg)
	}
	startOpt := stan.StartAt(pb.StartPosition_LastReceived)
	//_, err = sc.QueueSubscribe("foo1", "", mcb, startOpt)   // 也可以用queue subscribe
	_, err = sc.Subscribe("foo1", mcb, startOpt)
	if err != nil {
		sc.Close()
		log.Fatal(err)
	}

	// 创建一个channel，阻塞着
	signalChan := make(chan int)
	<-signalChan
}

// Subscriber2 ..
func Subscriber2() {
	// 消费数据
	sc, err := stan.Connect(clusterID, "subscriber2", stan.NatsURL(natsURL))
	if err != nil {
		log.Fatal(err)
		return
	}

	i := 0
	mcb := func(msg *stan.Msg) {
		i++
		log.Println("[INFO] Subscriber2 subscribe:", i, "---->", msg.Subject, msg)
	}
	startOpt := stan.StartAt(pb.StartPosition_LastReceived)
	//_, err = sc.QueueSubscribe("foo1", "", mcb, startOpt)   // 也可以用queue subscribe
	_, err = sc.Subscribe("foo1", mcb, startOpt, stan.DurableName("foo1-durable"))
	if err != nil {
		sc.Close()
		log.Fatal(err)
	}

	// 创建一个channel，阻塞着
	signalChan := make(chan int)
	<-signalChan
}

// Subscriber ..
func Subscriber() {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}

	// uniqueReplyTo := nats.NewInbox()
	subj := "time"
	sub, err := nc.SubscribeSync(subj)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			// Read the reply
			msg, err := sub.NextMsg(3 * time.Second)
			if err != nil {
				log.Fatal(err)
			}
			// Use the response
			log.Printf("Data:%s Reply: %s", msg.Data, msg.Reply)

			// Get the time
			timeAsBytes := []byte(time.Now().String())

			// Send the time as the response.
			msg.Respond(timeAsBytes)
		}
	}()

}

// PublisherReply ..
func PublisherReply() {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		m := 0
		for {
			m++
			nc.PublishRequest("time", "test_reply", []byte(fmt.Sprintf("msg%d----->", m)))
			time.Sleep(time.Second)
		}
	}()
}

// SubscriberAck ..
func SubscriberAck() {
	sc, err := stan.Connect(clusterID, "subscriber_ack", stan.NatsURL(natsURL))
	if err != nil {
		log.Fatal(err)
		return
	}

	i := 0
	_, err = sc.Subscribe("foo1",
		func(m *stan.Msg) {
			i++
			log.Println("[INFO] SubscriberAck subscribe:", i, "---->", m.Subject, m)
			time.Sleep(2 * time.Second)
			m.Ack()
		}, stan.SetManualAckMode(), stan.AckWait(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	// PubSubWithReply()
	// // 创建一个channel，阻塞着
	// signalChan := make(chan int)
	// <-signalChan

	SubscriberAck()
	Publisher()
	Subscriber1()
	// 创建一个channel，阻塞着
	signalChan := make(chan int)
	<-signalChan

}
