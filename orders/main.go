package main

import (
	"context"
	dapr "github.com/dapr/go-sdk/client"
	"log"
	"strconv"
	"time"
)

var _client dapr.Client

func main() {
	var err error
	_client, err = dapr.NewClient()
	if err != nil {
		log.Fatalf("error creating dapr client: %v", err)
	}
	defer _client.Close()

	for i := 1; i <= 10; i++ {
		order := `{"orderId"` + strconv.Itoa(i) + `}`

		err := _client.PublishEvent(context.Background(), "pubsub", "order.created", []byte(order))
		if err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}
}
