package mask

// mask маскирует ссылку
func Mask(n int, um []byte, ms byte) string {
	for i := n; i < len(um); i++ {
		if um[i] == ' ' {
			break
		}
		if um[i] != ' ' {
			um[i] = ms
		}
	}
	return string(um)
}
