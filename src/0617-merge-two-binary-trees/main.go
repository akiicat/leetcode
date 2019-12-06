package main
import "fmt"
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  t1, t2, o := NewTreeNode("1,3,2,5"), NewTreeNode("2,1,3,null,4,null,7"), NewTreeNode("3,4,5,5,4,null,7")
  fmt.Printf("Input:  %s, %s\n", t1.ToStr(), t2.ToStr())
  fmt.Printf("Output: %s\n", mergeTrees(t1, t2).ToStr())
  fmt.Printf("Expect: %s\n", o.ToStr())
}

// T: O(n)
// M: O(h)
// -- start --

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
  // if t1 == nil && t2 == nil {
  //   return nil
  // }

  if t2 == nil {
    return t1
  }

  if t1 == nil {
    return t2
  }

  t1.Val += t2.Val

  t1.Left = mergeTrees(t1.Left, t2.Left)
  t1.Right = mergeTrees(t1.Right, t2.Right)

  return t1
}

// -- end --

