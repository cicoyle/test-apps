module test-apps/scheduler-actor-reminders/player-actor

go 1.23.1

require (
	github.com/dapr/go-sdk v1.11.0
	test-apps/scheduler-actor-reminders/api v0.0.0-00010101000000-000000000000
)

require (
	github.com/dapr/dapr v1.14.4 // indirect
	github.com/dapr/go-sdk/examples/actor v0.0.0-20240626135542-c417f950fe1d // indirect
	github.com/go-chi/chi/v5 v5.1.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	go.opentelemetry.io/otel v1.27.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240624140628-dc46fd24d27d // indirect
	google.golang.org/grpc v1.65.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace test-apps/scheduler-actor-reminders/api => ../api

replace github.com/dapr/go-sdk => /Users/cassie/go/src/github.com/go-sdk
