package leetcode

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}

	// 在最前頭增加一位數，而後面的位數皆為 0
	newDigits := make([]int, len(digits)+1)
	newDigits[0] = 1

	return newDigits
}
