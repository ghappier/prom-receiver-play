package pool

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var MaxLength int64 = 102400

/*
type Producter struct {
	job Job
}

func NewProducter(maxJob int) *Producter {
	job := Job{Payload: Payload(maxJob)}
	return &Producer{job: job}
}
*/

func PayloadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PayloadHandler")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Read the body into a string for json decoding
	var content = &PayloadCollection{}
	err := json.NewDecoder(io.LimitReader(r.Body, MaxLength)).Decode(&content)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("error")
		return
	}
	fmt.Println("aa")
	// Go through each payload and queue items individually to be posted to S3
	for _, payload := range content.Payloads {
		fmt.Println(payload)
		// let's create a job with the payload
		work := Job{Payload: payload}

		// Push the work onto the queue.
		JobQueue <- work
	}

	w.WriteHeader(http.StatusOK)
}
