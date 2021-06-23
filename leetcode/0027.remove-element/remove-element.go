package leetcode

// 雙指針解(Two Pointers): O(n)
func removeElement(nums []int, val int) int {
	i := 0

	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}
