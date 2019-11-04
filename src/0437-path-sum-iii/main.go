package main
import "fmt"
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i, s, o := NewTreeNode("10,5,-3,3,2,null,11,3,-2,null,1"), 8, 3
  fmt.Printf("Input:  %s sum=%d\n", i.ToStr(), s)
  fmt.Printf("Output: %d\n", pathSum(i, s))
  fmt.Printf("Expect: %d\n", o)

  i, s, o = NewTreeNode("1,-2,-3"), -1, 1
  fmt.Printf("Input:  %s sum=%d\n", i.ToStr(), s)
  fmt.Printf("Output: %d\n", pathSum(i, s))
  fmt.Printf("Expect: %d\n", o)
}

// T: O(n)
// M: O(h)
// -- start --

// T: O(n)
// T: O(h)
func pathSum(root *TreeNode, sum int) int {
  m := make(map[int]int, 0)
  m[0] = 1

  return DFS(root, 0, sum, m)
}

func DFS(root *TreeNode, cur, sum int, m map[int]int) int {
  if root == nil {
    return 0
  }

  cur += root.Val
  count := m[cur-sum]

  m[cur]++

  count += DFS(root.Left, cur, sum, m)
  count += DFS(root.Right, cur, sum, m)

  m[cur]--

  return count
}

// T: O(n)
// M: O(h)
func pathSumRecursive(root *TreeNode, sum int) int {
    if root == nil { return 0 }
    return sumUp(root, 0, sum) + pathSumRecursive(root.Left, sum) + pathSumRecursive(root.Right, sum)
}

func sumUp(root *TreeNode, pre int, sum int) int {
    if root == nil { return 0 }
    current := pre + root.Val

    c:= 0

    if current == sum { c = 1 }
    return c + sumUp(root.Left, current, sum) + sumUp(root.Right, current, sum)
}


// T: O(n ^ 2)
// M: O(n)
func pathSumArray(root *TreeNode, sum int) int {
  return R(root, []int{}, sum)
}

func R(root *TreeNode, prev []int, sum int) int {
  if root == nil {
    return 0
  }

  prev = append(prev, root.Val)

  c := 0
  for i := 0; i < len(prev); i++ {
    s := 0
    for j := i; j < len(prev); j++ {
      s += prev[j]
    }
    if s == sum {
      c++
    }
  }

  return c + R(root.Left, prev, sum) + R(root.Right, prev, sum)
}

// -- end --

