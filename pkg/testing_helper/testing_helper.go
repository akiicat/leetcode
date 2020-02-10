package testing_helper

// import . "main/pkg/tree_node"
import . "main/pkg/list_node"
import "fmt"
import "testing"

type Value struct {
  Val string
  Err error
}

func S(i ...interface{}) *Value {
  if len(i) > 1 {
    a, b := S(i[0]), S(i[1:]...)

    if a.Err != nil {
      return a
    }
    if b.Err != nil {
      return b
    }

    return &Value{Val: a.Val + " " + b.Val}
  }

  switch v := i[0].(type) {
  case []string:
    return &Value{Val: fmt.Sprintf("%v", v)}
  case string:
    return &Value{Val: fmt.Sprintf("%s", v)}
  case int:
    return &Value{Val: fmt.Sprintf("%d", v)}
  case *ListNode:
    return &Value{Val: fmt.Sprintf("%s", v.ToStr())}
  default:
    return &Value{Err: fmt.Errorf("failed at type %T", v)}
  }
}

func T(t *testing.T, i, r, o *Value) {

  if i.Err != nil {
    t.Fatalf(i.Err.Error())
  }
  if r.Err != nil {
    t.Fatalf(r.Err.Error())
  }
  if o.Err != nil {
    t.Fatalf(o.Err.Error())
  }

  if r.Val != o.Val {
    t.Errorf("\nInput:  %s\nOutput: %s\nExpect: %s\n", i.Val, r.Val, o.Val)
  }
}

