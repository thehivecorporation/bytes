package bytes

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Float(t *testing.T) {
	t.Run("Test bytes with array reduction", func(t *testing.T) {
		var n float64 = math.MaxFloat64
		byt := Float64ToBytes(n)
		assert.Equal(t, n, Float64FromBytes(byt))

		var n32 float32 = math.MaxFloat32
		byt = Float32ToBytes(n32)
		assert.Equal(t, n32, Float32FromBytes(byt))

		n = math.SmallestNonzeroFloat64
		byt = Float64ToBytes(n)
		assert.Equal(t, n, Float64FromBytes(byt))

		n32 = math.SmallestNonzeroFloat32
		byt = Float32ToBytes(n32)
		assert.Equal(t, n32, Float32FromBytes(byt))
	})

	t.Run("Float64 to bytes and to Float64", func(t *testing.T) {
		rand.Seed(time.Now().UTC().UnixNano())

		input := getRandomSignedFloat64()

		bytes := Float64ToBytes(input)

		if input != Float64FromBytes(bytes) {
			t.Errorf(fmt.Sprintf("Failed Conversion for float64: %v", input))
		}
	})
}
