package bytes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Rune(t *testing.T) {
	var r rune = 'a'
	byt := RuneToBytes(r)
	assert.Equal(t, r, RuneFromBytes(byt))
}
