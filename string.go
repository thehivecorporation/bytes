package bytes

func StringToBytes(s string) []byte {
	return []byte(s)
}

func StringFromBytes(b []byte) (s string) {
	if len(b) == 0 {
		return
	}

	s = string(b)

	return
}
