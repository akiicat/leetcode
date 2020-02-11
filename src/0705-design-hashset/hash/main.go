package main

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

