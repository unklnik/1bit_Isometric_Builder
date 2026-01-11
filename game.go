package main

var (
	levW         = 28
	blokW        = U1
	selectTILE   = BLANKNUM
	currentFLOOR int
	gridON       = true
	selectCOLOR  = gameColor
)

// MARK: MAKE
func mLEVEXTRAS() {

}
func mLEV() {
	FLOORS = append(FLOORS, mISOGRIDCNT(levW, levW, blokW))
}
