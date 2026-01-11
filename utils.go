package main

import (
	"math"
	"math/rand/v2"
	"os"
	"path/filepath"
	"sort"
	"strings"

	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	BLANKNUM = 7777777777777777777
)

// MARK: ISO
func ZISORT(b []ISOBLOK) []ISOBLOK {
	sort.Slice(b, func(i, j int) bool { return b[i].zi > b[j].zi })
	return b
}
func cPOINTISOGRID(v2 z.Vector2, ig ISOGRID) int {
	num := MSBLOKNUM
	for i := range ig.b {
		if cPOINTISOBLOK(v2, ig.b[i]) {
			num = i
			break
		}
	}
	return num
}
func cPOINTISOBLOK(v2 z.Vector2, b ISOBLOK) bool {
	collides := false
	if cPOINTTRI(v2, b.tri[0]) || cPOINTTRI(v2, b.tri[1]) {
		collides = true
	}
	return collides
}

// MARK: TRI
func cPOINTTRI(v2 z.Vector2, t TRI) bool {
	return z.CheckCollisionPointTriangle(v2, t.v2[0], t.v2[1], t.v2[2])
}

// MARK: RECS
func RECCNT(r z.Rectangle) z.Vector2 {
	return z.NewVector2(r.X+r.Width/2, r.Y+r.Height/2)
}
func RECFROMCNT(cnt z.Vector2, w, h float32) z.Rectangle {
	return z.NewRectangle(cnt.X-w/2, cnt.Y-h/2, w, h)
}
func RECSMLR(r z.Rectangle, offset float32) z.Rectangle {
	return R(r.X+offset, r.Y+offset, r.Width-offset*2, r.Height-offset*2)
}
func RECLRGR(r z.Rectangle, offset float32) z.Rectangle {
	return R(r.X-offset, r.Y-offset, r.Width+offset*2, r.Height+offset*2)
}

// MARK:COLLIS
func cPOINTREC(v2 z.Vector2, r z.Rectangle) bool {
	return z.CheckCollisionPointRec(v2, r)
}
func cPOINTBLOKSCLICE(v2 z.Vector2, b []BLOK) int {
	num := 0
	for i := range b {
		if z.CheckCollisionPointRec(v2, b[i].r) {
			num = i
			break
		}
	}
	return num
}

func cRECBLOKSLICE(r z.Rectangle, b []BLOK) bool {
	collides := false
	for i := range b {
		if z.CheckCollisionRecs(r, b[i].r) {
			collides = true
			break
		}
	}
	return collides
}
func cRECRECSLICE(r z.Rectangle, rs []z.Rectangle) bool {
	collides := false
	for i := range rs {
		if z.CheckCollisionRecs(r, rs[i]) {
			collides = true
			break
		}
	}
	return collides
}
func cRECREC(r, r2 z.Rectangle) bool {
	return z.CheckCollisionRecs(r, r2)
}

func cBLOKRECINCIRC(b BLOK, cnt z.Vector2, rad float32) bool {
	collides := false
	if z.CheckCollisionCircleRec(cnt, rad, b.r) {
		collides = true
	}
	return collides
}

// MARK:VECTOR2
func V2(x, y float32) z.Vector2 {
	return z.NewVector2(x, y)
}
func NORMALIZE(v z.Vector2) z.Vector2 {
	length := float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
	if length == 0 {
		return z.NewVector2(0, 0)
	}
	return z.NewVector2(v.X/length, v.Y/length)
}

// MARK: FILES
func cPNG(path string) []string {
	var pngFiles []string

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".png") {
			pngFiles = append(pngFiles, path)
		}
		return nil
	})
	return pngFiles
}

// MARK: MATH
func ABSDIFF(a, b float32) float32 { return float32(math.Abs(float64(a - b))) }

func U1DIV(div float32) float32 {
	return U1 / div
}
func UMULTIDIV(div, multi float32) float32 {
	return (U1 / div) * multi
}

// MARK: RANDOM
func ROLL6() int {
	return rand.IntN(6) + 1
}
func FLIPCOIN() bool {
	return rand.IntN(2) == 0
}
func RINT(min, max int) int {
	return min + rand.IntN(max-min)
}
func RF32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
func RUINT8(min, max int) uint8 {
	return uint8(min) + uint8(rand.IntN(int(max-min+1)))
}
