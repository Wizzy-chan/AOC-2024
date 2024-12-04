package util

// There is no math.Abs for ints its kinda fucked up
func Abs(x int) int {
	if (x < 0) {
		return -x
	}
	return x
}
