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

func invertTree(root *TreeNode) *TreeNode {
  if root == nil {
    return root
  }

  invertTree(root.Left)
  invertTree(root.Right)

  root.Left, root.Right = root.Right, root.Left

  return root
}

// -- end --

