package mockeable

import "sync"

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
