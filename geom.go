package main

import (
	"fmt"

	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	FLOORS []ISOGRID
)

type BLOK struct {
	r     z.Rectangle
	c, c2 z.Color
	nm    string
	cnt   z.Vector2

	solid, off, activ bool

	lit int
}

type ISOBLOK struct {
	cnt         z.Vector2
	tri         []TRI
	v2          []z.Vector2
	im          IM
	imLIST      []IM
	rD, rDAbove z.Rectangle
	zi          int
	numtype     int
	solid       bool
}
type TRI struct {
	v2 []z.Vector2
}
type ISOGRID struct {
	tri []TRI
	v2  []z.Vector2
	b   []ISOBLOK
}

// MARK:ISO DRAW
func dISOGRIDIMLISTIMCOLOR(ig ISOGRID) {
	dig := ZISORT(ig.b)
	for i := range dig {
		if len(dig[i].imLIST) > 0 {
			r := dig[i].rDAbove
			for j := range len(dig[i].imLIST) {
				dIM(dig[i].imLIST[j], r, dig[i].imLIST[j].c)
				r.Y -= r.Height / 2
			}
		}
	}
}
func dISOGRIDIMLIST(ig ISOGRID) {
	dig := ZISORT(ig.b)
	for i := range dig {
		if len(dig[i].imLIST) > 0 {
			r := dig[i].rDAbove
			for j := range len(dig[i].imLIST) {
				dIM(dig[i].imLIST[j], r, gameColor)
				r.Y -= r.Height / 2
			}
		}
	}
}
func dISOGRIDBORDER(ig ISOGRID, lineW float32, c z.Color) {
	z.DrawLineEx(ig.v2[0], ig.v2[1], lineW, c)
	z.DrawLineEx(ig.v2[1], ig.v2[2], lineW, c)
	z.DrawLineEx(ig.v2[2], ig.v2[3], lineW, c)
	z.DrawLineEx(ig.v2[3], ig.v2[0], lineW, c)
}
func dISOGRIDIM(ig []ISOBLOK) {
	dig := ZISORT(ig)

	for i := range dig {
		dIM(dig[i].im, dig[i].rD, z.White)
		if dig[i].numtype == 1 {
			//dIM(dig[i].imWall, dig[i].rDWall, z.White)
		} else if dig[i].numtype == 2 {
			//z.DrawCircleV(dig[i].cnt, 100, z.Orange)
			//dIM(ETCIM[0], R(100, 100, 32, 32), z.White)
		}

		if DBG {
			dTXT1XY(fmt.Sprint(dig[i].zi), dig[i].rD.X, dig[i].rD.Y)
		}
	}
}
func dISOGRIDRECS(ig ISOGRID, c z.Color) {
	for i := range len(ig.b) {
		dISOBLOKREC(ig.b[i], c)
	}
}
func dISOGRIDLINES(ig ISOGRID, c z.Color) {
	for i := range len(ig.b) {
		z.DrawLineV(ig.b[i].v2[0], ig.b[i].v2[1], c)
		z.DrawLineV(ig.b[i].v2[1], ig.b[i].v2[2], c)
		z.DrawLineV(ig.b[i].v2[2], ig.b[i].v2[3], c)
		z.DrawLineV(ig.b[i].v2[3], ig.b[i].v2[0], c)
		if DBG {
			dTXT1CNT(fmt.Sprint(i), ig.b[i].cnt)
		}
		if DMSBLOK && MSIN {
			if MSBLOKNUM == i {
				dISOBLOKREC(ig.b[i], CA(c, 100))
			}
		}
	}
}
func dISOBLOKREC(b ISOBLOK, c z.Color) {
	dTRI(b.tri[0], c)
	dTRI(b.tri[1], c)
}

// MARK:ISO MAKE
func mISOGRIDCNT(numW, numH int, wBlok float32) ISOGRID {
	var ig ISOGRID
	var ib []ISOBLOK
	h := float32(numH) * wBlok
	w := float32(numW) * wBlok
	x, y := CNT.X, CNT.Y
	y += h / 2
	ig.v2 = append(ig.v2, V2(x, y))
	ig.v2 = append(ig.v2, V2(x-w, y-h/2))
	ig.v2 = append(ig.v2, V2(x, y-h))
	ig.v2 = append(ig.v2, V2(x+w, y-h/2))
	ig.tri = m2XTRI4POINTS(ig.v2)
	ox, oy := x, y
	a := numW * numH
	c := 0
	zi := 0
	ozi := zi
	for a > 0 {
		b := ISOBLOK{}
		b.v2 = mISOBLOKPOINTS(V2(x, y), wBlok)
		b.cnt = mISOBLOKCNT(b)
		b.tri = mISOBLOKTRI(b)
		b.rD = R(b.v2[1].X, b.v2[2].Y, wBlok*2, wBlok*2)
		b.rDAbove = b.rD
		b.rDAbove.Y -= b.rDAbove.Height / 2
		b.zi = zi
		ib = append(ib, b)
		x -= wBlok
		y -= wBlok / 2
		a--
		zi++
		c++
		if c == numW {
			c = 0
			x = ox
			x += wBlok
			ox = x
			y = oy
			y -= wBlok / 2
			oy = y
			zi = ozi
			zi++
			ozi = zi
		}
	}
	ig.b = ib

	return ig
}

func mISOBLOKCNT(b ISOBLOK) z.Vector2 {
	return z.NewVector2(b.v2[0].X, b.v2[1].Y)
}
func mISOBLOKPOINTS(bV2 z.Vector2, w float32) []z.Vector2 {
	var v2 []z.Vector2
	v2 = append(v2, bV2)
	bV2.X -= w
	bV2.Y -= w / 2
	v2 = append(v2, bV2)
	bV2.X += w
	bV2.Y -= w / 2
	v2 = append(v2, bV2)
	bV2.X += w
	bV2.Y += w / 2
	v2 = append(v2, bV2)
	return v2
}

// MARK: LINES
func dLINEXY(x1, y1, x2, y2, lineW float32, c z.Color) {
	z.DrawLineEx(V2(x1, y1), V2(x2, y2), lineW, c)
}

// MARK: TRIANGLES
func dTRI(t TRI, c z.Color) {
	z.DrawTriangle(t.v2[1], t.v2[0], t.v2[2], c)
}
func mISOBLOKTRI(b ISOBLOK) []TRI {
	var tri []TRI
	t := TRI{}
	t.v2 = append(t.v2, b.v2[0], b.v2[1], b.v2[2])
	tri = append(tri, t)
	t = TRI{}
	t.v2 = append(t.v2, b.v2[0], b.v2[2], b.v2[3])
	tri = append(tri, t)
	return tri
}
func m2XTRI4POINTS(v2 []z.Vector2) []TRI {
	var tri []TRI
	t := TRI{}
	t.v2 = append(t.v2, v2[0], v2[1], v2[2])
	tri = append(tri, t)
	t = TRI{}
	t.v2 = append(t.v2, v2[0], v2[2], v2[3])
	tri = append(tri, t)
	return tri
}

// MARK:RECTANGLES
func R(x, y, w, h float32) z.Rectangle {
	return z.NewRectangle(x, y, w, h)
}
func dREC(r z.Rectangle, c z.Color) {
	z.DrawRectangleRec(r, c)

}
func dRECLINES(r z.Rectangle, lineW float32, c z.Color) {
	z.DrawRectangleLinesEx(r, lineW, c)
}
func dRECXY(x, y, w, h int32, c z.Color) {
	z.DrawRectangle(x, y, w, h, c)
}
func dRECXYLINES(x, y, w, h int32, c z.Color) {
	z.DrawRectangleLines(x, y, w, h, c)
}
func dRECCNT(cnt z.Vector2, w, h int32, c z.Color) {
	z.DrawRectangle(int32(cnt.X)-w/2, int32(cnt.Y)-h/2, w, h, c)
}
func dRECCNTLINES(cnt z.Vector2, w, h int32, c z.Color) {
	z.DrawRectangleLines(int32(cnt.X)-w/2, int32(cnt.Y)-h/2, w, h, c)
}
