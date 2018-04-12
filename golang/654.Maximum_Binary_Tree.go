// https://leetcode.com/problems/maximum-binary-tree/description/

package main

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	max := nums[0]
	maxIdx := 0
	for i, n := range nums {
		if n > max {
			max = n
			maxIdx = i
		}
	}
	leftNums := nums[:maxIdx]
	rightNums := nums[maxIdx+1:]

	return &TreeNode{
		Val:   max,
		Left:  constructMaximumBinaryTree(leftNums),
		Right: constructMaximumBinaryTree(rightNums),
	}
}
