package common

func Unpack7Bit(data []byte) []byte {
	var result []byte
	for i := 0; i < len(data); i += 8 {
		if i+8 > len(data) {
			break
		}
		msb := data[i]
		for j := 0; j < 7; j++ {
			b := data[i+1+j]
			if msb&(1<<j) != 0 {
				b |= 0x80
			}
			result = append(result, b)
		}
	}
	return result
}
