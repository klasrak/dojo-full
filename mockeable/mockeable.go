package mockeable

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Mockeable interface {
	GetFuncControls() []*CallsFuncControl
	CleanUp()
	Use()
}

func CleanUpAndAssertControls(t *testing.T, mock Mockeable) {
	defer mock.CleanUp()
	for _, control := range mock.GetFuncControls() {
		if !control.IgnoreCallsAssertion {
			assert.Equal(t, control.ExpectedCalls, control.funcCalls, fmt.Sprintf("expected calls for func %s does not match actual", control.funcName))
		}
	}
}

type CallsFuncControl struct {
	funcName             string
	funcCalls            int
	ExpectedCalls        int
	IgnoreCallsAssertion bool
	mu                   sync.Mutex
}

func (c *CallsFuncControl) SetFuncName(name string) {
	c.funcName = name
}

func (c *CallsFuncControl) IncreaseCallCount() {
	c.mu.Lock()
	c.funcCalls++
	c.mu.Unlock()
}
