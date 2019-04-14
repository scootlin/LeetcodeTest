/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    var total int
    search(root,&total,"")
    return total
}
func search(root *TreeNode,total *int,position string){
    if root!= nil{
        if root.Left == nil && root.Right == nil{
            if position == "L"{
                *total += root.Val
            }
        }
        if root.Left!=nil{
            search(root.Left,total,"L")
        }
        if root.Right!=nil{
            search(root.Right,total,"R")
        }        
    }
}
