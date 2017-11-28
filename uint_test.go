package bytes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"math"
)

func Test_Uint(t *testing.T) {
	t.Run("Negative numbers and big ops", func(t *testing.T) {
		var n uint = math.MaxUint32
		byt := UintToBytes(n)
		assert.Equal(t, n, uintFromBytes(byt))

		var n8 uint8 = math.MaxUint8
		byt = Uint8ToBytes(n8)
		assert.Equal(t, n8, uint8FromBytes(byt))

		var n16 uint16 = math.MaxUint16
		byt = Uint16ToBytes(n16)
		assert.Equal(t, n16, uint16FromBytes(byt))

		n16 = math.MaxUint8
		byt = Uint16ToBytes(n16)
		assert.Equal(t, n16, uint16FromBytes(byt))

		var n32 uint32 = math.MaxUint32
		byt = Uint32ToBytes(n32)
		assert.Equal(t, n32, Uint32FromBytes(byt))

		n32 = math.MaxUint16
		byt = Uint32ToBytes(n32)
		assert.Equal(t, n32, Uint32FromBytes(byt))

		var n64 uint64 = math.MaxUint64
		byt = Uint64ToBytes(n64)
		assert.Equal(t, n64, uint64FromBytes(byt))

		n64 = math.MaxUint32
		byt = Uint64ToBytes(n64)
		assert.Equal(t, n64, uint64FromBytes(byt))
	})
}
