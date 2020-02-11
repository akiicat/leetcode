package main
import "strconv"
import . "main/pkg/tree_node"

// T: O(n)
// M: O(log(n))
// -- start --

// T: O(n)
// M: O(log(n))
func binaryTreePaths(root *TreeNode) []string {
  if root == nil {
    return []string{}
  }

  if root.Left == nil && root.Right == nil {
    return []string{strconv.Itoa(root.Val)}
  }

  l := binaryTreePaths(root.Left)
  r := binaryTreePaths(root.Right)

  for i, v := range l {
    l[i] = strconv.Itoa(root.Val) + "->" + v
  }

  for i, v := range r {
    r[i] = strconv.Itoa(root.Val) + "->" + v
  }

  return append(l, r...)
}

// -- end --

