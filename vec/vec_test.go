package vec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVecDev(t *testing.T) {
	a := New(200, 0, 0)
	assert.Equal(t, New(100, 0, 0), a.Dev(2))

	o := New(0, 0, 0)
	h := New(400.0, 0, 0)
	v := New(0, 200.0, 0)
	llc := Sub(o, h.Dev(2), v.Dev(2), New(0, 0, 1.0))
	assert.Equal(t, New(-200, -100, -1), llc)
}
