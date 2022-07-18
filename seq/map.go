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
