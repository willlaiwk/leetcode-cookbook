package leetcode

// 方法一: O(n^2)
func mergeSortedArraySol1(nums1 []int, m int, nums2 []int, n int) {
	j := 0
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[j]
		j++
	}

	for i := 1; i < len(nums1); i++ {
		for j := len(nums1) - 1; j >= i; j-- {
			if nums1[j] < nums1[j-1] {
				tmp := nums1[j]
				nums1[j] = nums1[j-1]
				nums1[j-1] = tmp
			}
		}
	}
}

// 方法二: O(n)
func mergeSortedArraySol2(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			k--
			i--
		} else {
			nums1[k] = nums2[j]
			k--
			j--
		}
	}
}
