package main

import (
	"fmt"

	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	DBG       bool
	debugRecW int32 = 300
)

func DEBUG() {
	dRECXY(0, 0, debugRecW, SCRH, CA(z.Maroon, 50))
	var x, y float32 = 10, 10
	dTXT1XY("MS.X "+fmt.Sprintf("%.0f", MS.X)+" MS.Y "+fmt.Sprintf("%.0f", MS.Y), x, y)
	y += F1H
	dTXT1XY("CAM.Zoom "+fmt.Sprint(CAM.Zoom)+" fadeA "+fmt.Sprint(fadeA), x, y)
	y += F1H
	dTXT1XY("MSIN "+fmt.Sprint(MSIN)+" MSBLOKNUM "+fmt.Sprint(MSBLOKNUM), x, y)
	y += F1H
	dTXT1XY("MSINMENU "+fmt.Sprint(MSINMENU)+" fadeA "+fmt.Sprint(fadeA), x, y)
	y += F1H

}
