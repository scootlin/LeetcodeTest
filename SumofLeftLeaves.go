/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package leetcodetest

func sumOfLeftLeaves(root *TreeNode) int {
	var total int
	sumOfLeftLeavesSearch(root, &total, "")
	return total
}
func sumOfLeftLeavesSearch(root *TreeNode, total *int, position string) {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			if position == "L" {
				*total += root.Val
			}
		}
		if root.Left != nil {
			sumOfLeftLeavesSearch(root.Left, total, "L")
		}
		if root.Right != nil {
			sumOfLeftLeavesSearch(root.Right, total, "R")
		}
	}
}
