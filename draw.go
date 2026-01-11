package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

func dCAM() {

	dISOGRIDIMLISTIMCOLOR(FLOORS[currentFLOOR])

	if gridON {
		dISOGRIDLINES(FLOORS[currentFLOOR], CA(gameColorDARK, 30))
	}

	if DBG {
		dISOGRIDLINES(FLOORS[currentFLOOR], z.Magenta)
		dISOGRIDBORDER(FLOORS[currentFLOOR], 4, z.Orange)
	}

}

func dNOCAM() {

	if SCANLINES {
		dSCAN(3, 1, CA(z.Black, 150))
	}

	dMENUS()
	dINFO()

	if DBG {

		DEBUG()
		dIMSHEET(DICE, float32(debugRecW)+10, 10, 10, 1)

	}

	//MS CURSOR
	if selectTILE != BLANKNUM {
		dIM(TILES[selectTILE], R(MS.X, MS.Y, U1, U1), selectCOLOR)
	} else {
		dIM(ETCIM[2], R(MS.X, MS.Y, U1, U1), gameColor2)
	}

}
