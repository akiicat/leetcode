package main
import "fmt"
import . "tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  //     1
  //    / \
  //   2   2
  //  / \ / \
  // 3  4 4  3
  i, o := NewTreeNode("1,2,2,3,4,4,3"), true
  fmt.Printf("Input:  %s\n", i.Sprintf())
  fmt.Printf("Output: %t\n", isSymmetric(i))
  fmt.Printf("Expect: %t\n", o)

  //   1
  //  / \
  // 2   2
  //  \   \
  //  3    3
  i, o = NewTreeNode("1,2,2,null,3,null,3"), false
  fmt.Printf("Input:  %s\n", i.Sprintf())
  fmt.Printf("Output: %t\n", isSymmetric(i))
  fmt.Printf("Expect: %t\n", o)

  i, o = NewTreeNode(""), true
  fmt.Printf("Input:  %s\n", i.Sprintf())
  fmt.Printf("Output: %t\n", isSymmetric(i))
  fmt.Printf("Expect: %t\n", o)
}

// T: O(n)
// M: O(n)
// -- start --

func isSymmetric(root *TreeNode) bool {
  return isSymmetricIterative(root)
  return isSymmetricRecursive(root, root)
}

func isSymmetricIterative(root *TreeNode) bool {
  if root == nil {
    return true
  }

  stack, l, r := []*TreeNode{root, root}, root, root

  for len(stack) != 0 {
    stack, l, r = stack[:len(stack)-2], stack[len(stack)-2], stack[len(stack)-1]

    if l.Val != r.Val {
      return false
    }

    ll, rr := l.Left != nil, r.Right != nil
    lr, rl := l.Right != nil, r.Left != nil
    if ll && rr {
      stack = append(stack, l.Left, r.Right)
    }

    if lr && rl {
      stack = append(stack, l.Right, r.Left)
    }

    if ll != rr || lr != rl {
      return false
    }
  }
  return true
}

func isSymmetricRecursive(l *TreeNode, r *TreeNode) bool {
  if l == nil && r == nil {
    return true
  }

  if l == nil || r == nil {
    return false
  }

  return l.Val == r.Val && isSymmetricRecursive(l.Left, r.Right) && isSymmetricRecursive(l.Right, r.Left)
}

// -- end --

