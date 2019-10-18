package main
import "fmt"
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i, o := NewTreeNode("6,2,8,0,4,7,9,null,null,3,5"), 6
  p, q := i.Left, i.Right
  fmt.Printf("Input:  %s p=%d q=%d\n", i.ToStr(), p.Val, q.Val)
  fmt.Printf("Output: %d\n", lowestCommonAncestor(i, p, q).Val)
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// T: O(1)
// -- start --

// T: O(n)
// T: O(1)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
  if p.Val < root.Val && q.Val < root.Val {
    return lowestCommonAncestor(root.Left, p, q)
  }

  if p.Val > root.Val && q.Val > root.Val {
    return lowestCommonAncestor(root.Right, p, q)
  }

  return root
}

// T: O(n)
// M: O(n)
func lowestCommonAncestorArray(root, p, q *TreeNode) *TreeNode {
  if root == nil {
    return nil
  }

  a, _ := findPath(root, p)
  b, _ := findPath(root, q)

  l := Min(len(a), len(b))

  for i := 0; i < l; i++ {
    if a[i].Val != b[i].Val {
      return a[i-1]
    }
  }

  return a[l-1]
}

func findPath(root, n *TreeNode) ([]*TreeNode, bool) {
  if root == nil {
    return []*TreeNode{}, false
  }

  if root.Val == n.Val {
    return []*TreeNode{root}, true
  }

  l, ok1 := findPath(root.Left, n)
  r, ok2 := findPath(root.Right, n)

  if ok1 {
    return append([]*TreeNode{root}, l...), true
  }

  if ok2 {
    return append([]*TreeNode{root}, r...), true
  }

  return []*TreeNode{}, false
}

func Min(a, b int) int {
  if a > b {
    return b
  }
  return a
}

// -- end --

