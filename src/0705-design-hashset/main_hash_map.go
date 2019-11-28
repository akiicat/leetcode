package main
import "fmt"

func main() {
  i1 := []string{"MyHashSet","add","add","contains","contains","add","contains","remove","contains"}
  i2 := [][]int{[]int{},[]int{1},[]int{2},[]int{1},[]int{3},[]int{2},[]int{2},[]int{2},[]int{2}}
  o := []string{"null","null","null","true","false","null","true","null","false"}

  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %v\n", run(i1, i2))
  fmt.Printf("Expect: %v\n", o)
}

func run(intrs []string, values [][]int) []string {
  var obj MyHashSet
  rtn := make([]string, len(intrs))

  for i, intr := range intrs {
    switch intr {
    case "MyHashSet":
      obj = Constructor()
      rtn[i] = "null"
    case "add":
      obj.Add(values[i][0])
      rtn[i] = "null"
    case "contains":
      rtn[i] = fmt.Sprintf("%t", obj.Contains(values[i][0]))
    case "remove":
      obj.Remove(values[i][0])
      rtn[i] = "null"
    default:
      fmt.Printf("Command Not Found\n")
      rtn[i] = "null"
    }
  }
  return rtn
}

// T: O(n)
// M: O(n)
// -- start --

type MyHashSet struct {
  m map[int]bool
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
  return MyHashSet{
    m: make(map[int]bool),
  }
}


func (this *MyHashSet) Add(key int)  {
  this.m[key] = true
  return
}


func (this *MyHashSet) Remove(key int)  {
  _, ok := this.m[key]
  if ok {
    delete(this.m, key)
  }
  return
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
  _, ok := this.m[key]
  return ok
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

// -- end --

