package main

import (
	"fmt"
	"time"
	"tipsy/game"
)

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())
	game := game.New()
	fmt.Println(game)
}
