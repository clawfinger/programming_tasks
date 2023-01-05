package commonsubstr

func TwoStrings(s1 string, s2 string) string {
	// Write your code here
	substrCache := make(map[string]struct{})
	strLen1 := len(s1)
	for left := 0; left < strLen1; left++ {
		for right := left + 1; right < strLen1; right++ {
			substrCache[s1[left:right]] = struct{}{}
		}
	}
	strLen2 := len(s2)
	for left := 0; left < strLen2; left++ {
		for right := left + 1; right < strLen2; right++ {
			_, ok := substrCache[s2[left:right]]
			if ok {
				return "YES"
			}
		}
	}
	return "NO"
}
