package main
import "fmt"
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  s, t, o := NewTreeNode("3,4,5,1,2"), NewTreeNode(""), false
  fmt.Printf("Input:  %s %s\n", s.ToStr(), t.ToStr())
  fmt.Printf("Output: %t\n", isSubtree(s, t))
  fmt.Printf("Expect: %t\n", o)

  s, t, o = NewTreeNode("3,4,5,1,2"), NewTreeNode("4,1,2"), true
  fmt.Printf("Input:  %s %s\n", s.ToStr(), t.ToStr())
  fmt.Printf("Output: %t\n", isSubtree(s, t))
  fmt.Printf("Expect: %t\n", o)

  s, t, o = NewTreeNode("3,4,5,1,2,null,null,0"), NewTreeNode("4,1,2"), false
  fmt.Printf("Input:  %s %s\n", s.ToStr(), t.ToStr())
  fmt.Printf("Output: %t\n", isSubtree(s, t))
  fmt.Printf("Expect: %t\n", o)

  s, t, o = NewTreeNode("1,1"), NewTreeNode("1"), true
  fmt.Printf("Input:  %s %s\n", s.ToStr(), t.ToStr())
  fmt.Printf("Output: %t\n", isSubtree(s, t))
  fmt.Printf("Expect: %t\n", o)
}

// leetcode 100.
// T: O(m*n)
// M: O(h)
// -- start --

func isSubtree(s *TreeNode, t *TreeNode) bool {
  if t == nil || s == nil {
    return false
  }

  l := isSubtree(s.Left, t)
  r := isSubtree(s.Right, t)

  return l || r || isSameTree(s, t)
}

// 0100-same-tree
func isSameTree(s *TreeNode, t *TreeNode) bool {
  if s == nil && t == nil {
    return true
  }

  if s == nil || t == nil || s.Val != t.Val {
    return false
  }

  return isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}

// -- end --

