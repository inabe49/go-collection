package seq

func Filter[A any](slice []A, p func(A) bool) []A {
	length := len(slice)
	n := 0
	temp := make([]A, length)

	for _, item := range slice {
		if p(item) {
			temp[n] = item
			n += 1
		}
	}

	result := make([]A, n)
	copy(result, temp)

	return result
}
