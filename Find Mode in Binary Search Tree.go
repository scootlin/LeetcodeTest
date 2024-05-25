/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package leetcodetest

var maxArr []int
var max int = 0
var count int = 0
var BinTreeNodeList map[int]int

func findMode(root *TreeNode) []int {
	BinTreeNodeList = make(map[int]int)
	searchNode(root)
	for i, v := range BinTreeNodeList {
		if max < v {
			maxArr = maxArr[:0]
			maxArr = append(maxArr, i)
			max = v
		} else if max == v {
			maxArr = append(maxArr, i)
		}
	}
	return maxArr
}

func searchNode(node *TreeNode) {
	if node != nil {
		BinTreeNodeList[node.Val]++
		searchNode(node.Left)
		searchNode(node.Right)
	}
}
