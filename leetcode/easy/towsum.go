package easy


func TowSum(target int, nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]+nums[i] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0,0}
}

