package minimumswaps

// Complete the minimumSwaps function below.
//	array := []int32{4, 3, 1, 2}

func MinimumSwaps(arr []int32) int32 {
	value2index := make(map[int32]int)
	for index, value := range arr {
		if value != int32(index+1) {
			value2index[value] = index + 1
		}
	}
	var res int32
	l := len(value2index)
	for l > 0 {
		for value, index := range value2index {
			_ = index
			correctValue, ok := value2index[value]
			if ok {
				tmp := value2index[value]
				value2index[value] = value2index[int32(correctValue)]
				value2index[int32(correctValue)] = tmp
				delete(value2index, int32(correctValue))
				if value2index[value] == int(value) {
					delete(value2index, int32(value))
				}
				res++
				l = len(value2index)
				break
			}
		}
	}

	return res
}
