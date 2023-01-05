package trailingmedian

func swap(arr []int32, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func partition(arr []int32, left int, right int) int {
	i := left
	j := right
	pivot := arr[right]
	for {
		for arr[i] < pivot {
			i++
		}

		for arr[j] >= pivot {
			if j == i {
				break
			}
			j--
		}
		if i >= j {
			break
		}
		swap(arr, i, j)
	}
	swap(arr, i, right)
	return i
}

func getMedianPos(slice []int32, left int, right int) int {
	pivotPos := left + (right-left)/2
	newPivotPos := partition(slice, left, right)
	i := left
	j := right
	for {
		if pivotPos == newPivotPos {
			return pivotPos
		} else if newPivotPos < pivotPos {
			i = newPivotPos + 1
			newPivotPos = partition(slice, i, j)
		} else {
			j = newPivotPos - 1
			newPivotPos = partition(slice, i, j)
		}

	}
}

func GetMedianWithSelect(slice []int32) float32 {
	median := getMedianPos(slice, 0, len(slice)-1)
	if len(slice)%2 == 0 {
		a := slice[median]
		b := slice[median+1]
		return (float32(a) + float32(b)) / 2.0
	} else {
		return float32(slice[median])
	}
}

func ActivityNotifications(expenditure []int32, d int32) int32 {
	var res int32
	for i := int(d); i < len(expenditure); i++ {
		median := GetMedianWithSelect(expenditure[(i - int(d)):i])
		if float32(expenditure[i]) >= median*2 {
			res++
		}
	}
	return res
}

func ActivityNotificationsWithArray(expenditure []int32, d int32) (int32, []float64) {
	res := make([]float64, 0)
	for i := int(d); i <= len(expenditure); i++ {
		window := make([]int32, 0)

		median := GetMedianWithSelect(append(window, expenditure[(i-int(d)):i]...))
		res = append(res, float64(median))
	}
	return 0, res
}
