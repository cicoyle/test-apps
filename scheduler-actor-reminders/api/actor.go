package api

import (
	"context"
	"fmt"

	"github.com/dapr/go-sdk/actor"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/examples/actor/api"
)

const playerActorType = "playerActorType"

type PlayerActor struct {
	actor.ServerImplBaseCtx
	DaprClient dapr.Client
	Health     int
}

func (p *PlayerActor) Type() string {
	return playerActorType
}

type GetPlayerRequest struct {
	ActorID string
}

// GetUser retrieving the state of the PlayerActor
func (p *PlayerActor) GetUser(ctx context.Context, player *GetPlayerRequest) (*PlayerActor, error) {
	if player.ActorID == p.ID() {
		fmt.Printf("Player Actor ID: %s has a health level of: %d\n", p.ID(), p.Health)
	}
	return p, nil
}

// Invoke invokes an action on the actor
func (p *PlayerActor) Invoke(ctx context.Context, req string) (string, error) {
	fmt.Println("get req = ", req)
	return req, nil
}

// ReminderCall executes logic to handle what happens when the reminder is triggered
// Dapr automatically calls this method when a reminder fires for the player actor
func (p *PlayerActor) ReminderCall(reminderName string, state []byte, dueTime string, period string) {
	fmt.Println("receive reminder = ", reminderName, " state = ", string(state), "duetime = ", dueTime, "period = ", period)
	if reminderName == "healthReminder" {
		fmt.Println("cassieeeeee")
		p.Health += 10 // Restore some health
		fmt.Printf("Player health restored. Current health: %d\n", p.Health)
	}

}

// StartReminder registers a reminder for the actor
func (p *PlayerActor) StartReminder(ctx context.Context, req *api.ReminderRequest) error {
	fmt.Println("Starting reminder:", req.ReminderName)
	return p.DaprClient.RegisterActorReminder(ctx, &dapr.RegisterActorReminderRequest{
		ActorType: p.Type(),
		ActorID:   p.ID(),
		Name:      req.ReminderName,
		DueTime:   req.Duration,
		Period:    req.Period,
		Data:      []byte(req.Data),
	})
}

// Start initializes the actor and its reminders
func (p *PlayerActor) Start(ctx context.Context) error {
	fmt.Println("PlayerActor started.")

	err := p.StartReminder(ctx, &api.ReminderRequest{
		ReminderName: "healthReminder",
		Period:       "10s",
		Duration:     "10s",
		Data:         `"Reminder triggered"`,
	})
	if err != nil {
		return fmt.Errorf("failed to start reminder: %w", err)
	}

	return nil
}
