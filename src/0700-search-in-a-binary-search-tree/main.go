package main
import "fmt"
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i, n, o := NewTreeNode("4,2,7,1,3"), 2, NewTreeNode("2,1,3")
  fmt.Printf("Input:  %s, %d\n", i.ToStr(), n)
  fmt.Printf("Output: %s\n", searchBST(i, n).ToStr())
  fmt.Printf("Expect: %s\n", o.ToStr())
}

// T: O(n)
// M: O(h)
// -- start --

func searchBST(root *TreeNode, val int) *TreeNode {
  if root == nil {
    return nil
  }

  if root.Val > val {
    return searchBST(root.Left, val)
  }

  if root.Val < val {
    return searchBST(root.Right, val)
  }

  return root
}

// -- end --

