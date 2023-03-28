package main


func merge(nums1 []int, m int, nums2 []int, n int) {
	insertIdx := len(nums1) - 1
	first := m - 1
	second := n - 1

	for insertIdx >= 0 {
		if first >= 0 && second >= 0 {
			if nums1[first] > nums2[second] {
				nums1[insertIdx] = nums1[first]
				first--
			} else {
				nums1[insertIdx] = nums2[second]
				second--
			}
		} else if first >= 0 {
			nums1[insertIdx] = nums1[first]
			first--
		} else if second >= 0 {
			nums1[insertIdx] = nums2[second]
			second--
		}
		insertIdx--
	}
}

func main() {
	one := []int{1, 2, 3, 0, 0, 0}
	two := []int{2, 5, 6}
	merge(one, 3, two, 3)
}
