package main
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

// T: O(n)
// M: O(h)
// -- start --

// T: O(n)
// M: O(n)
func minDiffInBST(root *TreeNode) int {
  m := []int{}

  R(root, &m)

  min := 1<<31
  for i := 0; i < len(m)-1; i++ {
    min = Min(min, m[i+1] - m[i])
  }

  return min
}

func R(root *TreeNode, m *[]int) {
  if root == nil {
    return
  }

  R(root.Left, m)
  *m = append(*m, root.Val)
  R(root.Right, m)
}

// the diff between root.Val and -1<<31 cannot smaller than the answer
// T: O(n)
// M: O(h)
func minDiffInBSTDFS(root *TreeNode) int {
  prev := -1<<31
  return R1(root, &prev)
}

func R1(root *TreeNode, prev *int) int {
  if root == nil {
    return 1<<31
  }

  l := R1(root.Left, prev)

  if *prev != -1<<31 {
    l = Min(l, root.Val - *prev)
  }
  *prev = root.Val

  r := R1(root.Right, prev)

  return Min(l, r)
}

func Min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

// -- end --

