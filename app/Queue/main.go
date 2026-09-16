package Queue

import (
	"errors"
	"fmt"
	"sync"

	JobQueue "github.com/elrefai99/go-backend/app/Queue/job"
	"github.com/elrefai99/go-backend/app/Queue/model"
)

type Queue struct {
	jobs chan model.IJob
	stop chan struct{}

	nextID int
	closed bool

	mu   sync.Mutex
	wg   sync.WaitGroup
	once sync.Once
}

var returnError = errors.New("Queue is stopped")

func LeoxyWorker(size int) *Queue {
	q := &Queue{
		jobs: make(chan model.IJob, size),
		stop: make(chan struct{}),
	}
	q.Start(3)
	return q
}

func (q *Queue) Start(size int) {
	for i := 1; i <= size; i++ {
		q.wg.Add(1)
		go q.Worker(i)
	}
}

func (q *Queue) Add(j model.IJob) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.closed {
		return returnError
	}

	q.nextID++
	j.ID = q.nextID

	select {
	case q.jobs <- j:
		return nil
	case <-q.stop:
		return returnError
	}
}

func (q *Queue) Worker(id int) {
	defer q.wg.Done()
	for {
		select {
		case <-q.stop:
			return

		case job := <-q.jobs:
			if err := q.process(id, job); err != nil {
				fmt.Printf("Worker %d failed: %v\n", id, err)
				q.Stop()
				return
			}
		}
	}

}
func (q *Queue) process(workerID int, job model.IJob) error {
	fmt.Printf(
		"Worker %d processing job %d type=%s\n",
		workerID,
		job.ID,
		job.Type,
	)

	switch job.Type {
	case "email":
		JobQueue.SendEmail(job)
		return nil

	default:
		return fmt.Errorf("unknown job type: %s", job.Type)
	}
}

func (q *Queue) Stop() {
	q.once.Do(func() {
		q.mu.Lock()
		q.closed = true
		q.mu.Unlock()

		close(q.stop)
	})
}

func (q *Queue) Close() {
	q.Stop()
	q.wg.Wait()
}
