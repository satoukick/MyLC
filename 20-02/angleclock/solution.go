package main

import "fmt"

func angleClock(hour int, minutes int) float64 {
	anglePerMinute := 360 / 60
	hourAnglePerMinute := float64(360) / 12 / 60
	anglePerHour := 360 / 12

	hAng := float64((hour%12)*anglePerHour) + float64(minutes)*hourAnglePerMinute
	minAng := float64(minutes * anglePerMinute)
	result := hAng - minAng
	if result < 0 {
		result = -result
	}
	if result > 180 {
		result = 360 - result
	}
	return result
}

func main() {
	fmt.Println(angleClock(12, 30))
	fmt.Println(angleClock(3, 15))
	fmt.Println(angleClock(3, 30))
	fmt.Println(angleClock(4, 50))
	fmt.Println(angleClock(12, 0))
	fmt.Println(angleClock(1, 57))
}
