package main
import "fmt"
import . "tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i, o := NewTreeNode("1,2,3,null,5"), 2
  fmt.Printf("Input:  %s\n", i.Sprintf())
  fmt.Printf("Output: %d\n", minDepth(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = NewTreeNode("3,9,20,null,null,15,7"), 2
  fmt.Printf("Input:  %s\n", i.Sprintf())
  fmt.Printf("Output: %d\n", minDepth(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = NewTreeNode("1,2,2,3,3,null,null,4,4"), 2
  fmt.Printf("Input:  %s\n", i.Sprintf())
  fmt.Printf("Output: %d\n", minDepth(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = NewTreeNode("1,2,2,3,null,null,3,4,null,null,4"), 4
  fmt.Printf("Input:  %s\n", i.Sprintf())
  fmt.Printf("Output: %d\n", minDepth(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = NewTreeNode("1,2"), 2
  fmt.Printf("Input:  %s\n", i.Sprintf())
  fmt.Printf("Output: %d\n", minDepth(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = NewTreeNode("1"), 1
  fmt.Printf("Input:  %s\n", i.Sprintf())
  fmt.Printf("Output: %d\n", minDepth(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = NewTreeNode(""), 0
  fmt.Printf("Input:  %s\n", i.Sprintf())
  fmt.Printf("Output: %d\n", minDepth(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(h)
// -- start --

func minDepth(root *TreeNode) int {
  if root == nil {
    return 0
  }

  if root.Left == nil && root.Right == nil {
    return 1
  }

  if root.Left == nil {
    return minDepth(root.Right) + 1
  }

  if root.Right == nil {
    return minDepth(root.Left) + 1
  }

  return Min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func Min(a, b int) int {
  if a < b {
    return a
  }
  return b
}


// -- end --

