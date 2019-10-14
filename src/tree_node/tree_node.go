package list_node
import "fmt"
import "strings"
import "strconv"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func NewTreeNode(input string) *TreeNode {
  if len(input) == 0 {
    return nil
  }

  strs := strings.Split(input, ",")

  size := 0
  for _, v := range strs {
    if v != "null" {
      size++
    }
  }

  nodes := make([]TreeNode, size)

  cur := &nodes[0]
  cur.Val, _ = strconv.Atoi(strs[0])

  queue, par := []**TreeNode{&cur.Left, &cur.Right}, &cur

  next := 1
  for _, v := range strs[1:] {
    queue, par = queue[1:], queue[0]

    if v == "null" {
      continue
    }

    cur = &nodes[next]
    cur.Val, _ = strconv.Atoi(v)

    *par = cur

    queue = append(queue, &cur.Left, &cur.Right)
    next++
  }

  return &nodes[0]
}

// BFS
func (root *TreeNode) ToStr() string {
  if root == nil {
    return ""
  }

  var s strings.Builder
  var buf strings.Builder

  queue, node := []*TreeNode{root}, root

  fmt.Fprintf(&s, "%d", root.Val)

  for len(queue) != 0 {
    queue, node = queue[1:], queue[0]

    if node.Left == nil {
      fmt.Fprintf(&buf, ",null")
    } else {
      fmt.Fprintf(&s, "%s,%d", buf.String(), node.Left.Val)
      buf.Reset()
      queue = append(queue, node.Left)
    }

    if node.Right == nil {
      fmt.Fprintf(&buf, ",null")
    } else {
      fmt.Fprintf(&s, "%s,%d", buf.String(), node.Right.Val)
      buf.Reset()
      queue = append(queue, node.Right)
    }
  }

  return s.String()
}

