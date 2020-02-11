package main
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

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

