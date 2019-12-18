package main
import "fmt"
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i, l, r, o := NewTreeNode("1,0,2"), 1, 2, NewTreeNode("1,null,2")
  fmt.Printf("Input:  %s, l=%d, r=%d\n", i.ToStr(), l, r)
  fmt.Printf("Output: %s\n", trimBST(i, l, r).ToStr())
  fmt.Printf("Expect: %s\n", o.ToStr())

  i, l, r, o = NewTreeNode("3,0,4,null,2,null,null,1"), 1, 3, NewTreeNode("3,2,null,1")
  fmt.Printf("Input:  %s, l=%d, r=%d\n", i.ToStr(), l, r)
  fmt.Printf("Output: %s\n", trimBST(i, l, r).ToStr())
  fmt.Printf("Expect: %s\n", o.ToStr())
}

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

