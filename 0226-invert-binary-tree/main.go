package main
import "fmt"
import . "tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i, o := NewTreeNode("4,2,7,1,3,6,9"), NewTreeNode("4,7,2,9,6,3,1")
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %s\n", invertTree(i).ToStr())
  fmt.Printf("Expect: %s\n", o.ToStr())
}

// T: O(n)
// M: O(h)
// -- start --

func invertTree(root *TreeNode) *TreeNode {
  if root == nil {
    return root
  }

  invertTree(root.Left)
  invertTree(root.Right)

  root.Left, root.Right = root.Right, root.Left

  return root
}

// -- end --

