package main
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

// T: O(n)
// M: O(n)
// -- start --

func levelOrderBottom(root *TreeNode) [][]int {
  stack := [][]int{}

  if root == nil {
    return stack
  }

  queue := []*TreeNode{root}

  for len(queue) != 0 {
    l := len(queue)
    arr := make([]int, l)

    for i := 0; i < l; i++ {
      node := queue[i]
      arr[i] = node.Val

      if node.Left != nil {
        queue = append(queue, node.Left)
      }

      if node.Right != nil {
        queue = append(queue, node.Right)
      }
    }

    queue = queue[l:]
    stack = append([][]int{arr}, stack...)
  }

  return stack
}


func levelOrderBottomRecursive(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	return bfs([]*TreeNode{root})
}

func bfs(roots []*TreeNode) [][]int {
	if len(roots) == 0 {
		return [][]int{}
	}

  arr := make([]int, 0, len(roots))
  queue := make([]*TreeNode, 0, 2*len(roots))

	for _, root := range roots {
		arr = append(arr, root.Val)

		if root.Left != nil {
			queue = append(queue, root.Left)
		}
		if root.Right != nil {
			queue = append(queue, root.Right)
		}
	}

	return append(bfs(queue), arr)
}

// -- end --

