package tools

import (
	"strconv"
	"strings"
)

//GetPositionFromKey return a position from a key "x:y"
func GetPositionFromKey(key string) [2]int {
	positions := strings.Split(key, ":")
	x, xErr := strconv.Atoi(positions[0])
	y, yErr := strconv.Atoi(positions[1])
	if !(xErr == nil) || !(yErr == nil) {
		panic("Invalid key")
	}
	return [2]int{x, y}
}

//GetKeyFromPosition return a key from a position [2]int
func GetKeyFromPosition(position [2]int) string {
	return strconv.Itoa(position[0]) + ":" + strconv.Itoa(position[1])
}
