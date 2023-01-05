package anagramspairs

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func SherlockAndAnagrams(s string) int32 {
	// Write your code here
	var res int32
	substrCache := make(map[string]struct{})
	strLen1 := len(s)
	for left := 0; left < strLen1; left++ {
		for right := left + 1; right <= strLen1; right++ {
			substr := s[left:right]
			var revSubstr string
			if len(substr) == 1 {
				revSubstr = substr
			} else {
				revSubstr = Reverse(substr)
			}
			_, ok := substrCache[revSubstr]
			if ok {
				res++
			}
			substrCache[substr] = struct{}{}
		}
	}
	return res
}
