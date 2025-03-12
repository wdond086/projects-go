package task

import (
	"fmt"
)

type Status struct {
	text string
}

var (
	Pending    = Status{"Pending"}
	InProgress = Status{"InProgress"}
	Done       = Status{"Completed"}
	Cancelled  = Status{"Cancelled"}
)

var statuses = []Status{
	Pending,
	InProgress,
	Done,
	Cancelled,
}

func StatusFromString(statusStr string) (Status, error) {
	for _, status := range statuses {
		if status.text == statusStr {
			return status, nil
		}
	}
	return Status{}, fmt.Errorf("Unknown '%s' status", statusStr)
}

func (status Status) Validate() error {
	voidStatus := Status{}
	if status == voidStatus {
		return fmt.Errorf("Invalid Status")
	}
	return nil
}

func (status Status) String() string {
	if err := status.Validate(); err != nil {
		return ""
	}
	return status.text
}
