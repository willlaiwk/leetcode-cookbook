package leetcode

// 暴力解: O(n^2)
func twoSumSol1(nums []int, target int) []int {
	var result = make([]int, 2)

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result[0] = i
				result[1] = j
			}
		}
	}

	return result
}

// Map解: O(n)
func twoSumSol2(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if v, ok := numMap[target-nums[i]]; ok {
			return []int{i, v}
		}

		numMap[nums[i]] = i
	}

	return nil
}
