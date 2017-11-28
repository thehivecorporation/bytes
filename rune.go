package bytes

func RuneFromBytes(b []byte) rune {
	if len(b) == 0 {
		return 0
	}

	return Int32FromBytes(b)
}

func RuneToBytes(r rune) []byte {
	return Int32ToBytes(int32(r))
}
