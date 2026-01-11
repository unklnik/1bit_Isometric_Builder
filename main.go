package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	z.InitWindow(SCRW, SCRH, "2601")

	INIT()

	for !z.WindowShouldClose() {
		FRAMES++
		z.BeginDrawing()

		z.ClearBackground(z.Black)

		//DRAW CAM
		z.BeginMode2D(CAM)
		dCAM()
		z.EndMode2D()

		//END DRAW CAM

		//DRAW NOCAM
		dNOCAM()
		//ENDDRAW NOCAM
		z.EndDrawing()
		UP()
	}

	EXIT()

	z.CloseWindow()
}
