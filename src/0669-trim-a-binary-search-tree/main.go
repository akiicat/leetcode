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

func trimBST(root *TreeNode, L int, R int) *TreeNode {
  if root == nil {
    return nil
  }

  l := trimBST(root.Left, L, R)
  r := trimBST(root.Right, L, R)

  if root.Val < L {
    return r
  }

  if root.Val > R {
    return l
  }

  root.Left = l
  root.Right = r

  return root
}

// -- end --

