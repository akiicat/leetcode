package main
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

// leetcode 100.
// T: O(m*n)
// M: O(h)
// -- start --

func isSubtree(s *TreeNode, t *TreeNode) bool {
  if t == nil || s == nil {
    return false
  }

  l := isSubtree(s.Left, t)
  r := isSubtree(s.Right, t)

  return l || r || isSameTree(s, t)
}

// 0100-same-tree
func isSameTree(s *TreeNode, t *TreeNode) bool {
  if s == nil && t == nil {
    return true
  }

  if s == nil || t == nil || s.Val != t.Val {
    return false
  }

  return isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}

// -- end --

