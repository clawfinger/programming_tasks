package validstring

func IsValid(s string) string {
	// cache := make(map[rune]int)
	// for _, letter := range s {
	// 	cache[letter]++
	// }
	// diffCache := make(map[int]int)

	// for _, count := range cache {
	// 	for letter, opposite := range cache {
	// 		diff := math.Abs(float64(count - opposite))
	// 		if diff != 0 {
	// 			diffCache[letter]++
	// 		}
	// 	}
	// }
	// if len(diffCache) > 1 {
	// 	return "NO"
	// }
	return "YES"
}
