package bytes

import (
	"encoding/binary"
)

func Uint8ToBytes(u uint8) []byte {
	return []byte{u}
}

func Uint16ToBytes(u uint16) []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, u)

	if buf[0] != 0 {
		return buf
	}

	return buf[1:]
}

func Uint32ToBytes(u uint32) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, u)

	for i := 0; i < 3; i++ {
		if buf[i] != 0 {
			return buf[i:]
		}
	}

	return buf[3:]
}

func Uint64ToBytes(u uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, u)

	for i := 0; i < 7; i++ {
		if buf[i] != 0 {
			return buf[i:]
		}
	}

	return buf[7:]
}

func UintToBytes(u uint) []byte {
	return Uint32ToBytes(uint32(u))
}

func uintFromBytes(byt []byte) uint {
	if len(byt) == 0 {
		return 0
	}

	if len(byt) != 4 {
		prefix := make([]byte, 4-len(byt))
		return uint(binary.BigEndian.Uint32(append(prefix, byt...)))
	}

	return uint(binary.BigEndian.Uint32(byt))
}

func uint8FromBytes(b []byte) uint8 {
	if len(b) == 0 {
		return 0
	}

	return uint8(binary.BigEndian.Uint16(append([]byte{0}, b...)))
}

func uint16FromBytes(b []byte) uint16 {
	if len(b) == 0 {
		return 0
	}
	if len(b) != 2 {
		prefix := make([]byte, 2-len(b))
		return binary.BigEndian.Uint16(append(prefix, b...))
	}

	return binary.BigEndian.Uint16(b)
}

func Uint32FromBytes(b []byte) uint32 {
	if len(b) == 0 {
		return 0
	}

	if len(b) != 4 {
		prefix := make([]byte, 4-len(b))
		return binary.BigEndian.Uint32(append(prefix, b...))
	}

	return binary.BigEndian.Uint32(b)
}

func uint64FromBytes(b []byte) uint64 {
	if len(b) == 0 {
		return 0
	}
	if len(b) != 8 {
		prefix := make([]byte, 8-len(b))
		return binary.BigEndian.Uint64(append(prefix, b...))
	}

	return binary.BigEndian.Uint64(b)
}
