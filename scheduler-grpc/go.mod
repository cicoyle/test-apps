module scheduler-grpc

go 1.22.2

toolchain go1.22.3

require (
	github.com/dapr/dapr v1.13.2
	google.golang.org/grpc v1.63.2
	google.golang.org/protobuf v1.33.0
)

require (
	go.opentelemetry.io/otel v1.21.0 // indirect
	go.opentelemetry.io/otel/trace v1.21.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240304212257-790db918fca8 // indirect
)

replace github.com/dapr/dapr => ../../dapr
