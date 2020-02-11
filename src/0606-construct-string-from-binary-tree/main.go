package main
import "strconv"
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

// T: O(n)
// M: O(h)
// -- start --

func tree2str(t *TreeNode) string {
  if t == nil {
    return ""
  }

  l := tree2str(t.Left)
  r := tree2str(t.Right)

  rtn := strconv.Itoa(t.Val)
  switch {
  case r != "":
    rtn += "(" + l + ")" + "(" + r + ")"
  case l != "":
    rtn += "(" + l + ")"
  }

  return rtn
}

// -- end --

