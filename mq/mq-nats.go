package main

import (
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
	go Consumer1()
	go Consumer2()
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

// Consumer1 ..
func Consumer1() {
	// 消费数据
	sc, err := stan.Connect(clusterID, "consumer1", stan.NatsURL(natsURL))
	if err != nil {
		log.Fatal(err)
		return
	}

	i := 0
	mcb := func(msg *stan.Msg) {
		i++
		log.Println("[INFO] Consumer1 consume:", i, "---->", msg.Subject, msg)
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

// Consumer2 ..
func Consumer2() {
	// 消费数据
	sc, err := stan.Connect(clusterID, "consumer2", stan.NatsURL(natsURL))
	if err != nil {
		log.Fatal(err)
		return
	}

	i := 0
	mcb := func(msg *stan.Msg) {
		i++
		log.Println("[INFO] Consumer2 consume:", i, "---->", msg.Subject, msg)
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
