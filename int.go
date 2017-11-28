package bytes

import (
	"encoding/binary"
)

func Int8FromBytes(b []byte) int8 {
	if len(b) == 0 {
		return 0
	}

	return int8(uint8FromBytes(b))
}

func Int16FromBytes(b []byte) int16 {
	if len(b) == 0 || len(b) != 2 {
		return 0
	}

	return int16(uint16FromBytes(b))
}

func Int32FromBytes(b []byte) int32 {
	if len(b) == 0 {
		return 0
	}

	return int32(Uint32FromBytes(b))
}

func Int64FromBytes(b []byte) int64 {
	if len(b) == 0 || len(b) != 8 {
		return 0
	}

	return int64(uint64FromBytes(b))
}

func IntFromBytes(b []byte) int {
	if len(b) == 0 {
		return 0
	}

	if len(b) != 4 {
		prefix := make([]byte, 4-len(b))
		return int(binary.BigEndian.Uint32(append(prefix, b...)))
	}

	return int(binary.BigEndian.Uint32(b))
}

func IntToBytes(i int) []byte {
	return UintToBytes(uint(i))
}

func Int8ToBytes(i int8) (byt []byte) {
	return []byte{uint8(i)}
}

func Int16ToBytes(u int16) []byte {
	return Uint16ToBytes(uint16(u))
}

func Int32ToBytes(i int32) (byt []byte) {
	return Uint32ToBytes(uint32(i))
}

func Int64ToBytes(i int64) []byte {
	return Uint64ToBytes(uint64(i))
}
