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

func findTilt(root *TreeNode) int {
  sum := 0
  SumTree(root, &sum)
  return sum
}

func SumTree(root *TreeNode, sum *int) int {
  if root == nil {
    return 0
  }

  l := SumTree(root.Left, sum)
  r := SumTree(root.Right, sum)

  *sum += Abs(l, r)

  return l + r + root.Val
}

func Abs(a, b int) int {
  if a > b {
    return a - b
  }
  return b - a
}

// -- end --

