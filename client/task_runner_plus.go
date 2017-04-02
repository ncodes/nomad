package client

import (
	"fmt"
	"log"
	"time"

	"github.com/ncodes/cocoon/tools"
)

// TaskRunnerPlus includes unofficial task runner functionalities.
// Functionalities include the ability to stop and remove a docker container
// identified by an id specified in the task environment.
type TaskRunnerPlus struct {
	ContainerEnvKey string
	taskEnv         map[string]string
	l               *log.Logger
}

// NewTaskRunnerPlus creates a new task runner
func NewTaskRunnerPlus(logger *log.Logger, taskEnv map[string]string) *TaskRunnerPlus {
	return &TaskRunnerPlus{
		ContainerEnvKey: "CONTAINER_ID",
		taskEnv:         taskEnv,
		l:               logger,
	}
}

// Unofficial Feature: Deletes any docker image with a matching id as the `CONTAINER_ID` in the TaskEnv
func (r *TaskRunnerPlus) stopContainer() error {

	containerID := r.taskEnv[r.ContainerEnvKey]
	if containerID == "" {
		return fmt.Errorf("Container id is required")
	}

	r.l.Printf("[DEBUG] driver.raw_exec: Attempting to stop associated container (if any)")

	err := tools.DeleteContainer(containerID, false, false, false)
	if err != nil {
		if err == tools.ErrContainerNotFound {
			r.l.Printf("[DEBUG] driver.raw_exec: No associated container found")
			return nil
		}
		return fmt.Errorf("failed to delete container attached to task")
	}

	r.l.Printf("[DEBUG] driver.raw_exec: Successfully stopped task container")
	return nil
}

// SendGRPCSignal sends an artifical signal to a GRPC server running
// on a container (or anywhere). If request does not return within the timeout
// period it returns with no error
func (r *TaskRunnerPlus) SendGRPCSignal(timeout time.Duration) error {
	return nil // not implemented
}
