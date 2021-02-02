package tools

func ArrayContains(obstacles [][]int, position []int) bool {
	for _, v := range obstacles {
		if v[0] == position[0] && v[1] == position[1] {
			return true
		}
	}
	return false
}
