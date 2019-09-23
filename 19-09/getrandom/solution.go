package main

import (
	"math/rand"
)

/*
Design a data structure that supports all following operations in average O(1) time.

insert(val): Inserts an item val to the set if not already present.
remove(val): Removes an item val from the set if present.
getRandom: Returns a random element from current set of elements. Each element must have the same probability of being returned.
Example:

// Init an empty set.
RandomizedSet randomSet = new RandomizedSet();

// Inserts 1 to the set. Returns true as 1 was inserted successfully.``
randomSet.insert(1);

// Returns false as 2 does not exist in the set.
randomSet.remove(2);

// Inserts 2 to the set, returns true. Set now contains [1,2].
randomSet.insert(2);

// getRandom should return either 1 or 2 randomly.
randomSet.getRandom();

// Removes 1 from the set, returns true. Set now contains [2].
randomSet.remove(1);

// 2 was already in the set, so return false.
randomSet.insert(2);

// Since 2 is the only number in the set, getRandom always return 2.
randomSet.getRandom();
*/

type RandomizedSet struct {
	val []int
	pos map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		val: []int{},
		pos: make(map[int]int),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.pos[val]; ok {
		return false
	}
	r.val = append(r.val, val)
	r.pos[val] = len(r.val) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (r *RandomizedSet) Remove(val int) bool {
	p, ok := r.pos[val]
	if !ok {
		return false
	}
	r.val[p], r.val[len(r.val)-1] = r.val[len(r.val)-1], r.val[p]
	r.pos[r.val[p]] = p
	delete(r.pos, val)
	r.val = r.val[:len(r.val)-1]

	return true
}

/** Get a random element from the set. */
func (r *RandomizedSet) GetRandom() int {
	n := rand.Intn(len(r.val))
	return r.val[n]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	obj := Constructor()
	obj.Insert(0)
	obj.Insert(1)
	obj.Remove(0)
	obj.Insert(2)
	obj.Remove(1)
	obj.GetRandom()
}
