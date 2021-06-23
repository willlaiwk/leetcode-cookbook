/**
 * 暴力解: O(n^2)
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
 var twoSumSol1 = function(nums, target) {
    for (let i = 0; i < nums.length-1; i++) {
        for (let j = i+1; j < nums.length; j++) {
            if (nums[i]+nums[j] === target) {
                return [i, j]
            }
        }
    }

    return null
 }

/**
 * Map解: O(n)
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSumSol2 = function(nums, target) {
    const numMap = new Map()

    for (let i = 0; i < nums.length; i++) {
        if (numMap.has(target-nums[i])) {
            return [i, numMap.get(target-nums[i])]
        }
        numMap.set(nums[i], i)
    }

    return null
 }