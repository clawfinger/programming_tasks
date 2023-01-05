package geometricprogressiontriplets

// Complete the countTriplets function below.
func CountTriplets(arr []int64, r int64) int64 {
	entryIndices := make(map[int64][]int)
	for index, entry := range arr {
		entryIndices[entry] = append(entryIndices[entry], index)
	}
	var res int64
	for _, entry := range arr {
		second := entry * r
		third := second * r
		secondEntries, secondOk := entryIndices[second]
		thirdEntries, thirdOk := entryIndices[third]
		if secondOk && thirdOk {
			secondLen := len(secondEntries)
			thirdLen := len(thirdEntries)
			if secondLen == 1 && thirdLen == 1 {
				res++
			} else if secondLen == 1 {
				res += int64(thirdLen)
			} else {
				res += int64(secondLen)
			}
		}

	}
	return res
}
