package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	running                               bool
	screenWidth, screenHeight, frameCount int32

	// background \\
	bkgColour        rl.Color
	bkg1, bkg2, bkg3 Sprite

	// player \\
	player Player
)

type Player struct {
	*Sprite
	speed, dir, playerFrame int32
}

type Sprite struct {
	X, Y         int32
	sprite       rl.Texture2D
	dest, source rl.Rectangle
}

const ()

func input() {
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		player.dir--
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyLeft) {
		player.dir++
	}
}

func update() {
	frameCount++
	if frameCount >= 60 {
		frameCount = 0
	}
	running = !rl.WindowShouldClose()
	if frameCount%12 == 0 {
		player.playerFrame++
		if player.playerFrame >= 6 {
			player.playerFrame = 0
		}
		fmt.Println(player.playerFrame)
	}
	screenWidth = 1280
	screenHeight = 720
	player.dest.X += float32(player.speed) * float32(player.dir)
	player.dir = 0
	player.source.X = float32(56 * player.playerFrame)
}

func draw() {
	rl.BeginDrawing()

	// background
	rl.ClearBackground(bkgColour)
	rl.DrawTexturePro(bkg1.sprite, bkg1.source, bkg1.dest, rl.NewVector2(0, 0), 0, rl.RayWhite)
	rl.DrawTexturePro(bkg2.sprite, bkg2.source, bkg2.dest, rl.NewVector2(0, 0), 0, rl.RayWhite)
	rl.DrawTexturePro(bkg3.sprite, bkg3.source, bkg3.dest, rl.NewVector2(0, 0), 0, rl.RayWhite)

	// player
	rl.DrawTexturePro(player.sprite, player.source, player.dest, rl.NewVector2(0, 0), 0, rl.RayWhite)

	rl.DrawFPS(0, 0)
	rl.EndDrawing()

}

func init() {
	screenHeight = 720
	screenWidth = 1280
	bkgColour = rl.SkyBlue

	running = true

	rl.InitWindow(screenWidth, screenHeight, "I LOVE MY NEW KEYBOARD")
	rl.SetExitKey(rl.KeyF12)
	rl.SetTargetFPS(60)

	// background \\
	bkg1 = Sprite{
		0, 0,
		rl.LoadTexture("res/background/background_layer_1.png"),
		rl.NewRectangle(0, 0, float32(screenWidth), float32(screenHeight)),
		rl.NewRectangle(0, 0, 320, 180)}
	bkg2 = Sprite{
		0, 0,
		rl.LoadTexture("res/background/background_layer_2.png"),
		rl.NewRectangle(0, 0, float32(screenWidth), float32(screenHeight)),
		rl.NewRectangle(0, 0, 320, 180)}
	bkg3 = Sprite{
		0, 0,
		rl.LoadTexture("res/background/background_layer_3.png"),
		rl.NewRectangle(0, 0, float32(screenWidth), float32(screenHeight)),
		rl.NewRectangle(0, 0, 320, 180)}
	player = Player{
		&Sprite{
			0, 0,
			rl.LoadTexture("res/character/char_blue.png"),
			rl.NewRectangle(0, 0, 168, 168),
			rl.NewRectangle(0, 0, 56, 56)},
		10, 0, 0}
}

func quit() {
	rl.CloseWindow()
}

func main() {
	defer quit()
	for running {
		input()
		update()
		draw()
	}
}
