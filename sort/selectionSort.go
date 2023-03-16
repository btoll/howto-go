package sort

func swap(nums []int, j, k int) {
	temp := nums[j]
	nums[j] = nums[k]
	nums[k] = temp
}

func SelectionSort(nums []int) []int {
	i := 0
	for i < len(nums) {
		min_index := i
		for j := i + 1; j < len(nums); j++ {
			if nums[min_index] > nums[j] {
				min_index = j
			}
		}
		if i != min_index {
			swap(nums, i, min_index)
		}
		i += 1
	}
	return nums
}
