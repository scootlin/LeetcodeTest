/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// var maxDepth int = 0
// var level int = 0
// var recmap map[int]TreeNode
package leetcodetest

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	_, maxDepth := SumallestSubtreeSearch(root)
	return maxDepth
}

func SumallestSubtreeSearch(root *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}
	ldep, lnode := SumallestSubtreeSearch(root.Left)
	rdep, rnode := SumallestSubtreeSearch(root.Right)
	max := Max(ldep, rdep) + 1
	if ldep == rdep {
		return max, root
	}
	if ldep > rdep {
		return max, lnode
	} else {
		return max, rnode
	}
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
