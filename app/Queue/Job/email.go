package queue

import "sync"

type EmailJob struct {
	ID      int
	To      string
	Subject string
	Body    string
}

type EmailQueue struct {
	jobs chan EmailJob
	wg   sync.WaitGroup
}
