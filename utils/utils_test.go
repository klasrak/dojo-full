package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToJSON(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		b := ToJSON(map[string]string{"foo": "bar"})
		assert.Equal(t, `{"foo":"bar"}`, string(b))
	})

	t.Run("Panic", func(t *testing.T) {
		assert.Panics(t, func() {
			ToJSON(make(chan int, 1))
		})
	})
}
