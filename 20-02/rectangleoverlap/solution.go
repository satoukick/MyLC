package rectangleoverlap

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// rec2 above rec1
	if rec2[1] >= rec1[3] {
		return false
	}
	// right
	if rec2[0] >= rec1[2] {
		return false
	}
	// below
	if rec2[3] <= rec1[1] {
		return false
	}
	// left
	if rec2[2] <= rec1[0] {
		return false
	}
	return true
}
