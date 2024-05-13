package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"

	commonv1 "github.com/dapr/dapr/pkg/proto/common/v1"
	rtv1 "github.com/dapr/dapr/pkg/proto/runtime/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"scheduler-grpc/job/github.com/cicoyle/test-apps/scheduler-grpc/job"
)

// JobService represents the gRPC service
type JobService struct {
	job.JobServiceServer
}
type Job struct {
	TypeURL string `json:"type_url"`
	Value   string `json:"value"`
}

// WatchJobs implements the gRPC method for watching jobs
func (s *JobService) WatchJobs(ctx context.Context, req *job.WatchJobsRequest) (*job.WatchJobsResponse, error) {
	rawData := req.GetData()
	log.Printf("Raw request data: %s\n", rawData)

	decodedValue, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		return nil, fmt.Errorf("error decoding base64: %v", err)
	}

	log.Printf("Decoded JOB value: %s\n", decodedValue)
	return &job.WatchJobsResponse{}, nil
}

func (s *JobService) ReceiveJobs(ctx context.Context, in *job.ReceiveJobsRequest) (*job.ReceiveJobsResponse, error) {
	if strings.HasPrefix(in.GetMethod(), "job/cass") {
		log.Printf("Method: %s\n", in.GetMethod())
		log.Printf("Data: %s\n", in.GetData())
		var jobData Job
		//dataBytes := in.GetData()
		//
		//if err := json.Unmarshal(dataBytes, &jobData); err != nil {
		//	return nil, fmt.Errorf("error decoding data: %v", err)
		//}
		decodedValue, err := base64.StdEncoding.DecodeString(jobData.Value)
		if err != nil {
			return nil, fmt.Errorf("error decoding base64: %v", err)
		}

		log.Printf("Decoded value: %s", strings.TrimSpace(string(decodedValue)))
		log.Println()
	}
	return nil, nil
}

func (s *JobService) OnJobEvent(ctx context.Context, in *rtv1.JobEventRequest) (*rtv1.JobEventResponse, error) {
	log.Printf("CASSIE IN ONJOBEVENT: %+v\n", in)

	log.Printf("Method: %s\n", in.GetMethod())
	log.Printf("Data: %s\n", in.GetData().String())

	//in.GetName() -> enable users to do switch off job name and/or method
	if strings.HasPrefix(in.GetMethod(), "job/cass") {
		log.Printf("Method: %s\n", in.GetMethod())
		log.Printf("Data: %s\n", in.GetData().String())
		var jobData Job
		dataBytes := in.GetData().GetValue()

		if err := json.Unmarshal(dataBytes, &jobData); err != nil {
			return nil, fmt.Errorf("error decoding data: %v", err)
		}
		decodedValue, err := base64.StdEncoding.DecodeString(jobData.Value)
		if err != nil {
			return nil, fmt.Errorf("error decoding base64: %v", err)
		}

		log.Printf("Decoded value: %s", strings.TrimSpace(string(decodedValue)))
		log.Println()
	}
	return nil, nil
}

func (s *JobService) OnInvoke(ctx context.Context, in *commonv1.InvokeRequest) (*commonv1.InvokeResponse, error) {
	if strings.HasPrefix(in.GetMethod(), "job/cass") {
		log.Printf("Method: %s\n", in.GetMethod())
		log.Printf("Data: %s\n", in.GetData().String())
		var jobData Job
		dataBytes := in.GetData().GetValue()

		if err := json.Unmarshal(dataBytes, &jobData); err != nil {
			return nil, fmt.Errorf("error decoding data: %v", err)
		}
		decodedValue, err := base64.StdEncoding.DecodeString(jobData.Value)
		if err != nil {
			return nil, fmt.Errorf("error decoding base64: %v", err)
		}

		log.Printf("Decoded value: %s", strings.TrimSpace(string(decodedValue)))
		log.Println()
	}
	return nil, nil
}

func (s *JobService) OnBindingEvent(ctx context.Context, in *rtv1.BindingEventRequest) (*rtv1.BindingEventResponse, error) {
	return nil, nil
}

func (s *JobService) ListTopicSubscriptions(context.Context, *emptypb.Empty) (*rtv1.ListTopicSubscriptionsResponse, error) {
	return nil, nil
}

func (s *JobService) OnTopicEvent(ctx context.Context, in *rtv1.TopicEventRequest) (*rtv1.TopicEventResponse, error) {
	log.Printf("Raw request data: %+v\n", in)
	return nil, nil
}

func (s *JobService) ListInputBindings(context.Context, *emptypb.Empty) (*rtv1.ListInputBindingsResponse, error) {
	return nil, nil
}

func main() {
	server := grpc.NewServer()
	js := &JobService{}
	// Register JobService with the server
	job.RegisterJobServiceServer(server, js)
	// Register the JobService with Dapr for handling callbacks
	rtv1.RegisterAppCallbackServer(server, js)
	listener, err := net.Listen("tcp", "127.0.0.1:7878")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Start the server
	log.Println("Server started on port 7878")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
