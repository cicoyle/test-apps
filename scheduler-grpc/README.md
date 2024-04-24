1. go mod init scheduler-grpc
2. go mod tidy
3. protoc --go_out=./job --go-grpc_out=./job proto/job.proto
4. go run scheduler-grpc.go