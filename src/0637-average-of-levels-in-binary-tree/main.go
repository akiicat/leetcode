package main
import "fmt"
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i, o := NewTreeNode("3,9,20,15,7"), []float64{3.00000,14.50000,11.00000}
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %v\n", averageOfLevels(i))
  fmt.Printf("Expect: %v\n", o)
}

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

