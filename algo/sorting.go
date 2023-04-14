package algo

/*
 * Merge Sort Algorith
 */
func MergeSort(list []int64) []int64 {
	if len(list) == 1 {
		return list
	}

	left := MergeSort(list[:len(list)/2])
	right := MergeSort(list[(len(list) / 2):])
	return merge(left, right)
}

func merge(list1 []int64, list2 []int64) []int64 {
	var res []int64
	var li, ri int
	for li < len(list1) && ri < len(list2) {
		if list1[li] < list2[ri] {
			res = append(res, list1[li])
			li++
		} else {
			res = append(res, list2[ri])
			ri++
		}
	}

	if li < len(list1) {
		res = append(res, list1[li:]...)
	} else {
		res = append(res, list2[ri:]...)
	}
	return res
}

/*
* Sorting based on callback
 */
 
