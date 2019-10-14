package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func main() {
  //     5
  //    / \
  //   3   6
  //  / \   \
  // 2   4   7
  t := &TreeNode{Val: 5}
  t.Left = &TreeNode{Val: 3}
  t.Right = &TreeNode{Val: 6}
  t.Left.Left = &TreeNode{Val: 2}
  t.Left.Right = &TreeNode{Val: 4}
  t.Right.Right = &TreeNode{Val: 7}

  fmt.Printf("Input:  [5, 3, 6, 2, 4, 0, 7]\nOutput: %t\nExpect: true\n" , findTarget(t, 9))
}

// T: O(N)
// M: O(N)
// -- start --

func findTarget(root *TreeNode, k int) bool {
  m := make(map[int]bool)

  return r(root, k, m)
}

func r(root *TreeNode, k int, m map[int]bool) bool {
  if root == nil {
    return false
  }

  _, ok := m[root.Val]
  if ok {
    return true
  }

  m[k - root.Val] = true

  return r(root.Left, k, m) || r(root.Right, k, m)
}

// -- end --

