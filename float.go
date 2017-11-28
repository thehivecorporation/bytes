package bytes

import (
	"encoding/binary"
	"math"
)

func Float64ToBytes(f float64) (b []byte) {
	b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, math.Float64bits(f))
	return b
}

func Float64FromBytes(b []byte) (f float64) {
	if len(b) != 8 {
		return 0
	}

	bits := binary.BigEndian.Uint64(b)
	f = math.Float64frombits(bits)
	return
}

func Float32ToBytes(f float32) (b []byte) {
	b = make([]byte, 4)
	binary.BigEndian.PutUint32(b, math.Float32bits(f))
	return b
}

func Float32FromBytes(b []byte) (f float32) {
	if len(b) != 4 {
		return 0
	}

	bits := binary.BigEndian.Uint32(b)
	f = math.Float32frombits(bits)
	return
}
