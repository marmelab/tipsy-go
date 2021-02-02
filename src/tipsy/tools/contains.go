package tools

//ArrayContains return true if an array of [][2]int contains a [2]int, false otherwise.
func ArrayContains(obstacles [][2]int, position []int) bool {
	for _, v := range obstacles {
		if v[0] == position[0] && v[1] == position[1] {
			return true
		}
	}
	return false
}
