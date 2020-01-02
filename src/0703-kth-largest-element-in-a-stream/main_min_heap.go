package main
import "fmt"
import "strconv"
import "container/heap"

func main() {
  i1 := []string{"KthLargest","add","add","add","add","add"}
  i2 := []interface{}{
    []interface{}{3,[]int{4,5,8,2}},
    []int{3},
    []int{5},
    []int{10},
    []int{9},
    []int{4},
  }
  o := []string{"null","4","5","5","8","8"}

  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %v\n", run(i1, i2))
  fmt.Printf("Expect: %v\n", o)

  i1 = []string{"KthLargest","add","add","add","add","add"}
  i2 = []interface{}{
    []interface{}{3,[]int{4}},
    []int{3},
    []int{5},
    []int{10},
    []int{9},
    []int{4},
  }
  o = []string{"null","-1","3","4","5","5"}

  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %v\n", run(i1, i2))
  fmt.Printf("Expect: %v\n", o)

  i1 = []string{"KthLargest","add","add","add","add","add"}
  i2 = []interface{}{
    []interface{}{1,[]int{}},
    []int{-3},
    []int{-2},
    []int{-4},
    []int{0},
    []int{4},
  }
  o = []string{"null","-3","-2","-2","0","4"}

  fmt.Printf("Input:  %v %v\n", i1, i2)
  fmt.Printf("Output: %v\n", run(i1, i2))
  fmt.Printf("Expect: %v\n", o)
}

func run(intrs []string, values []interface{}) []string {
  var obj KthLargest
  rtn := make([]string, len(intrs))

  for i, intr := range intrs {
    switch intr {
    case "KthLargest":
      v := values[i].([]interface{})
      obj = Constructor(v[0].(int), v[1].([]int))
      rtn[i] = "null"
    case "add":
      v := values[i].([]int)
      rtn[i] = strconv.Itoa(obj.Add(v[0]))
    default:
      fmt.Printf("Command Not Found\n")
      rtn[i] = "null"
    }
  }
  return rtn
}

// T: O(n * log(k))
// M: O(k)
// -- start --

type KthLargest struct {
  k int
  h Heap // min heap
}

func Constructor(k int, nums []int) KthLargest {
  h := Heap(nums)

  heap.Init(&h)

  for len(h) > k {
    heap.Pop(&h)
  }

  return KthLargest{k, h}
}


func (this *KthLargest) Add(val int) int {
  heap.Push(&this.h, val)

  for len(this.h) > this.k {
    heap.Pop(&this.h)
  }

  if len(this.h) < this.k {
    return -1
  }

  return this.h[0]
}

type Heap []int

func (h Heap) Len() int {
  return len(h)
}

func (h Heap) Less(i, j int) bool {
  return h[i] < h[j]
}

func (h Heap) Swap(i, j int) {
  h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push (v interface{}) {
  *h = append(*h, v.(int))
}

func (h Heap) Peak() int {
  return h[0]
}

func (h *Heap) Pop() interface{} {
  n := len(*h)
  v := (*h)[n-1]
  *h = (*h)[:n-1]
  return v
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

// -- end --

