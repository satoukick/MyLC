package main

//TLE
func getTriggerTime(increase [][]int, requirements [][]int) []int {
	triggers := make([]int, len(requirements))
	for i := 0; i < len(triggers); i++ {
		triggers[i] = -1
	}
	status := []int{0, 0, 0}
	for i := 0; i < len(requirements); i++ {
		if requirements[i][0]+requirements[i][1]+requirements[i][2] == 0 {
			triggers[i] = 0
		}
	}
	for i := 0; i < len(increase); i++ {
		incr := increase[i]
		status[0], status[1], status[2] = status[0]+incr[0], status[1]+incr[1], status[2]+incr[2]
		for index := 0; index < len(requirements); index++ {
			if triggers[index] != -1 {
				continue
			}
			if status[0] >= requirements[index][0] &&
				status[1] >= requirements[index][1] &&
				status[2] >= requirements[index][2] {
				triggers[index] = i + 1
			}
		}
	}
	return triggers
}
