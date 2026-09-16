package JobQueue

import (
	"fmt"

	"github.com/elrefai99/go-backend/app/Queue/model"
)

func SendEmail(jobs model.IJob) {
	fmt.Printf(
		"Sending email: %s\n",
		jobs.Payload,
	)
}
