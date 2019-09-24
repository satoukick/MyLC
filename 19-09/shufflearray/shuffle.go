package main

import (
	"math/rand"
)

type Solution struct {
	cur  []int
	orig []int
}

func Constructor(nums []int) Solution {
	return Solution{
		cur:  make([]int, len(nums)),
		orig: nums,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	this.cur = make([]int, len(this.orig))
	copy(this.cur, this.orig)
	return this.orig
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	copy(this.cur, this.orig)
	for i := 0; i < len(this.cur); i++ {
		n := rand.Intn(len(this.cur))
		this.cur[i], this.cur[n] = this.cur[n], this.cur[i]
	}
	return this.cur
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
