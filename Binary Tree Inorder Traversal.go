/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
    List := &[]int{}
    search(root,List)
    return *List
}

func search(node *TreeNode, List *[]int){
    if node != nil{
        search(node.Left,List)
        *List = append(*List,node.Val)
        search(node.Right,List)
    }
}
