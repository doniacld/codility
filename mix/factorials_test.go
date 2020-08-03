package mix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorials(t *testing.T) {
	res := Factorials(3)
	assert.Equal(t, "6, 2, 1", res)
}
