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
  s, sum := 0, 0
  R(root, s, &sum)
  return sum
}

func R(root *TreeNode, s int, sum *int) {
  if root == nil {
    return
  }

  s = s << 1 | root.Val
  if root.Left == nil && root.Right == nil {
    *sum += s
    return
  }

  R(root.Left, s, sum)
  R(root.Right, s, sum)
}

// -- end --

