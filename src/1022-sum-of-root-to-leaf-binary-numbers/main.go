package main
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

// T: O(n)
// M: O(1)
// -- start --

func sumRootToLeaf(root *TreeNode) int {
  return R(root, 0)
}

func R(root *TreeNode, s int) int {
  if root == nil {
    return
  }

  s = s << 1 | root.Val
  if root.Left == nil && root.Right == nil {
    return s
  }

  return R(root.Left, s, sum) + R(root.Right, s, sum)
}

// -- end --

