package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	F1, F2   z.Font
	F1H, F2H float32
	F1SIZE   int32 = 18
	RUBIK    FONT
)

type FONT struct {
	fon     []z.Font
	sizes   []int32
	heights []float32
}

// MARK: DRAW
func dTXTFONT(t string, x, y float32, f FONT, size01234 int, c z.Color) float32 {
	z.DrawTextEx(f.fon[size01234], t, V2(x, y), float32(f.fon[size01234].BaseSize), 0, c)
	return cTXTFONTLEN(t, f, size01234)
}
func dTXT1XY(t string, x, y float32) {
	z.DrawTextEx(F1, t, V2(x, y), float32(F1.BaseSize), 0, z.White)
}
func dTXT1CNT(t string, cnt z.Vector2) {
	l := cTXT1LEN(t)
	x, y := cnt.X, cnt.Y
	x -= l / 2
	y -= F1H / 2
	dTXT1XY(t, x, y)
}
func dTXT1XYSIZECOLOR(t string, x, y, sizeMULTI float32, c z.Color) {
	z.DrawTextEx(F1, t, V2(x, y), float32(F1.BaseSize)*sizeMULTI, 0, c)
}

// MARK: CHECK
func cTXT1LEN(t string) float32 {
	v2 := z.MeasureTextEx(F1, t, float32(F1SIZE), 0)
	return v2.X
}
func cTXT1LENSIZE(t string, sizeMULTI float32) float32 {
	v2 := z.MeasureTextEx(F1, t, float32(F1SIZE)*sizeMULTI, 0)
	return v2.X
}
func cTXTFONTLEN(t string, f FONT, size01234 int) float32 {
	v2 := z.MeasureTextEx(f.fon[size01234], t, float32(f.sizes[size01234]), 0)
	return v2.X
}

// MARK: MAKE
func mTXT() {
	F1 = z.LoadFontEx("res/Rubik-Medium.ttf", F1SIZE, nil)
	F1H = F1.Recs.Height

	RUBIK = mFONT(RUBIK, "res/Rubik-Medium.ttf", []int32{14, 18, 24, 28, 32})

}

func mFONT(f FONT, path string, sizes []int32) FONT {
	f.sizes = sizes
	for i := range sizes {
		f.fon = append(f.fon, z.LoadFontEx(path, sizes[i], nil))
	}
	for i := range f.fon {
		f.heights = append(f.heights, f.fon[i].Recs.Height)
	}
	return f
}
