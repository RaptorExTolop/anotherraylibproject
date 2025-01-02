package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	defer rl.CloseWindow()
	rl.InitWindow(1280, 720, "Just another day at the office")
	rl.SetTargetFPS(60)
	for rl.WindowShouldClose() {
		fmt.Println("running")
	}
	fmt.Println("quit")
}
