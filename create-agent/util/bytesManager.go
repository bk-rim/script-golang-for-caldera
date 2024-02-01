package util

func BytesContains(data []byte, sub []byte) bool {
	return len(data) >= len(sub) && bytesIndex(data, sub) != -1
}

func bytesIndex(data []byte, sub []byte) int {
	for i := range data {
		if len(data)-i < len(sub) {
			return -1
		}
		if bytesEqual(data[i:i+len(sub)], sub) {
			return i
		}
	}
	return -1
}

func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
