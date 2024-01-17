package algo

func PrintArrayInRev(list []int) []int {
	left := 0
	right := len(list) - 1

	for left < right {
		list[left], list[right] = list[right], list[left]
		left++
		right--
	}

	return list
}

func ModifiedBinarySearch(nums []int, target int) int {
	if len(nums) == 1 && target != nums[0] {
		return -1
	}

	mid := len(nums) / 2

	if nums[mid] == target {
		return mid
	}

	left := false

	if nums[mid-1] == target {
		return mid - 1
	} else if nums[mid-1] > nums[mid] {

		if nums[0] == target {
			return 0
		}

		if target > nums[0] {
			left = true
		}
	} else {
		n := len(nums)

		if nums[n-1] == target {
			return n - 1
		}

		if target > nums[n-1] {
			left = true
		} else if target < nums[n-1] && target > nums[mid] {
			left = false
		} else {
			left = true
		}
	}

	if left {
		//return searchA(nums[:mid], target)
	} else {
		//return searchA(nums[mid:], target)
	}
	return 0
}
