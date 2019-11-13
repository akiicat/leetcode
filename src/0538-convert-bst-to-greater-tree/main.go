package main
import "fmt"
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i, o := NewTreeNode("5,2,13"), NewTreeNode("18,20,13")
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %s\n", convertBST(i).ToStr())
  fmt.Printf("Expect: %s\n", o.ToStr())
}

// T: O(n)
// M: O(h)
// -- start --

func convertBST(root *TreeNode) *TreeNode {
  prev := 0
  DFS(root, &prev)
  return root
}

func DFS(root *TreeNode, prev *int) {
  if root == nil {
    return
  }

  DFS(root.Right, prev)

  *prev += root.Val
  root.Val = *prev

  DFS(root.Left, prev)
}

// -- end --

