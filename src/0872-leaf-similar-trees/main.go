package main
import "fmt"
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i1, i2, o := NewTreeNode("3,5,1,6,2,9,8,null,null,7,4"), NewTreeNode("3,5,1,6,7,4,2,null,null,null,null,null,null,9,8"), true
  fmt.Printf("Input:  %s %s\n", i1.ToStr(), i2.ToStr())
  fmt.Printf("Output: %t\n", leafSimilar(i1, i2))
  fmt.Printf("Expect: %t\n", o)
}

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

