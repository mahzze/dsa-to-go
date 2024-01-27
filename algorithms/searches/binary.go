package searches

func BinarySearch(arr []int, target int) bool {
	high := len(arr)
	low := 0

	for low < high {
		middle := (high-low)/2 + low
		print(middle, "\n")
		if arr[middle] == target {
			return true
		} else if arr[middle] > target {
			high = middle
		} else {
			low = middle + 1
		}
	}
	return false
}
