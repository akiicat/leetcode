package main
import "fmt"
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i, o := NewTreeNode("1,2,3,null,5"), [][]int{[]int{5}, []int{2,3}, []int{1}}
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %v\n", levelOrderBottomRecursive(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = NewTreeNode("1,2,3,4,5"), [][]int{[]int{4,5}, []int{2,3}, []int{1}}
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %v\n", levelOrderBottomRecursive(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = NewTreeNode("3,9,20,null,null,15,7"), [][]int{[]int{15,7}, []int{9,20}, []int{3}}
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %v\n", levelOrderBottomRecursive(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = NewTreeNode("1,2,2,3,3,null,null,4,4"), [][]int{[]int{4,4},[]int{3,3},[]int{2,2},[]int{1}}
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %v\n", levelOrderBottomRecursive(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = NewTreeNode("1,2,2,3,null,null,3,4,null,null,4"), [][]int{[]int{4,4},[]int{3,3},[]int{2,2},[]int{1}}
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %v\n", levelOrderBottomRecursive(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = NewTreeNode(""), [][]int{}
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %v\n", levelOrderBottomRecursive(i))
  fmt.Printf("Expect: %v\n", o)
}

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

