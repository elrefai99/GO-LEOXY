package queue

import (
	"fmt"

	"github.com/elrefai99/go-backend/app/Queue"
)

func SendEmail(jobs Queue.Job) {
	fmt.Printf(
		"Sending email: %s\n",
		jobs.Payload,
	)
}
