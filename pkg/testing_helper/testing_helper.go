package testing_helper

import "fmt"
import "sort"
import "testing"
import . "main/pkg/list_node"
import . "main/pkg/tree_node"
import . "main/pkg/nested_integer"

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
  case []interface{}:
    return &Value{Val: fmt.Sprintf("%v", v)}
  case []string:
    return &Value{Val: fmt.Sprintf("%v", v)}
  case string:
    return &Value{Val: fmt.Sprintf("%s", v)}
  case [][]int:
    return &Value{Val: fmt.Sprintf("%v", v)}
  case []int:
    return &Value{Val: fmt.Sprintf("%v", v)}
  case int:
    return &Value{Val: fmt.Sprintf("%d", v)}
  case int32:
    return &Value{Val: fmt.Sprintf("%d", v)}
  case uint32:
    return &Value{Val: fmt.Sprintf("%d(%032b)", v, v)}
  case [][]byte: // [][]uint8
    return &Value{Val: fmt.Sprintf("%s", v)}
  case []byte: // []uint8
    return &Value{Val: fmt.Sprintf("%s", v)}
  case byte: // uint8
    return &Value{Val: fmt.Sprintf("%c", v)}
  case []float64:
    return &Value{Val: fmt.Sprintf("%v", v)}
  case float64:
    return &Value{Val: fmt.Sprintf("%f", v)}
  case []bool:
    return &Value{Val: fmt.Sprintf("%t", v)}
  case bool:
    return &Value{Val: fmt.Sprintf("%t", v)}
  case *ListNode:
    return &Value{Val: fmt.Sprintf("%s", v.ToStr())}
  case *TreeNode:
    return &Value{Val: fmt.Sprintf("%s", v.ToStr())}
  case *NestedInteger:
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

func Sort(i interface{}) interface{} {
  switch v := i.(type) {
  case []string:
    sort.Strings(v)
    return v
  case []int:
    sort.Ints(v)
    return v
  case [][]int:
    sort.Slice(v[:], func(i, j int) bool {
      if len(v[i]) != len(v[j]) {
        return len(v[i]) < len(v[j])
      }

      for x := range v[i] {
        if v[i][x] != v[j][x] {
          return v[i][x] < v[j][x]
        }
      }

      return false
    })
    return v
  default:
    return fmt.Errorf("Sort: failed at type %T", v)
  }
}

