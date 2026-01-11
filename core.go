package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	SCRW, SCRH, FPS int32 = 1920, 1080, 60
	CAM             z.Camera2D
	CNT             z.Vector2
	FRAMES          int32

	//UNITS
	U1 float32 = 32

	UQTR, UHALF, U8TH, U3RD = U1 / 4, U1 / 2, U1 / 8, U1 / 3

	U2, U3, U4, U5, U6, U7, U8, U9, U10 = U1 * 2, U1 * 3, U1 * 4, U1 * 5, U1 * 6, U1 * 7, U1 * 8, U1 * 9, U1 * 10
)

func INIT() {
	z.SetTargetFPS(FPS)
	z.SetWindowState(z.FlagMsaa4xHint | z.FlagWindowUndecorated)
	z.HideCursor()

	CNT = V2(float32(SCRW/2), float32(SCRH/2))
	CAM.Target = CNT
	CAM.Offset = V2(float32(SCRW/2), float32(SCRH/2))
	CAM.Rotation = 0.0
	CAM.Zoom = 1

	gameColor = SEPIA()
	gameColorDARK = SEPIADARK()
	gameColor2 = z.Magenta
	gameColor3 = z.Orange
	selectCOLOR = gameColor

	mCOLORS()
	mTXT()
	mIMGAME()
	mLEV()
	mOBJ()
	mLEVEXTRAS()
	mMENUS()
	//mINFOBAR()
	//mPL()

	DMSBLOK = true
}
func RESTART() {
	FLOORS[currentFLOOR] = ISOGRID{}
	mLEV()
	mLEVEXTRAS()
}
func EXIT() {
	z.UnloadFont(F1)
	z.UnloadFont(F2)

	for i := range RUBIK.fon {
		z.UnloadFont(RUBIK.fon[i])
	}

	for i := range len(TEX) {
		z.UnloadTexture(TEX[i])
	}

}
