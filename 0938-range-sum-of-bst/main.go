package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func main() {
  // 10,5,15,3,7,null,18
  //                10
  //              /    \
  //           5         15
  //          / \      /    \
  //        3    7    null  18
  i := &TreeNode{Val: 10}
  i.Left = &TreeNode{Val: 5}
  i.Right = &TreeNode{Val: 15}
  i.Left.Left = &TreeNode{Val: 3}
  i.Left.Right = &TreeNode{Val: 7}
  i.Right.Left = nil
  i.Right.Right = &TreeNode{Val: 18}
  L, R, o := 7, 15, 32
  fmt.Printf("Input:  %s, %d, %d\n", i, L, R)
  fmt.Printf("Output: %d\n", rangeSumBST(i, L, R))
  fmt.Printf("Expect: %d\n", o)

  // 10,5,15,3,7,13,18,1,null,6
  //                  10
  //                /    \
  //             5         15
  //            / \      /    \
  //        3       7    13    18
  //       / \     /
  //      1  null 6
  i = &TreeNode{Val: 10}
  i.Left = &TreeNode{Val: 5}
  i.Right = &TreeNode{Val: 15}
  i.Left.Left = &TreeNode{Val: 3}
  i.Left.Right = &TreeNode{Val: 7}
  i.Right.Left = &TreeNode{Val: 13}
  i.Right.Right = &TreeNode{Val: 18}
  i.Left.Left.Left = &TreeNode{Val: 1}
  i.Left.Left.Right = nil
  i.Left.Right.Left = &TreeNode{Val: 6}
  L, R, o = 6, 10, 23
  fmt.Printf("Input:  %s, %d, %d\n", i, L, R)
  fmt.Printf("Output: %d\n", rangeSumBST(i, L, R))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(1)
// -- start --

func rangeSumBST(root *TreeNode, L int, R int) int {
  return rangeSumBSTStackTopBottom(root, L, R)
}

// T: O(n)
// M: O(h)
func rangeSumBSTStackTopBottom(root *TreeNode, L int, R int) int {
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

