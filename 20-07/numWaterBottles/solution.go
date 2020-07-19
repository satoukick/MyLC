package numWaterBottles

func numWaterBottles(numBottles int, numExchange int) int {
	sum := numBottles
	for numBottles >= numExchange {
		sum += numBottles / numExchange
		remain := numBottles % numExchange
		numBottles = numBottles/numExchange + remain
	}
	return sum
}
