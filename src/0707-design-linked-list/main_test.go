package main
import "testing"
import . "main/pkg/testing_helper"
import "fmt"
import "strconv"

func TestMyLinkedList(_t *testing.T) {
  i1 := []string{"MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"}
  i2 := [][]int{[]int{},[]int{1},[]int{3},[]int{1,2},[]int{1},[]int{1},[]int{1}}
  o := []string{"null","null","null","null","2","null","3"}
  T(_t, S(i1, i2), S(run(i1, i2)), S(o))

  i1 = []string{"MyLinkedList","addAtIndex","addAtIndex","addAtIndex","get"}
  i2 = [][]int{[]int{},[]int{0,10},[]int{0,20},[]int{1,30},[]int{0}}
  o = []string{"null","null","null","null","20"}
  T(_t, S(i1, i2), S(run(i1, i2)), S(o))

  i1 = []string{"MyLinkedList","addAtHead","deleteAtIndex","addAtHead","addAtHead","addAtHead","addAtHead","addAtHead","addAtTail","get","deleteAtIndex","deleteAtIndex"}
  i2 = [][]int{[]int{},[]int{2},[]int{1},[]int{2},[]int{7},[]int{3},[]int{2},[]int{5},[]int{5},[]int{5},[]int{6},[]int{4}}
  o = []string{"null","null","null","null","null","null","null","null","null","2","null","null"}
  T(_t, S(i1, i2), S(run(i1, i2)), S(o))

  i1 = []string{"MyLinkedList","addAtHead","get","addAtHead","addAtHead","deleteAtIndex","addAtHead","get","get","get","addAtHead","deleteAtIndex"}
  i2 = [][]int{[]int{},[]int{4},[]int{1},[]int{1},[]int{5},[]int{3},[]int{7},[]int{3},[]int{3},[]int{3},[]int{1},[]int{4}}
  o = []string{"null","null","null","null","null","null","null","null","null","2","null","null"}
  T(_t, S(i1, i2), S(run(i1, i2)), S(o))

}

func run(intrs []string, values [][]int) []string {
  var obj MyLinkedList
  rtn := make([]string, len(intrs))

  for i, intr := range intrs {
    switch intr {
    case "MyLinkedList":
      obj = Constructor()
      rtn[i] = "null"
    case "get":
      rtn[i] = strconv.Itoa(obj.Get(values[i][0]))
    case "addAtHead":
      obj.AddAtHead(values[i][0])
      rtn[i] = "null"
    case "addAtTail":
      obj.AddAtTail(values[i][0])
      rtn[i] = "null"
    case "addAtIndex":
      obj.AddAtIndex(values[i][0], values[i][1])
      rtn[i] = "null"
    case "deleteAtIndex":
      obj.DeleteAtIndex(values[i][0])
      rtn[i] = "null"
    default:
      fmt.Printf("Command Not Found\n")
      rtn[i] = "null"
    }
  }
  return rtn
}
