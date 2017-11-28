package bytes

func ToBytes(i interface{}) []byte {
	switch t := i.(type) {
	case float64:
		return Float64ToBytes(t)

	case string:
		return StringToBytes(t)

	case bool:
		return BoolToBytes(t)

	case uint:
		return UintToBytes(t)
	case uint8:
		return Uint8ToBytes(t)
	case uint16:
		return Uint16ToBytes(t)
	case uint32:
		return Uint32ToBytes(t)
	case uint64:
		return Uint64ToBytes(t)

	case int:
		return IntToBytes(t)
	case int8:
		return Int8ToBytes(t)
	case int16:
		return Int16ToBytes(t)
	case int32:
		return Int32ToBytes(t)
	case int64:
		return Int64ToBytes(t)
	}

	return nil
}

func Equal(b1, b2 []byte) (r bool) {
	if len(b1) != len(b2) {
		return
	}

	for i, v2 := range b2 {
		if v2 != b1[i] {
			return
		}
	}

	return true
}
