package twosum

func TwoSum(nums []int, target int) []int {
	arr := make([]int, 2)
	l := len(nums)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				arr[0], arr[1] = i, j
				return arr
			}
		}
	}

	return arr
}
