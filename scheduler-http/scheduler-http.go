package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Job struct {
	TypeURL string `json:"type_url"`
	Value   string `json:"value"`
}

func main() {
	// List
	http.HandleFunc("/job/", handleJob)
	// Get
	http.HandleFunc("/job/cass1", cass1Job)
	http.HandleFunc("/job/sam1", sam1Job)

	log.Println("Server started on port 7979")
	log.Fatal(http.ListenAndServe("127.0.0.1:7979", nil))
}

func sam1Job(w http.ResponseWriter, r *http.Request) {
	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		// TODO
	}
	fmt.Println("SAM: Raw request body:", string(rawBody))
}

func cass1Job(w http.ResponseWriter, r *http.Request) {
	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		// TODO
	}
	fmt.Println("CASSIE: Raw request body:", string(rawBody))
}

func handleJob(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CASSIE IN HANDLE JOB")
	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("error reading request body: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Println("Raw request body:", string(rawBody))

	var jobData Job
	if err := json.Unmarshal(rawBody, &jobData); err != nil {
		http.Error(w, fmt.Sprintf("error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}

	decodedValue, err := base64.StdEncoding.DecodeString(jobData.Value)
	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding base64: %v", err), http.StatusBadRequest)
		return
	}

	fmt.Printf("Decoded JOB value: %s\n", decodedValue)

	w.WriteHeader(http.StatusOK)
}
