package main

func UP() {
	INP()
	uPL()
	uCAM()
	uOBJ()
	if len(MN) > 0 {
		uMENUS()
	}
	TIMERS()
}

// MARK: PLAYER
func uPL() {

}

// MARK: OBJ
func uOBJ() {

}

// MARK: BLOKS

// MARK: CAM
func uCAM() {

}

// MARK: TIMERS
func TIMERS() {
	if fadeliteONOFF {
		if fadeALITE >= fadeMINLITE {
			fadeALITE -= fadeSPD
		} else {
			fadeliteONOFF = false
		}
	} else {
		if fadeALITE <= fadeMAXLITE {
			fadeALITE += fadeSPD
		} else {
			fadeliteONOFF = true
		}
	}
	if fadeONOFF {
		if fadeA >= fadeMIN {
			fadeA -= fadeSPD
		} else {
			fadeONOFF = false
		}
	} else {
		if fadeA <= fadeMAX {
			fadeA += fadeSPD
		} else {
			fadeONOFF = true
		}
	}

}
