package main
import "fmt"
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func main() {
  i, o := []int{-10,-3,0,5,9}, NewTreeNode("0,-3,9,-10,null,5")
  fmt.Printf("Input:  %v\n", i)
  fmt.Printf("Output: %s\n", sortedArrayToBST(i).ToStr())
  fmt.Printf("Expect: %s\n", o.ToStr())
}

// T: O(n)
// M: O(log(n))
// -- start --

func sortedArrayToBST(nums []int) *TreeNode {
  if len(nums) == 0 {
    return nil
  }
  mid := len(nums) / 2
  return &TreeNode{
    Val: nums[mid],
    Left: sortedArrayToBST(nums[:mid]),
    Right: sortedArrayToBST(nums[mid+1:]) }
}

// -- end --

