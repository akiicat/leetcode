package main
import "fmt"
import . "tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i, o := NewTreeNode("3,9,20,null,null,15,7"), 24
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %d\n", sumOfLeftLeaves(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = NewTreeNode("1"), 0
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %d\n", sumOfLeftLeaves(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(h)
// -- start --

// T: O(n)
// M: O(h)
func sumOfLeftLeaves(root *TreeNode) int {
  return r(root, false)
}

func r(root *TreeNode, isLeft bool) int {
  if root == nil {
    return 0
  }

  if isLeft && root.Left == nil && root.Right == nil {
    return root.Val
  }

  return r(root.Left, true) + r(root.Right, false)
}

// T: O(n)
// M: O(n)
func sumOfLeftLeavesStack(root *TreeNode) int {
  stack := []*TreeNode{nil, root}
  sum := 0

  for len(stack) != 0 {
    l, r := stack[0], stack[1]

    stack = stack[2:]

    if l != nil {
      if l.Left == nil && l.Right == nil {
        sum += l.Val
      } else {
        stack = append(stack, l.Left, l.Right)
      }
    }

    if r != nil {
      if !(r.Left == nil && r.Right == nil) {
        stack = append(stack, r.Left, r.Right)
      }
    }
  }

  return sum
}

// -- end --

