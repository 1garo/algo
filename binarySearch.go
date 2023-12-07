package main

func BinarySearch(items []int, find int) int {
	low := 0
	high := len(items) - 1

	for low < high {
		mid := low + (high - low) / 2
		guess := items[mid]

		if guess == find {
			return mid
		} else if find > guess {
			low = mid + 1
		} else if find < guess {
			high = mid
		}
	}

	return -1
}
