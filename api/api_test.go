package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test Run()
func TestRun(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		api := New()
		var err error

		running := make(chan struct{})
		done := make(chan struct{})

		go func() {
			close(running)
			err = api.Run()
			defer close(done)
		}()

		<-running
		assert.NotNil(t, api.Server)
		assert.NotNil(t, api.Server.Handler)
		assert.NotNil(t, api.Server.Addr)
		assert.Nil(t, err)
		// assert.NoError(t, api.Server.Shutdown(context.Background()))
		assert.NoError(t, api.Server.Close())
		<-done
	})
}
