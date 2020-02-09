package testing_helper

// import . "main/pkg/tree_node"
// import . "main/pkg/list_node"
import "fmt"
import "testing"

func S(i ...interface{}) string {
  if len(i) > 1 {
    return fmt.Sprintf("%s %s", S(i[0]), S(i[1:]...))
  }

  switch v := i[0].(type) {
  case []string:
    return fmt.Sprintf("%v", v)
  case string:
    return fmt.Sprintf("%s", v)
  case int:
    return fmt.Sprintf("%d", v)
  default:
    return fmt.Sprintf("failed at type %T", v)
  }
}

func T(t *testing.T, i, r, o string) {
  if r != o {
    t.Errorf("\nInput:  %s\nOutput: %s\nExpect: %s\n", i, r, o)
  }
}

