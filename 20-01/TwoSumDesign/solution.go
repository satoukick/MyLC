package main

type TwoSum struct {
	num map[int]int
}

/** Initialize your data structure here. */
func Constructor() TwoSum {
	return TwoSum{
		num: make(map[int]int),
	}
}

/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int) {
	this.num[number]++
}

/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
	for v1 := range this.num {
		v2 := value - v1
		if v2 == v1 {
			if this.num[v1] > 1 {
				return true
			}
		} else if this.num[v2] >= 1 {
			return true
		}
	}
	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
