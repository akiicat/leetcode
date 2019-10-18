package main
import "fmt"
import "strconv"
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i, o := NewTreeNode("1,2,3,null,5"), []string{"1->2->5", "1->3"}
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %v\n", binaryTreePaths(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = NewTreeNode("3,9,20,null,null,15,7"), []string{"3->9","3->20->15","3->20->7"}
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %v\n", binaryTreePaths(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = NewTreeNode("1,2,2,3,3,null,null,4,4"), []string{"1->2->3->4","1->2->3->4","1->2->3","1->2"}
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %v\n", binaryTreePaths(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = NewTreeNode("1,2,2,3,null,null,3,4,null,null,4"), []string{"1->2->3->4","1->2->3->4"}
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %v\n", binaryTreePaths(i))
  fmt.Printf("Expect: %v\n", o)

  i, o = NewTreeNode(""), []string{}
  fmt.Printf("Input:  %s\n", i.ToStr())
  fmt.Printf("Output: %v\n", binaryTreePaths(i))
  fmt.Printf("Expect: %v\n", o)
}

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

