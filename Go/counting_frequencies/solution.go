package countingfrequencies

// You are given  queries. Each query is of the form two integers described below:
// 1 - x : Insert x in your data structure.
// 2 - y : Delete one occurence of y from your data structure, if present.
// 3 - z : Check if any integer is present whose frequency is exactly . If yes, print 1 else 0.

// The queries are given in the form of a 2-D array  of size  where  contains the operation, and  contains the data element.

// Example

// The results of each operation are:

// Operation   Array   Output
// (1,1)       [1]
// (2,2)       [1]
// (3,2)                   0
// (1,1)       [1,1]
// (1,1)       [1,1,1]
// (2,1)       [1,1]
// (3,2)                   1
// Return an array with the output: .

// Complete the freqQuery function below.
func FreqQuery(queries [][]int32) []int32 {
	res := []int32{}
	cache := map[int32]int32{}
	reverseCache := map[int32]int32{}

	for _, query := range queries {
		switch query[0] {
		case 1:
			prev := cache[query[1]]
			current := prev + 1
			cache[query[1]] = current
			reverseCache[current]++
			if prev != 0 {
				reverseCache[prev]--
				if reverseCache[prev] <= 0 {
					delete(reverseCache, prev)
				}
			}
		case 2:
			prev, ok := cache[query[1]]
			if ok {
				current := prev - 1
				if current == 0 {
					delete(cache, query[1])
				} else {
					cache[query[1]] = current
				}
				reverseCache[current]++
				if prev != 0 {
					reverseCache[prev]--
					if reverseCache[prev] <= 0 {
						delete(reverseCache, prev)
					}
				}
			}
		case 3:
			_, ok := reverseCache[query[1]]
			var queryRes int32
			if ok {
				queryRes = 1
			}
			res = append(res, queryRes)

		}
	}
	return res
}
