package main

import (
	"fmt"

	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	TEX                []z.Texture2D
	TILES, ETCIM, DICE []IM
	BLANKIM            IM
)

type IM struct {
	tex int
	r   z.Rectangle
	c   z.Color
}
type ANIM struct {
	ims                       []IM
	frameT                    int32
	totalFrames, currentFrame int
}

// MARK: GAME SPECIFIC
func mIMGAME() {
	mETCIM("res/etc")
	TILES = mIMSHEETpath("res/isotiles")
	DICE = mIMSHEETXY("res/dice.png", 6, 1, 0, 0, 16, 16, 0, 0)
}

// MARK: ANIM
func dANIM(anm ANIM, r z.Rectangle) ANIM {
	dIM(anm.ims[anm.currentFrame], r, z.White)
	if FRAMES%anm.frameT == 0 {
		anm.currentFrame++
	}
	if anm.currentFrame == anm.totalFrames {
		anm.currentFrame = 0
	}
	return anm
}
func dANIMCOLOR(anm ANIM, r z.Rectangle, c z.Color) ANIM {
	dIM(anm.ims[anm.currentFrame], r, c)
	if FRAMES%anm.frameT == 0 {
		anm.currentFrame++
	}
	if anm.currentFrame == anm.totalFrames {
		anm.currentFrame = 0
	}
	return anm
}
func dANIMCOLORFLIP(anm ANIM, r z.Rectangle, c z.Color) ANIM {
	dIMFLIP(anm.ims[anm.currentFrame], r, c)
	if FRAMES%anm.frameT == 0 {
		anm.currentFrame++
	}
	if anm.currentFrame == anm.totalFrames {
		anm.currentFrame = 0
	}
	return anm
}
func mANIMIMSHEET(ims []IM, fps int32, numFrames int) ANIM {
	anm := ANIM{}
	anm.frameT = FPS / fps
	anm.totalFrames = numFrames
	for i := range numFrames {
		anm.ims = append(anm.ims, ims[i])
	}
	return anm
}

// MARK: DRAW
func dIM(im IM, r z.Rectangle, c z.Color) {
	z.DrawTexturePro(TEX[im.tex], im.r, r, z.Vector2Zero(), 0, c)
}
func dIMFLIP(im IM, r z.Rectangle, c z.Color) {
	im.r.Width = -im.r.Width
	z.DrawTexturePro(TEX[im.tex], im.r, r, z.Vector2Zero(), 0, c)
}
func dIMSHEET(ims []IM, x, y, space, zoom float32) {
	ox := x
	for i := range ims {
		z.DrawTexturePro(TEX[ims[i].tex], ims[i].r, R(x, y, ims[i].r.Width*zoom, ims[i].r.Height*zoom), z.Vector2Zero(), 0, z.White)
		dTXT1XY(fmt.Sprint(i), x, y+ims[i].r.Height*zoom+2)
		x += space + ims[i].r.Width*zoom
		if x+ims[i].r.Width*zoom >= float32(SCRW) {
			x = ox
			y += space + ims[i].r.Height*zoom + F1H + 4
		}
	}

}

// MARK: MAKE
func mIMSHEETpath(path string) []IM {
	var ims []IM
	images := cPNG(path)
	println(len(images))
	for i := range len(images) {
		ims = append(ims, mIMfile(images[i]))
	}
	return ims
}
func mETCIM(path string) {
	images := cPNG(path)
	println(len(images))
	for i := range len(images) {
		ETCIM = append(ETCIM, mIMfile(images[i]))
	}
}
func mIMfile(path string) IM {
	im := IM{}
	image := z.LoadImage(path)
	TEX = append(TEX, z.LoadTextureFromImage(image))
	w := float32(image.Width)
	h := float32(image.Height)
	z.UnloadImage(image)
	im.r = R(0, 0, w, h)
	im.tex = len(TEX) - 1
	im.c = gameColor
	return im
}
func mIM(path string, x, y, w, h float32) IM {
	im := IM{}
	image := z.LoadImage(path)
	TEX = append(TEX, z.LoadTextureFromImage(image))
	z.UnloadImage(image)
	im.r = R(x, y, w, h)
	im.tex = len(TEX) - 1
	return im
}
func MIMSHEETIM(im IM, numW, numH int, x, y, w, h, offsetX, offsetY float32) []IM {
	var ims []IM
	a := numW * numH
	c := 0
	ox := x
	for a > 0 {
		im2 := IM{}
		im2.r = R(x, y, w, h)
		im2.tex = im.tex
		ims = append(ims, im2)
		x += w + offsetX
		c++
		a--
		if c == numW {
			c = 0
			x = ox
			y += h + offsetY
		}
	}
	return ims
}
func mIMSHEETXY(path string, numW, numH int, x, y, w, h, offsetX, offsetY float32) []IM {
	var ims []IM
	a := numW * numH
	c := 0
	ox := x
	image := z.LoadImage(path)
	TEX = append(TEX, z.LoadTextureFromImage(image))
	z.UnloadImage(image)
	for a > 0 {
		im := IM{}
		im.r = R(x, y, w, h)
		im.tex = len(TEX) - 1
		ims = append(ims, im)
		x += w + offsetX
		c++
		a--
		if c == numW {
			c = 0
			x = ox
			y += h + offsetY
		}
	}

	return ims
}
