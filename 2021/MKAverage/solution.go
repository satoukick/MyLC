package main

import (
	"math"
	"sort"
)

type MKAverage struct {
	data    []int
	min     int
	removes int
}

func Constructor(m int, k int) MKAverage {
	return MKAverage{
		data:    make([]int, 0),
		min:     m,
		removes: k,
	}
}

func (this *MKAverage) AddElement(num int) {
	this.data = append(this.data, num)
}

func (this *MKAverage) CalculateMKAverage() int {
	if len(this.data) < this.min {
		return -1
	}

	// fmt.Println(this.data)

	cac := make([]int, this.min)
	copy(cac, this.data[len(this.data)-this.min:])

	sort.Ints(cac)

	// fmt.Println(cac)

	cac = cac[this.removes : len(cac)-this.removes]

	// fmt.Println(cac)

	sum := 0
	for _, v := range cac {
		sum += v

	}
	return int(math.Floor(float64(sum) / float64(len(cac))))
}

/**
 * Your MKAverage object will be instantiated and called as such:
 * obj := Constructor(m, k);
 * obj.AddElement(num);
 * param_2 := obj.CalculateMKAverage();
 */
