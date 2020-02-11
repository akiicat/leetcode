package main
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

// leetcode 572.
// T: O(n)
// M: O(h)
// -- start --

func isSameTree(p *TreeNode, q *TreeNode) bool {
  if p == nil && q == nil {
    return true
  }

  if p == nil && q != nil || p != nil && q == nil {
    return false
  }

  return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSameTreeConcise(s *TreeNode, t *TreeNode) bool {
  if s == nil && t == nil {
    return true
  }

  if s == nil || t == nil || s.Val != t.Val {
    return false
  }

  return isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}

// -- end --

