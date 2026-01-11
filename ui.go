package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	INF         []INFO
	MN          []MENU
	MNSPD       = U1 / 2
	MENUISOPEN  bool
	MENUOPENNUM int
)

type INFO struct {
	nm  string
	r   z.Rectangle
	spc float32
}
type MENU struct {
	r, rTAB    z.Rectangle
	open, msin bool
	c, cBG     z.Color
	nm         string
}

// MARK: DRAW
func dMENUS() {

	for i := range len(MN) {

		if MN[i].r.X < float32(SCRW) {
			dREC(MN[i].r, MN[i].cBG)
			dRECLINES(MN[i].r, 1, MN[i].c)
		}
		if cMSREC(MN[i].rTAB) {
			dREC(MN[i].rTAB, CA(gameColor3, fadeA))
			if MN[i].open {
				txlen := cTXTFONTLEN("close >>  ", RUBIK, 4)
				dTXTFONT("close >>  ", MN[i].rTAB.X-txlen, MN[i].rTAB.Y, RUBIK, 4, gameColor)
			} else {
				txlen := cTXTFONTLEN(MN[i].nm+" <<  ", RUBIK, 4)
				dTXTFONT(MN[i].nm+" <<  ", MN[i].rTAB.X-txlen, MN[i].rTAB.Y, RUBIK, 4, gameColor)
			}
		}
		//RTAB IMG
		if MENUISOPEN {
			dIM(DICE[MENUOPENNUM], MN[MENUOPENNUM].rTAB, gameColor)
		} else if !MENUISOPEN && MN[MENUOPENNUM].r.X >= float32(SCRW) {
			switch i {
			case 0: //TILES
				dIM(DICE[0], MN[i].rTAB, gameColor)
			case 1: //COLORS
				dIM(DICE[1], MN[i].rTAB, gameColor)
			case 2: //SETTINGS
				dIM(DICE[2], MN[i].rTAB, gameColor)
			}
		}
		//dRECLINES(MN[i].rTAB, 1, MN[i].c)
		if MN[i].r.X < float32(SCRW) {
			spc := U1 / 4
			rW := U1
			switch i {
			case 0: //TILES
				x := MN[i].r.X + spc
				y := MN[i].r.Y + spc
				ox := x
				for j := range len(TILES) {
					if x+rW >= float32(SCRW)-rW {
						x = ox
						y += rW + spc
					}
					r := R(x, y, rW, rW)
					if cMSREC(r) {
						dREC(r, CA(gameColor3, fadeALITE))
						if MSL {
							selectTILE = j
						}
					}
					dIM(TILES[j], r, gameColor)
					x += rW + spc
				}
			case 1: //COLORS
				x := MN[i].r.X + spc
				y := MN[i].r.Y + spc
				ox := x
				for j := range len(COLORS) {
					if x+rW >= float32(SCRW)-rW {
						x = ox
						y += rW + spc
					}
					r := R(x, y, rW, rW)
					r2 := RECLRGR(r, 4)
					if cMSREC(r) {
						dREC(r2, CA(gameColor3, fadeALITE))
						if MSL {
							selectCOLOR = COLORS[j]
						}
					}
					dREC(r, COLORS[j])
					x += rW + spc
				}
			case 2: //SETTINGS
				x := MN[i].r.X + spc
				ox := x
				y := MN[i].r.Y + spc
				dTXTFONT("bloom shader ", x, y, RUBIK, 3, gameColor)
				x = float32(SCRW) - U2
				SHADER1ON = ONOFFSWITCH(SHADER1ON, x, y, UMULTIDIV(20, 14), UMULTIDIV(20, 3), gameColorDARK)
				y += RUBIK.heights[3]
				x = ox
				dTXTFONT("scan lines ", x, y, RUBIK, 3, gameColor)
				x = float32(SCRW) - U2
				SCANLINES = ONOFFSWITCH(SCANLINES, x, y, UMULTIDIV(20, 14), UMULTIDIV(20, 3), gameColorDARK)
				y += RUBIK.heights[3]
				x = ox

			}
		}
	}

}
func dINFO() {
	for i := range len(INF) {
		dREC(INF[i].r, gameColor)
		dTXTFONT(INF[i].nm, INF[i].r.X+INF[i].r.Width+INF[i].spc, INF[i].r.Y, RUBIK, 2, gameColor)
	}
}

// MARK: UPDATE
func uMENUS() {
	MSINMENU = false
	MENUISOPEN = false
	for i := range len(MN) {
		if MN[i].open {
			MENUISOPEN = true
			MENUOPENNUM = i
		}
		if cMSREC(MN[i].rTAB) || cMSREC(MN[i].r) {
			MSINMENU = true
		}
		if MN[i].open && MN[i].r.X > float32(SCRW)-MN[i].r.Width {
			MN[i].r.X -= MNSPD
		} else if MN[i].open && MN[i].r.X < float32(SCRW)-MN[i].r.Width {
			MN[i].r.X = float32(SCRW) - MN[i].r.Width
		}
		if !MN[i].open && MN[i].r.X < float32(SCRW) {
			MN[i].r.X += MNSPD
		} else if !MN[i].open && MN[i].r.X > float32(SCRW) {
			MN[i].r.X = float32(SCRW)
		}
		MN[i].rTAB.X = MN[i].r.X - MN[i].rTAB.Width
		MN[i].rTAB.X += 2
		if cPOINTREC(MS, MN[i].rTAB) {
			if MSL {
				MN[i].open = !MN[i].open
			}
		}

	}
}

// MARK: MAKE
func mMENUS() {
	tabsize := U1
	menuW := U10
	mn := MENU{}
	mn.c = gameColor
	mn.cBG = CA(z.Black, 150)
	mn.r = z.NewRectangle(float32(SCRW), 0, menuW, float32(SCRH))
	mn.rTAB = z.NewRectangle(mn.r.X-tabsize, mn.r.Y+UHALF, tabsize, tabsize)
	mn.rTAB.X += 2
	mn.nm = "tiles"
	MN = append(MN, mn)

	mn.rTAB = z.NewRectangle(mn.r.X-tabsize, mn.r.Y+UHALF+tabsize, tabsize, tabsize)
	mn.rTAB.X += 2
	mn.nm = "colors"
	MN = append(MN, mn)

	mn.rTAB = z.NewRectangle(mn.r.X-tabsize, mn.r.Y+UHALF+tabsize*2, tabsize, tabsize)
	mn.rTAB.X += 2
	mn.nm = "settings"
	MN = append(MN, mn)
}

func mINFOBAR() {
	siz := UHALF
	spc := U1DIV(4)
	txlen := cTXTFONTLEN("color", RUBIK, 2)
	r := R(CNT.X-siz/2, spc, siz, siz)
	r.X -= txlen + siz + spc
	in := INFO{}
	in.r = r
	in.nm = "color"
	in.spc = spc
	INF = append(INF, in)
}

// MARK: BUTTONS
func ONOFFSWITCH(onoff bool, x, y, siz, offset float32, c z.Color) bool {
	r := R(x, y, siz, siz)
	r2 := R(x+offset, y+offset, siz-offset*2, siz-offset*2)
	r3 := R(x-2, y-2, siz+4, siz+4)
	c2 := CA(c, 100)
	dRECLINES(r, 2, c)
	if onoff {
		dREC(r2, c)
	} else {
		dREC(r2, c2)
	}
	if cMSREC(r) {
		dRECLINES(r3, 2, CA(c, fadeA))
		if MSL {
			onoff = !onoff
		}
	}
	return onoff
}
