package mask

func Search(p, k []byte) int {
	var n, j int
	for i := n; i < len(p); i++ {
		if p[i] != k[j] {
			j = 0
		}
		if p[i] == k[j] {
			j++
			if j == len(k) { // сравнился ли ключ полностью
				n = i + 1 // индекс, с которого начинается ссылка
				j = 0
			}
		}
	}
	return n
}
