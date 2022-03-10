package main
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

// leetcode 501.
// T: O(n)
// M: O(h)
// -- start --

func getMinimumDifference(root *TreeNode) int {
  prev := -1
  return DFS(root, &prev)
}

func DFS(root *TreeNode, prev *int) int {
  if root == nil {
    return 1<<31
  }

  min := DFS(root.Left, prev)

  if *prev != -1 {
    min = Min(root.Val - *prev, min)
  }
  *prev = root.Val

  min = Min(DFS(root.Right, prev), min)

  return min
}

func Min(a, b int) int {
  if a > b {
    return b
  }
  return a
}

// -- end --

