package main
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

// T: O(n + m) n, m are the length of the given trees.
// M: O(n + m)
// -- start --

// DFS
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
  value1 := Leafs(root1)
  value2 := Leafs(root2)

  if len(value1) != len(value2) {
    return false
  }

  for i := 0; i < len(value1); i++ {
    if value1[i] != value2[i] {
      return false
    }
  }

  return true
}

func Leafs(root *TreeNode) []int {
  s, r := []*TreeNode{root}, root  // stack

  value := []int{}

  for len(s) != 0 {
    r, s = s[len(s)-1], s[:len(s)-1]

    if r.Right != nil {
      s = append(s, r.Right)
    }

    if r.Left != nil {
      s = append(s, r.Left)
    }

    if r.Left == nil && r.Right == nil {
      value = append(value, r.Val)
    }
  }

  return value
}

// -- end --

