package bytes

func BoolToBytes(t bool) []byte {
	if t {
		return []byte{1}
	}
	return []byte{0}
}

func BoolFromBytes(b []byte) (r bool) {
	if len(b) == 0 || len(b) > 1 {
		return
	}

	r = b[0] == 1

	return
}
