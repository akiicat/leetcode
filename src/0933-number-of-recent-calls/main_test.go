package main
import "testing"
import . "main/pkg/testing_helper"
import "fmt"
import "strconv"

func TestMinStack(_t *testing.T) {
  i1 := []string{"RecentCounter","ping","ping","ping","ping"}
  i2 := [][]int{[]int{},[]int{1},[]int{100},[]int{3001},[]int{3002}}
  o := []string{"null","1","2","3","3"}
  T(_t, S(i1, i2), S(run(i1, i2)), S(o))
}

func run(intrs []string, values [][]int) []string {
  var obj RecentCounter
  rtn := make([]string, len(intrs))

  for i, intr := range intrs {
    switch intr {
    case "RecentCounter":
      obj = Constructor()
      rtn[i] = "null"
    case "ping":
      rtn[i] = strconv.Itoa(obj.Ping(values[i][0]))
    default:
      fmt.Printf("Command Not Found\n")
      rtn[i] = "null"
    }
  }
  return rtn
}
