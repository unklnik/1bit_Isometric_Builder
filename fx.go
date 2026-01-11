package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	SHADER1ON, SCANLINES bool
)

func dSCAN(spc, lineW float32, c z.Color) {
	var x, y float32
	x2 := x + float32(SCRW)
	for y < float32(SCRH) {
		dLINEXY(x, y, x2, y, lineW, c)
		y += spc
	}
}
