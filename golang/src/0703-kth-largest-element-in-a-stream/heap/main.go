package main
import "container/heap"

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

