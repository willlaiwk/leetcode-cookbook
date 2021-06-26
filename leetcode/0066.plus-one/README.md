# [66. Plus One](https://leetcode.com/problems/plus-one/)
- Difficulty: `Easy`
- Tags: `Array`, `Math`

## 題目
Given a non-empty array of decimal digits representing a non-negative integer, increment one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contains a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

 

**Example 1:**
```
Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
```

**Example 2:**
```
Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
```

**Example 3:**
```
Input: digits = [0]
Output: [1]
```

**Constraints:**
- `1 <= digits.length <= 100`
- `0 <= digits[i] <= 9`


## 解題思路
此題用 Array 來表示一數，並做加 1 的操作。

從陣列後至前跑迴圈，並注意進位問題即可，時間複雜度為 `O(n)`。
