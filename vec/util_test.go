package vec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	a := New(0, 0, 0)
	b := New(1, 1, 1)

	assert.Equal(t, b, Add(a, b))
	assert.Equal(t, New(1, 1, 1), b)
	assert.Equal(t, New(2, 2, 2), Add(b, b))
	assert.Equal(t, New(3, 3, 3), Add(b, b, b))
}
