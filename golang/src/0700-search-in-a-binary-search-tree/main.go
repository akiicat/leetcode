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

