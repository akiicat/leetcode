package main
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

// T: O(n)
// M: O(n)
// -- start --

func findTarget(root *TreeNode, k int) bool {
  m := make(map[int]bool)

  return r(root, k, m)
}

func r(root *TreeNode, k int, m map[int]bool) bool {
  if root == nil {
    return false
  }

  _, ok := m[root.Val]
  if ok {
    return true
  }

  m[k - root.Val] = true

  return r(root.Left, k, m) || r(root.Right, k, m)
}

// -- end --

