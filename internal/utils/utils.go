package utils

func DivideString(input_string string) string {
	size := 1
	partSize := len(input_string) / size
	parts := make([][]rune, size)

	for i := 0; i < size; i++ {
		end := i + partSize
		if end > len(input_string) {
			end = len(input_string)
		}
		parts[i] = []rune(input_string)[i:end]
	}

	partOne := parts[0]
	return string(partOne)
}
