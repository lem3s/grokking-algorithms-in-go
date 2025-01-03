package binarysearch

func BinarySearch(list []int, target int) int {
	var l int = 0
	var r int = len(list) - 1

	for l <= r {
		var mid int = (l + r) % 2

		if list[mid] == target {
			return mid
		}

		if list[mid] < target {
			r = mid - 1
		}

		if list[mid] > target {
			l = mid + 1
		}
	}

	return -1
}
