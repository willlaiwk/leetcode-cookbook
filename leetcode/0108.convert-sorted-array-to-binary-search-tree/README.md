# [108. Convert Sorted Array to Binary Search Tree](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/)
- Difficulty: `Easy`
- Tags: `Array`, `Divide and Conquer`, `Tree`, `Binary Search Tree`, `Binary Tree`

## 題目
Given an integer array nums where the elements are sorted in **ascending order**, convert it to a **height-balanced** binary search tree.

A **height-balanced** binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.

**Example 1:**
![](https://assets.leetcode.com/uploads/2021/02/18/btree1.jpg)
```
Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: [0,-10,5,null,-3,null,9] is also accepted:
```
![](https://assets.leetcode.com/uploads/2021/02/18/btree2.jpg)


**Example 2:**
![](https://assets.leetcode.com/uploads/2021/02/18/btree.jpg)
```
Input: nums = [1,3]
Output: [3,1]
Explanation: [1,3] and [3,1] are both a height-balanced BSTs.
```

**Constraints:**
- `1 <= nums.length <= 10^4`
- `-10^4 <= nums[i] <= 10^4`
- `nums` is sorted in a **strictly increasing** order.

## 解題思路
本題是將傳進的陣列轉換成 [**Binary Search Tree**](https://www.geeksforgeeks.org/binary-search-tree-data-structure/)。

實作過程與 [Binary Search](https://www.pgwill.tech/post/2021-06-27-binary-search) 一樣:
1. 先找到陣列的中間值，此中間值即為 Tree 的 top 節點，並將陣列從中間值切開分成左右兩個陣列。
2. 找到左右兩個陣列的中間值，這兩個中間值分別是上一個節點的左右子節點。接著分別將左右兩個陣列沿著中間值再切出左右陣列。
3. 重複以上動作，直到陣列為空。

而程式方面可透過遞迴的方式來實現。
