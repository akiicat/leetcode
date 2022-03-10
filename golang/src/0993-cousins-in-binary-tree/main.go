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

func isCousins(root *TreeNode, x int, y int) bool {
  px, py := 0, 0
  return Depth(root, x, 0, &px) == Depth(root, y, 0, &py) && px != py
}

func Depth(root *TreeNode, match int, depth int, parent *int) int {
  if root == nil {
    return 0
  }

  if root.Val == match {
    return depth
  }

  *parent = root.Val
  l := Depth(root.Left, match, depth+1, parent)
  if l != 0 {
    return l
  }

  *parent = root.Val
  r := Depth(root.Right, match, depth+1, parent)
  if r != 0 {
    return r
  }

  return 0
}

// -- end --

