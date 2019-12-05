package main
import "fmt"
import "strconv"
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i, o := NewTreeNode("1,2,3,4"), "1(2(4))(3)"
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %s\n", tree2str(i))
  fmt.Printf("Expect: %s\n", o)

  i, o = NewTreeNode("1,2,3,null,4"), "1(2()(4))(3)"
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %s\n", tree2str(i))
  fmt.Printf("Expect: %s\n", o)
}

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

