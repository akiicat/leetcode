package main
import "fmt"
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i, o := NewTreeNode("1,2,3"), 1
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %d\n", findTilt(i))
  fmt.Printf("Expect: %d\n", o)

  i, o = NewTreeNode("1,2,3,4,5"), 9
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %d\n", findTilt(i))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(n)
// -- start --

func findTilt(root *TreeNode) int {
  sum := 0
  SumTree(root, &sum)
  return sum
}

func SumTree(root *TreeNode, sum *int) int {
  if root == nil {
    return 0
  }

  l := SumTree(root.Left, sum)
  r := SumTree(root.Right, sum)

  *sum += Abs(l, r)

  return l + r + root.Val
}

func Abs(a, b int) int {
  if a > b {
    return a - b
  }
  return b - a
}

// -- end --

