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

func findSecondMinimumValue(root *TreeNode) int {
  m1, m2 := 1 << 31, 1 << 31

  DFS(root, &m1, &m2)

  if m2 == 1 << 31 {
    return -1
  }

  return m2
}

func DFS(root *TreeNode, m1, m2 *int) {
  if root == nil {
    return
  }

  DFS(root.Left, m1, m2)

  if root.Val < *m1 {
    *m1, *m2 = root.Val, *m1
  } else if root.Val != *m1 && root.Val < *m2 {
    *m2 = root.Val
  }

  DFS(root.Right, m1, m2)
}

// -- end --

