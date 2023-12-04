package main

func BubbleSort(items []int) []int {
	for i := 0; i < len(items); i++ {
		for j := 0; j < len(items) - 1 - i; j++ {
			if items[j] > items[j+1] {
				curr := items[j]
				items[j] = items[j+1]
				items[j+1] = curr
			}
		}
	}

	return items
}
