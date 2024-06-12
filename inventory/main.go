package main

import (
	"context"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"log"
	"net/http"
)

// Subscription to tell the dapr what topic to subscribe.
//   - PubsubName: is the name of the component configured in the metadata of pubsub.yaml.
//   - Topic: is the name of the topic to subscribe.
//   - Route: tell dapr where to request the API to publish the message to the subscriber when get a message from topic.
//   - Match: (Optional) The CEL expression to match on the CloudEvent to select this route.
//   - Priority: (Optional) The priority order of the route when Match is specificed.
//     If not specified, the matches are evaluated in the order in which they are added.
var _defaultSubscription = &common.Subscription{
	PubsubName: "pubsub",
	Topic:      "order.created",
	Route:      "/orders",
}

func main() {
	s := daprd.NewService(":8001")

	if err := s.AddTopicEventHandler(_defaultSubscription, handler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listening: %v", err)
	}
}

func handler(_ context.Context, e *common.TopicEvent) (bool, error) {
	log.Printf("event - PubsubName: %s, Topic: %s, ID: %s, Data: %s", e.PubsubName, e.Topic, e.ID, e.Data)
	return false, nil
}
