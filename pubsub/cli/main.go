package main

import (
	"fmt"
	"time"

	"context"

	proto "github.com/go-micro/examples/pubsub/srv/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/util/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// send events using the publisher
func sendEv(topic string, p micro.Publisher) {
	t := time.NewTicker(time.Second)

	for _ = range t.C {
		// create new event
		ev := &proto.Notification{
			Event: &proto.Event{
				Time: timestamppb.Now(),
				Text: fmt.Sprintf("Messaging you all day on %s", topic),
			},
			Kind: proto.Notification_DownloadComplete,
		}

		log.Logf("publishing %+v\n", ev)

		// publish an event
		if err := p.Publish(context.Background(), ev); err != nil {
			log.Logf("error publishing: %v", err)
		}
	}
}

func main() {
	// create a service
	service := micro.NewService(
		micro.Name("go.micro.cli.pubsub"),
	)
	// parse command line
	service.Init()

	// create publisher
	pub1 := micro.NewEvent("example.topic.pubsub.1", service.Client())
	pub2 := micro.NewEvent("example.topic.pubsub.2", service.Client())

	// pub to topic 1
	go sendEv("example.topic.pubsub.1", pub1)
	// pub to topic 2
	go sendEv("example.topic.pubsub.2", pub2)

	// block forever
	select {}
}
