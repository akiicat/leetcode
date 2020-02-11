package main
import . "main/pkg/list_node"

// T: O(1)
// M: O(1)
// -- start --

func deleteNode(node *ListNode) {
  *node = *node.Next
}

func deleteNodeAssign(node *ListNode) {
  node.Val, node.Next = node.Next.Val, node.Next.Next
}

// -- end --

