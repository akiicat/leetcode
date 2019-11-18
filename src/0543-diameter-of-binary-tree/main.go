package main
import "fmt"
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i, o := NewTreeNode("1,2,3,4,5"), 3
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %d\n", diameterOfBinaryTree(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = NewTreeNode(""), 0
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %d\n", diameterOfBinaryTree(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = NewTreeNode("4,-7,-3,null,null,-9,-3,9,-7,-4,null,6,null,-6,-6,null,null,0,6,5,null,9,null,null,-1,-4,null,null,null,-2"), 8
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %d\n", diameterOfBinaryTree(i))
  fmt.Printf("Expect: %d\n", o)
}

// leetcode 104.
// T: O(n)
// M: O(h)
// -- start --

func diameterOfBinaryTree(root *TreeNode) int {
  max := 0
  maxDepth(root, &max)
  return max
}

// 0104-maximum-depth-of-binary-tree
func maxDepth(root *TreeNode, max *int) int {
  if root == nil {
    return 0
  }
  l := maxDepth(root.Left, max)
  r := maxDepth(root.Right, max)
  *max = Max(*max, l + r)
  return Max(l, r) + 1
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

// -- end --

