package operator

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGolang_LessThanEqual(t *testing.T) {
  assert.False(t, 10 < 10)
}
