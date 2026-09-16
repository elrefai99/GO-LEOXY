package Queue

import (
	"fmt"
	"time"
)

type Job struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

type Queue struct {
	jobs chan Job
}

func LeoxyWorker(size int) *Queue {
	return &Queue{
		jobs: make(chan Job, size),
	}
}

func (q *Queue) Add(j Job) {
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
			fmt.Printf(
				"Sending email: %s\n",
				jobs.Payload,
			)

			time.Sleep(2 * time.Second)
		}
	}
}
