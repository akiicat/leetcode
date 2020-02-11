package main

// T: O(n)
// M: O(n)
// -- start --

type MyHashSet struct {
  arr []bool
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
  return MyHashSet{
    arr: make([]bool, 1000001),
  }
}


func (this *MyHashSet) Add(key int)  {
  this.arr[key] = true
  return
}


func (this *MyHashSet) Remove(key int)  {
  this.arr[key] = false
  return
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
  return this.arr[key]
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

// -- end --

