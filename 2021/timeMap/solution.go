package main

type TimeMap struct {
	m map[string][]*item
}

type item struct {
	value     string
	timestamp int
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][]*item),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if this.m[key] == nil {
		this.m[key] = make([]*item, 0)
	}
	item := &item{
		value:     value,
		timestamp: timestamp,
	}
	this.m[key] = append(this.m[key], item)
}

// Time Complexity: O(logn)
// Space Complexity: O(n)
func (this *TimeMap) Get(key string, timestamp int) string {
	if len(this.m[key]) == 0 {
		return ""
	}

	if this.m[key][0].timestamp > timestamp {
		return ""
	}

	left, right := 0, len(this.m[key])-1
	var result int

	for left <= right {
		mid := left + (right-left)/2
		if this.m[key][mid].timestamp == timestamp {
			result = mid
			break
		}
		if this.m[key][mid].timestamp > timestamp {
			right = mid - 1
		} else {
			left = mid + 1
			result = mid
		}
	}

	return this.m[key][result].value
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
