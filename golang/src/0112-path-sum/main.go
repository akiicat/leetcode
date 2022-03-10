package main
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

// T: O(n)
// M: O(h)
// -- start --

// Recursive Depth First Search
func hasPathSum(root *TreeNode, sum int) bool {
  if root == nil {
    return false
  }

  // Leaf
  if root.Left == nil && root.Right == nil {
    return sum == root.Val
  }

  return hasPathSum(root.Left, sum - root.Val) || hasPathSum(root.Right, sum - root.Val)
}

// -- end --

