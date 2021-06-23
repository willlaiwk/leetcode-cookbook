# [Two Sum](https://leetcode.com/problems/two-sum/)
- Difficulty: `Easy`
- Tags: `Array`、`Hash Table`

## 題目
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

 
**Example 1:**

```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
```

**Example 2:**

```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

**Example 3:**

```
Input: nums = [3,3], target = 6
Output: [0,1]
```

**Constraints:**

- `2 <= nums.length <= 10^4`
- `-10^9 <= nums[i] <= 10^9`
- `-10^9 <= target <= 10^9`
- `Only one valid answer exists.`

## 解題思路
直覺想到的方法是使用兩層迴圈的暴力解，時間複雜度為 `O(n^2)`。

但我們可以用 Map (or Dictionary) 將時間複雜度降到 `O(n)`，

思路是 Loop `nums`：

1. 若 `map[target - nums[i]]` 不存在，則 `map[nums[i]] = i`
2. 若 `map[target - nums[i]]` 存在，則返回 `[i, map[target - nums[i]]]`
