package k8s

import (
	"context"
	"fmt"
	"testing"

	podPlugin "github.com/flyteorg/flyteplugins/go/tasks/plugins/k8s/pod"

	"github.com/stretchr/testify/assert"
)

func TestSubTaskExecutionContext(t *testing.T) {
	ctx := context.Background()

	tCtx := getMockTaskExecutionContext(ctx, 0)
	taskTemplate, err := tCtx.TaskReader().Read(ctx)
	assert.Nil(t, err)

	executionIndex := 0
	originalIndex := 5
	retryAttempt := uint64(1)

	stCtx := newSubTaskExecutionContext(tCtx, taskTemplate, executionIndex, originalIndex, retryAttempt)

	assert.Equal(t, stCtx.TaskExecutionMetadata().GetTaskExecutionID().GetGeneratedName(), fmt.Sprintf("notfound-%d-%d", executionIndex, retryAttempt))

	subtaskTemplate, err := stCtx.TaskReader().Read(ctx)
	assert.Nil(t, err)
	assert.Equal(t, int32(2), subtaskTemplate.TaskTypeVersion)
	assert.Equal(t, podPlugin.ContainerTaskType, subtaskTemplate.Type)
}
