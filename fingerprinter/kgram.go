package fingerprinter

func KGram(k int, input string) (result []string) {
	result = make([]string, 0)
	if k < 1 {
		return result
	}
	l := len(input)

	for i := k; i <= l; i++ {
		result = append(result, input[i-k:i])
	}

	return result
}
