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

func averageOfLevels(root *TreeNode) []float64 {
  rtn := []float64{}
  s := []*TreeNode{root}

  for len(s) != 0 {
    n, sum := len(s), 0
    for i := 0; i < n; i++ {
      sum += s[i].Val
      if s[i].Left != nil {
        s = append(s, s[i].Left)
      }
      if s[i].Right != nil {
        s = append(s, s[i].Right)
      }
    }
    rtn = append(rtn, float64(sum) / float64(n))
    s = s[n:]
  }

  return rtn
}

// -- end --

