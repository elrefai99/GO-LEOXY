package Queue

import (
	"fmt"
	"sync"
	"time"
)

type IJob struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}

type Queue struct {
	jobs   chan IJob
	nextID int
	mu     sync.Mutex
}

func LeoxyWorker(size int) *Queue {
	return &Queue{
		jobs: make(chan IJob, size),
	}
}

func (q *Queue) Add(j IJob) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.nextID++
	j.ID = q.nextID

	q.jobs <- j
}

func (q *Queue) Worker(id int) {
	for jobs := range q.jobs {
		fmt.Printf(
			"Worker %d processing job %d type=%s\n",
			id,
			jobs.ID,
			jobs.Type,
		)

		switch jobs.Type {
		case "email":
			SendEmail(jobs)

			time.Sleep(2 * time.Second)
		}
	}
}

func SendEmail(jobs IJob) {
	fmt.Printf(
		"Sending email: %s\n",
		jobs.Payload,
	)
}
