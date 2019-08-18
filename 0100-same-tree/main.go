package main
import "fmt"
import . "tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i1, i2, o := NewTreeNode("1,2,3"), NewTreeNode("1,null,3"), false
  fmt.Printf("Input:  %s, %s\n", i1.Sprintf(), i2.Sprintf())
  fmt.Printf("Output: %t\n", isSameTree(i1, i2))
  fmt.Printf("Expect: %t\n", o)

  i1, i2, o = NewTreeNode("1,2,3"), NewTreeNode("1,2,3"), true
  fmt.Printf("Input:  %s, %s\n", i1.Sprintf(), i2.Sprintf())
  fmt.Printf("Output: %t\n", isSameTree(i1, i2))
  fmt.Printf("Expect: %t\n", o)

  i1, i2, o = NewTreeNode("10,5,15"), NewTreeNode("10,5,null,null,15"), false
  fmt.Printf("Input:  %s, %s\n", i1.Sprintf(), i2.Sprintf())
  fmt.Printf("Output: %t\n", isSameTree(i1, i2))
  fmt.Printf("Expect: %t\n", o)
}

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

// -- end --

