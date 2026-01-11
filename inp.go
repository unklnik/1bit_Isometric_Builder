package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

var (

	//MOUSE
	MS, MSCAM z.Vector2
	MSBLOKNUM int

	MSIN, DMSBLOK, MSINMENU, MSL, MSR, MSM bool
)

func INP() {

	//INP PLAYER
	/*	move := z.NewVector2(0, 0)
		if z.IsKeyDown(z.KeyRight) || z.IsKeyDown(z.KeyD) {
			move.X += 1
			pl.lr = false
		}
		if z.IsKeyDown(z.KeyLeft) || z.IsKeyDown(z.KeyA) {
			move.X -= 1
			pl.lr = true
		}
		if z.IsKeyDown(z.KeyUp) || z.IsKeyDown(z.KeyW) {
			move.Y -= 1
		}
		if z.IsKeyDown(z.KeyDown) || z.IsKeyDown(z.KeyS) {
			move.Y += 1
		}
		move = NORMALIZE(move)
		if move.X == 0 && move.Y == 0 {
			pl.moving = false
		} else {
			v2 := pl.cnt
			v2.X += move.X * pl.spd
			v2.Y += move.Y * pl.spd
			v2.Y += pl.r.Height / 2
			if cPLMOVE(v2) {

			}
		}
	*/

	//ZOOM
	if z.IsKeyPressed(z.KeyLeftBracket) {
		switch CAM.Zoom {
		case 0.25:
			CAM.Zoom = 2
		case 0.5:
			CAM.Zoom = 0.25
		case 1:
			CAM.Zoom = 0.5
		case 1.5:
			CAM.Zoom = 1
		case 2:
			CAM.Zoom = 1.5
		}
	} else if z.IsKeyPressed(z.KeyRightBracket) {
		switch CAM.Zoom {
		case 0.25:
			CAM.Zoom = 0.5
		case 0.5:
			CAM.Zoom = 1
		case 1:
			CAM.Zoom = 1.5
		case 1.5:
			CAM.Zoom = 2
		case 2:
			CAM.Zoom = 0.25
		}
	}

	//CORE
	if z.IsKeyPressed(z.KeyF1) {
		DBG = !DBG
	}
	if z.IsKeyPressed(z.KeyF2) {
		RESTART()
	}

	//MOUSE
	MS = z.GetMousePosition()
	MSCAM = z.GetScreenToWorld2D(MS, CAM)
	if cPOINTTRI(MSCAM, FLOORS[currentFLOOR].tri[0]) || cPOINTTRI(MSCAM, FLOORS[currentFLOOR].tri[1]) {
		MSIN = true
		MSBLOKNUM = cPOINTISOGRID(MSCAM, FLOORS[currentFLOOR])
	} else {
		MSIN = false
		MSBLOKNUM = -1
	}

	if z.IsMouseButtonPressed(z.MouseButtonLeft) {
		MSL = true
	} else {
		MSL = false
	}
	if z.IsMouseButtonPressed(z.MouseButtonRight) {
		MSR = true
	} else {
		MSR = false
	}
	if z.IsMouseButtonPressed(z.MouseButtonMiddle) {
		MSM = true
	} else {
		MSM = false
	}
	//PLACE BLOCK
	if MSIN && MSL && selectTILE != BLANKNUM {
		blok := TILES[selectTILE]
		blok.c = selectCOLOR
		FLOORS[currentFLOOR].b[MSBLOKNUM].imLIST = append(FLOORS[currentFLOOR].b[MSBLOKNUM].imLIST, blok)
	}
	if selectTILE != BLANKNUM && MSR {
		selectTILE = BLANKNUM
	}
	if selectTILE != BLANKNUM && MSM {
		FLOORS[currentFLOOR].b[MSBLOKNUM].imLIST = append(FLOORS[currentFLOOR].b[MSBLOKNUM].imLIST, ETCIM[3])
	}

}
func cMSREC(r z.Rectangle) bool {
	return cPOINTREC(MS, r)
}
