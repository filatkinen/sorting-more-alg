package sort

func QuickSort(array []int) {
	quickSort(array, 0, len(array)-1)
}

func quickSort(array []int, start, end int) {
	left, right := start, end
	pivot := array[(start+end)/2]
	for left <= right {
		for array[left] < pivot {
			left++
		}
		for array[right] > pivot {
			right--
		}
		if left <= right {
			array[left], array[right] = array[right], array[left]
			left++
			right--
		}
	}
	if right > start {
		quickSort(array, start, right)
	}
	if left < end {
		quickSort(array, left, end)
	}
}

func QuickSortAppend(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	pivot := slice[0]
	var less, greater []int
	for _, num := range slice[1:] {
		if pivot > num {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	result := append(QuickSortAppend(less), pivot)
	result = append(result, QuickSortAppend(greater)...)
	return result
}
