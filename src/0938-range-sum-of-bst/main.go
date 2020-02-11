package main
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

// T: O(n)
// M: O(1)
// -- start --

// Top Bottom
// T: O(n)
// M: O(h)
func rangeSumBST(root *TreeNode, L int, R int) int {
  var sum int
  var node *TreeNode

  stack := []*TreeNode{root}

  for len(stack) > 0 {
    node, stack = stack[len(stack)-1], stack[:len(stack)-1]

    if node == nil {
      continue
    }

    if L <= node.Val && node.Val <= R {
      sum += node.Val
    }
    if L < node.Val {
      stack = append(stack, node.Left)
    }
    if node.Val < R {
      stack = append(stack, node.Right)
    }
  }

  return sum
}

// T: O(n)
// M: O(1)
func rangeSumBSTMergeAlgorithm(root *TreeNode, L int, R int) int {
  sum := 0

  if root == nil {
    return 0
  }

  if L <= root.Val && root.Val <= R {
    sum += root.Val
  }

  if L <= root.Val {
    sum += rangeSumBSTMergeAlgorithm(root.Left, L, R)
  }

  if root.Val <= R {
    sum += rangeSumBSTMergeAlgorithm(root.Right, L, R)
  }

  return sum
}

// T: O(n)
// M: O(1)
func rangeSumBSTDFSButtomUpRecursive(root *TreeNode, L int, R int) int {
  if root == nil {
    return 0
  }

  if root.Val < L {
    return rangeSumBST(root.Right, L, R)
  }

  if root.Val > R {
    return rangeSumBST(root.Left, L, R)
  }

  return root.Val +
         rangeSumBST(root.Left, L, R) +
         rangeSumBST(root.Right, L, R)
}

// -- end --

