package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"test-apps/scheduler-actor-reminders/api"
	"time"

	dapr "github.com/dapr/go-sdk/client"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Actor ID for the player 'session'
	actorID := "player-1"

	getPlayerRequest := &api.GetPlayerRequest{ActorID: actorID}
	requestData, err := json.Marshal(getPlayerRequest)
	if err != nil {
		log.Fatalf("error marshaling request data: %v", err)
	}

	req := &dapr.InvokeActorRequest{
		ActorType: "playerActorType",
		ActorID:   actorID,
		Method:    "GetUser",
		Data:      requestData,
	}
	resp, err := client.InvokeActor(ctx, req)
	if err != nil {
		log.Fatalf("error invoking actor method GetUser: %v", err)
	}
	fmt.Printf("cassie resp: %+v\n", string(resp.Data))

	err = client.RegisterActorReminder(ctx, &dapr.RegisterActorReminderRequest{
		ActorType: "playerActorType",
		ActorID:   actorID,
		Name:      "healthReminder",
		DueTime:   "1s",
		Period:    "2s",
		Data:      []byte(`"Reminder triggered"`),
	})
	if err != nil {
		log.Fatalf("error starting actor reminder: %v", err)
	}
	fmt.Println("Started healthReminder for actor:", actorID)

	fmt.Println("Waiting 5s...")
	<-time.After(time.Second * 5) // allow time for reminders to trigger

	fmt.Println("Getting actor: ", actorID)
	req = &dapr.InvokeActorRequest{
		ActorType: "playerActorType",
		ActorID:   actorID,
		Method:    "GetUser",
		Data:      requestData,
	}
	resp, err = client.InvokeActor(ctx, req)
	if err != nil {
		log.Fatalf("error invoking actor method GetUser: %v", err)
	}
	fmt.Printf("cassie resp1: %+v\n", string(resp.Data))

	fmt.Println("Unregistering healthReminder for actor...")
	err = client.UnregisterActorReminder(ctx, &dapr.UnregisterActorReminderRequest{
		ActorType: "playerActorType",
		ActorID:   actorID,
		Name:      "healthReminder",
	})
	if err != nil {
		log.Fatalf("error unregistering actor reminder: %v", err)
	}
	fmt.Println("Unregistered healthReminder for actor:", actorID)

	// Graceful shutdown on Ctrl+C or SIGTERM (for Docker/K8s graceful shutdown)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	fmt.Println("Shutting down...")
}
