package seq

func Map[A, B any](source []A, p func(A) B) []B {
	var result = make([]B, len(source))

	for i, item := range source {
		result[i] = p(item)
	}

	return result
}

func FlatMap[A, B any](source []A, p func(A) []B) []B {
	var result = make([]B, 0)

	for _, item := range source {
		result = append(result, p(item)...)
	}

	return result
}

func Drop[A any](slice []A, n int) []A {
	length := len(slice) - n

	if length < 1 {
		return []A{}
	}

	var result = make([]A, length)

	for i, item := range slice {
		if i >= n {
			result[i-n] = item
		}
	}

	return result
}

func DropRight[A any](slice []A, n int) []A {
	length := len(slice) - n

	if length < 1 {
		return []A{}
	}

	var result = make([]A, length)

	for i, item := range slice {
		if i < length {
			result[i] = item
		}
	}

	return result
}

func DropWhile[A any](slice []A, p func(A) bool) []A {
	length := len(slice)
	start := false
	n := 0
	temp := make([]A, length)

	for i, item := range slice {
		if start {
			temp[i-n] = item
		} else {
			if p(item) {
				n += 1
			} else {
				temp[i-n] = item
				start = true
			}
		}
	}

	result := make([]A, length-n)
	copy(result, temp)

	return result
}

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
