package braillegraph

func encodeUnary(i uint8) uint64 {
	if i > 64 {
		return 0
	}
	result := uint64(0)
	for idx := 0; uint8(idx) < i; idx++ {
		result = (uint64(1) << uint(idx)) | result
	}
	return result
}

func toBoolArray(v, len uint) []bool {
	data := make([]bool, len)
	for i := uint(0); i < len; i++ {
		value := (v >> i) & 0x1
		data[i] = value == 1
	}
	return data
}

func isBetweenZeroAndOne(value float64) bool {
	return isInRange(value, 0, 1)
}

func isInRange(value, lower, upper float64) bool {
	return value >= lower && value <= upper
}
