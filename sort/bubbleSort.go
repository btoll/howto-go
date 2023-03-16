package sort

func BubbleSort(nums []int) []int {
	var temp int
	i := 0
	l := len(nums)
	for i < l {
		jlen := l - i
		// Every pass through the slice will sort a number
		// at the end, so the length of the inner loop (`jlen`)
		// can be decremented each time.
		for j := 1; j < jlen; j++ {
			if nums[j] < nums[j-1] {
				temp = nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = temp
			}
		}
		i += 1
	}
	return nums
}
