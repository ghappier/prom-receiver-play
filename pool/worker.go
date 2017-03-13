package pool

import (
	//"log"
	//"os"
	"fmt"

	"github.com/lizq/prom-receiver-play/storage"
)

var (
	//MaxWorker = os.Getenv("MAX_WORKERS")
	//MaxQueue  = os.Getenv("MAX_QUEUE")
	MaxWorker = 100
	MaxQueue  = 100000

	// A buffered channel that we can send work requests on.
	JobQueue chan Job = make(chan Job, 1)
)

// Job represents the job to be run
type Job struct {
	Payload Payload
}

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
	holder     *storage.StorageHolder
}

func NewWorker(workerPool chan chan Job, h *storage.StorageHolder) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
		holder:     h,
	}
}

// Start method starts the run loop for the worker, listening for a quit channel
// in case we need to stop it
func (w Worker) Start() {

	go func() {

		for {
			//register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// we have received a work request.
				job.Payload.holder = w.holder
				if err := job.Payload.Save(); err != nil {
					fmt.Printf("Error do payload function ：%s", err.Error())
				}
			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
