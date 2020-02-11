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
func isBalanced(root *TreeNode) bool {
  _, r := R(root)
  return r
}

func R(root *TreeNode) (int, bool) {
  if root == nil {
    return 0, true
  }

  l, o1 := R(root.Left)
  r, o2 := R(root.Right)

  return Max(l, r) + 1, o1 && o2 && Abs(l - r) <= 1
}


// T: O(n ** 2)
// M: O(h)
func isBalancedRecursive(root *TreeNode) bool {
  if root == nil {
    return true
  }

  if Abs(maxDepth(root.Left) - maxDepth(root.Right)) > 1 {
    return false
  }

  return isBalanced(root.Left) && isBalanced(root.Right)
}

// 104-maximum-depth-of-binary-tree
// T: O(n)
// M: O(h)
func maxDepth(root *TreeNode) int {
  if root == nil {
    return 0
  }

  return Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func Abs(a int) int {
  if a < 0 {
    return -a
  }
  return a
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}


// -- end --

