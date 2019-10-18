package main
import "fmt"
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i1, i2, o := NewTreeNode("5,4,8,11,null,13,4,7,2,null,null,null,1"), 22, true
  fmt.Printf("Input:  %s %d\n", i1.ToStr(), i2)
  fmt.Printf("Output: %t\n", hasPathSum(i1, i2))
  fmt.Printf("Expect: %t\n", o)

  i1, i2, o = NewTreeNode("1,2"), 0, false
  fmt.Printf("Input:  %s %d\n", i1.ToStr(), i2)
  fmt.Printf("Output: %t\n", hasPathSum(i1, i2))
  fmt.Printf("Expect: %t\n", o)

  i1, i2, o = NewTreeNode("1,2"), 10, false
  fmt.Printf("Input:  %s %d\n", i1.ToStr(), i2)
  fmt.Printf("Output: %t\n", hasPathSum(i1, i2))
  fmt.Printf("Expect: %t\n", o)

  i1, i2, o = NewTreeNode("-2,null,-3"), -5, true
  fmt.Printf("Input:  %s %d\n", i1.ToStr(), i2)
  fmt.Printf("Output: %t\n", hasPathSum(i1, i2))
  fmt.Printf("Expect: %t\n", o)

  i1, i2, o = NewTreeNode(""), 0, false
  fmt.Printf("Input:  %s %d\n", i1.ToStr(), i2)
  fmt.Printf("Output: %t\n", hasPathSum(i1, i2))
  fmt.Printf("Expect: %t\n", o)

  i1, i2, o = NewTreeNode(""), 1, false
  fmt.Printf("Input:  %s %d\n", i1.ToStr(), i2)
  fmt.Printf("Output: %t\n", hasPathSum(i1, i2))
  fmt.Printf("Expect: %t\n", o)
}

// T: O(n)
// M: O(h)
// -- start --

// Recursive Depth First Search
func hasPathSum(root *TreeNode, sum int) bool {
  if root == nil {
    return false
  }

  // Leaf
  if root.Left == nil && root.Right == nil {
    return sum == root.Val
  }

  return hasPathSum(root.Left, sum - root.Val) || hasPathSum(root.Right, sum - root.Val)
}

// -- end --

