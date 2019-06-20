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
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
    _,maxDepth := search(root)
    return maxDepth
}

func search(root *TreeNode)(int,*TreeNode){
    if root == nil{
        return 0,nil
    }
    ldep,lnode := search(root.Left)
    rdep,rnode := search(root.Right)
    max := Max(ldep, rdep) + 1
    if ldep == rdep{
        return max,root
    }
    if ldep > rdep{
        return max,lnode
    }else{
        return max,rnode
    }
}

func Max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
