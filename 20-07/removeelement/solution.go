package main

func removeElement(nums []int, val int) int {
	head, rear := 0, len(nums)-1
	for head <= rear {
		if nums[head] == val {
			nums[head], nums[rear] = nums[rear], nums[head]
			rear--
		} else {
			head++
		}
	}
	return rear
}
