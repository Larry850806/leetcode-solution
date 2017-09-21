// https://leetcode.com/problems/trim-a-binary-search-tree/description/

package main

// TreeNode is definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// root, Left, Right can be nil
func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < L {
		return trimBST(root.Right, L, R)
	}

	if root.Val > R {
		return trimBST(root.Left, L, R)
	}

	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return root
}
