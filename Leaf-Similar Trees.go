/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    str1 := ""
    str2 := ""
    Travers(root1,&str1)
    Travers(root2,&str2)
    return str1 == str2
}
func Travers(node *TreeNode, str *string){
    if node == nil{
        return
    }
    if node.Left == nil && node.Right == nil{
        *str += strconv.Itoa(node.Val)
        return
    }
    Travers(node.Left,str)
    Travers(node.Right,str)
}
