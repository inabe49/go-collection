package seq

func Exists[A any](slice []A, p func(A) bool) bool {
	for _, item := range slice {
		if p(item) {
			return true
		}
	}

	return false
}
