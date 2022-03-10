package main

// T: O(n)
// M: O(n)
// -- start --

type Node struct {
  Key int
  Val int
  Next *Node
}

type MyHashMap struct {
  arr []*Node
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
  return MyHashMap{
    arr: make([]*Node, 1<<10),
  }
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
  id := key & (1<<10 - 1)

  // find
  var prev *Node
  cur := this.arr[id]
  for cur != nil && cur.Key != key {
    prev, cur = cur, cur.Next
  }

  // update
  if cur != nil {
    cur.Val = value
    return
  }

  // create
  obj := &Node{Key: key, Val: value}
  if prev == nil {
    this.arr[id] = obj
  } else {
    prev.Next = obj
  }

  return
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
  id := key & (1<<10 - 1)

  cur := this.arr[id]
  for cur != nil {
    if cur.Key == key {
      return cur.Val
    }
    cur = cur.Next
  }

  return -1
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
  id := key & (1<<10 - 1)

  var prev *Node
  cur := this.arr[id]
  for cur != nil && cur.Key != key {
    prev, cur = cur, cur.Next
  }

  if cur == nil {
    return
  }

  if prev == nil {
    this.arr[id] = cur.Next
  } else {
    prev.Next = cur.Next
  }
  return
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

// -- end --

