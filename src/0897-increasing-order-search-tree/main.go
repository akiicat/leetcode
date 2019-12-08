package main
import "fmt"
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i, o := NewTreeNode("5,3,6,2,4,null,8,1,null,null,null,7,9"), NewTreeNode("1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9")
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %s\n", increasingBST(i).ToStr())
  fmt.Printf("Expect: %s\n", o.ToStr())
}

// T: O(n)
// M: O(h)
// -- start --

// T: O(n)
// M: O(h)
func increasingBST(root *TreeNode) *TreeNode {
  head := &TreeNode{}
  tail := head
  R2(root, tail)
  return head.Right
}

func R2(root *TreeNode, tail *TreeNode) *TreeNode {
  if root == nil {
    return tail
  }

  tail = R2(root.Left, tail)

  root.Left = nil
  tail.Right, tail = root, root

  tail = R2(root.Right, tail)

  return tail
}

// T: O(n)
// M: O(h)
func increasingBSTReverseRecursive(root *TreeNode) *TreeNode {
  var child *TreeNode
  R1(root, &child)
  return child
}

func R1(root *TreeNode, child **TreeNode) {
  if root == nil {
    return
  }

  R1(root.Right, child)

  root.Right = *child
  *child = root

  R1(root.Left, child)

  root.Left = nil

  return
}

// -- end --

