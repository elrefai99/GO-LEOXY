package Queue

import (
	"fmt"
	"sync"
	"time"

	JobQueue "github.com/elrefai99/go-backend/app/Queue/job"
	"github.com/elrefai99/go-backend/app/Queue/model"
)

type Queue struct {
	jobs   chan model.IJob
	nextID int
	mu     sync.Mutex
}

func LeoxyWorker(size int) *Queue {
	return &Queue{
		jobs: make(chan model.IJob, size),
	}
}

func (q *Queue) Add(j model.IJob) {
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
			JobQueue.SendEmail(jobs)

			time.Sleep(2 * time.Second)
		}
	}
}
