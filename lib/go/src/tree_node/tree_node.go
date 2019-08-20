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

  arr := strings.Split(input, ",")

  head := make([]TreeNode, len(arr))

  for i, v := range arr {
    if v == "null" {
      continue
    }

    cur := &head[i]
    cur.Val, _ = strconv.Atoi(v)

    if i == 0 {
      continue
    }

    par := &head[(i-1)/2]
    isLeft := i % 2 == 1

    if isLeft {
      par.Left = cur
    } else {
      par.Right = cur
    }
  }

  return &head[0]
}

func (root *TreeNode) Printf() {
  fmt.Printf("%s", root.Sprintf())
}

// BFS
func (root *TreeNode) Sprintf() string {
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

