package main
import "fmt"
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i, o := NewTreeNode("1,null,2,2"), []int{2}
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %v\n", findMode(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = NewTreeNode("2147483647"), []int{2147483647}
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %v\n", findMode(i))
  fmt.Printf("Expect: %v\n", o)

}

// https://leetcode.com/problems/find-mode-in-binary-search-tree/discuss/329452/Golang-beats-100-O(1)-space
//
// To have O(1) space, one needs to keep track of the previously visited value.
// One can do this during traversal using the following (you'll find this extremely useful in many BST problems requiring O(1) space):
// 
// ```
// DFS(root, prev)
//     DFS(root.Left,prev)
// 	Update prev to current root value
// 	DFS(root.Right,prev)
// ```
// 
// To find BST's Mode value while traversal, one needs to keep track of number of nodes with same values, as follows:
// 
// ```
// DFS(root, prev)
//     DFS(root.Left,prev)
// 	if root.Val == prev {count++} else {count=1}
// 	Update prev to current root value
// 	DFS(root.Right,prev)
// ```
// 
// Now all that's left is to keep track of max length. If count becomes greater than max, then we update max, and wipe off the existing result array and create a new one using the new element.
//
// T: O(n)
// M: O(1)
// -- start --

// DFS
// T: O(n)
// M: O(1)
func findMode(root *TreeNode) []int {
  if root == nil {
    return []int{}
  }

  prev := -1<<31
  count := 0
  max := 1
  result := make([]int, 0)

  DFS(root, &prev, &count, &max, &result)

  return result
}

func DFS(root *TreeNode, prev, count, max *int, result *[]int) {
  if root == nil {
    return
  }

  DFS(root.Left, prev, count, max, result)

  if root.Val == *prev {
    *count++
  } else {
    *count = 1
  }

  if *count > *max {
    *result = []int{root.Val}
    *max = *count
  } else if *max == *count {
    *result = append(*result, root.Val)
  }
  *prev = root.Val


  DFS(root.Right, prev, count, max, result)
}


// BFS
// T: O(n)
// M: O(n)
func findModeMap(root *TreeNode) []int {
  m := make(map[int]int)
  m = find(root, m)

  max := 0
  for _, v := range m {
    if v > max {
      max = v
    }
  }

  rtn := []int{}
  for k, v := range m {
    if v == max {
      rtn = append(rtn, k)
    }
  }

  return rtn
}

func find(root *TreeNode, m map[int]int) map[int]int {
  if root == nil {
    return m
  }

  m[root.Val]++
  find(root.Left, m)
  find(root.Right, m)

  return m
}

// BFS
// T: O(n)
// M: O(n)
func findModeArray(root *TreeNode) []int {
  arr := R(root)

  m := map[int]int{}

  max := 0
  for _, v := range arr {
    m[v]++
    if m[v] > max {
      max = m[v]
    }
  }

  rtn := []int{}
  for k, v := range m {
    if v == max {
      rtn = append(rtn, k)
    }
  }

  return rtn
}

func R(root *TreeNode) []int {
  if root == nil {
    return []int{}
  }

  l := R(root.Left)
  r := R(root.Right)

  return append(append(l, r...), root.Val)
}

// -- end --

