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

// T: O(n)
// M: O(h)
func isSymmetric(root *TreeNode) bool {
  return isSymmetricRecursive(root, root)
}

func isSymmetricRecursive(l *TreeNode, r *TreeNode) bool {
  if l == nil && r == nil {
    return true
  }

  if l == nil || r == nil {
    return false
  }

  return l.Val == r.Val && isSymmetricRecursive(l.Left, r.Right) && isSymmetricRecursive(l.Right, r.Left)
}

// T: O(n)
// M: O(n)
func isSymmetricIterative(root *TreeNode) bool {
  if root == nil {
    return true
  }

  stack, l, r := []*TreeNode{root, root}, root, root

  for len(stack) != 0 {
    stack, l, r = stack[:len(stack)-2], stack[len(stack)-2], stack[len(stack)-1]

    if l.Val != r.Val {
      return false
    }

    ll, rr := l.Left != nil, r.Right != nil
    lr, rl := l.Right != nil, r.Left != nil
    if ll && rr {
      stack = append(stack, l.Left, r.Right)
    }

    if lr && rl {
      stack = append(stack, l.Right, r.Left)
    }

    if ll != rr || lr != rl {
      return false
    }
  }
  return true
}

// -- end --

