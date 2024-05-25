/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package leetcodetest

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	List := &[]int{}
	BinTreeSearch(root, List)
	return *List
}

func BinTreeSearch(node *TreeNode, List *[]int) {
	if node != nil {
		BinTreeSearch(node.Left, List)
		*List = append(*List, node.Val)
		BinTreeSearch(node.Right, List)
	}
}
