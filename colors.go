package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	gameColor, gameColor2, gameColor3, gameColorDARK z.Color

	fadeONOFF, fadeliteONOFF bool

	fadeA, fadeALITE, fadeMIN, fadeMAX, fadeMINLITE, fadeMAXLITE, fadeSPD uint8 = 100, 50, 100, 200, 50, 150, 8

	COLORS []z.Color
)

// MARK: UTILS
func CA(c z.Color, a uint8) z.Color {
	c.A = a
	return c
}

// MARK: CUSTOM
func SEPIA() z.Color {
	return z.NewColor(246, 220, 107, 255)
}
func SEPIADARK() z.Color {
	return z.NewColor(236, 198, 15, 255)
}

// MARK: PALETTE
func MAROON() z.Color            { return z.NewColor(uint8(128), uint8(0), uint8(0), 255) }
func DARKRED() z.Color           { return z.NewColor(uint8(139), uint8(0), uint8(0), 255) }
func BROWN() z.Color             { return z.NewColor(uint8(165), uint8(42), uint8(42), 255) }
func FIREBRICK() z.Color         { return z.NewColor(uint8(178), uint8(34), uint8(34), 255) }
func CRIMSON() z.Color           { return z.NewColor(uint8(220), uint8(20), uint8(60), 255) }
func RED() z.Color               { return z.NewColor(uint8(255), uint8(0), uint8(0), 255) }
func TOMATO() z.Color            { return z.NewColor(uint8(255), uint8(99), uint8(71), 255) }
func CORAL() z.Color             { return z.NewColor(uint8(255), uint8(127), uint8(80), 255) }
func INDIANRED() z.Color         { return z.NewColor(uint8(205), uint8(92), uint8(92), 255) }
func LIGHTCORAL() z.Color        { return z.NewColor(uint8(240), uint8(128), uint8(128), 255) }
func DARKSALMON() z.Color        { return z.NewColor(uint8(233), uint8(150), uint8(122), 255) }
func SALMON() z.Color            { return z.NewColor(uint8(250), uint8(128), uint8(114), 255) }
func LIGHTSALMON() z.Color       { return z.NewColor(uint8(255), uint8(160), uint8(122), 255) }
func ORANGERED() z.Color         { return z.NewColor(uint8(255), uint8(69), uint8(0), 255) }
func DARKORANGE() z.Color        { return z.NewColor(uint8(255), uint8(140), uint8(0), 255) }
func ORANGE() z.Color            { return z.NewColor(uint8(255), uint8(165), uint8(0), 255) }
func GOLD() z.Color              { return z.NewColor(uint8(255), uint8(215), uint8(0), 255) }
func DARKGOLDENROD() z.Color     { return z.NewColor(uint8(184), uint8(134), uint8(11), 255) }
func GOLDENROD() z.Color         { return z.NewColor(uint8(218), uint8(165), uint8(32), 255) }
func PALEGOLDENROD() z.Color     { return z.NewColor(uint8(238), uint8(232), uint8(170), 255) }
func DARKKHAKI() z.Color         { return z.NewColor(uint8(189), uint8(183), uint8(107), 255) }
func KHAKI() z.Color             { return z.NewColor(uint8(240), uint8(230), uint8(140), 255) }
func OLIVE() z.Color             { return z.NewColor(uint8(128), uint8(128), uint8(0), 255) }
func YELLOW() z.Color            { return z.NewColor(uint8(255), uint8(255), uint8(0), 255) }
func YELLOWGREEN() z.Color       { return z.NewColor(uint8(154), uint8(205), uint8(50), 255) }
func DARKOLIVEGREEN() z.Color    { return z.NewColor(uint8(85), uint8(107), uint8(47), 255) }
func OLIVEDRAB() z.Color         { return z.NewColor(uint8(107), uint8(142), uint8(35), 255) }
func LAWNGREEN() z.Color         { return z.NewColor(uint8(124), uint8(252), uint8(0), 255) }
func CHARTREUSE() z.Color        { return z.NewColor(uint8(127), uint8(255), uint8(0), 255) }
func GREENYELLOW() z.Color       { return z.NewColor(uint8(173), uint8(255), uint8(47), 255) }
func DARKGREEN() z.Color         { return z.NewColor(uint8(0), uint8(100), uint8(0), 255) }
func GREEN() z.Color             { return z.NewColor(uint8(0), uint8(128), uint8(0), 255) }
func FORESTGREEN() z.Color       { return z.NewColor(uint8(34), uint8(139), uint8(34), 255) }
func LIME() z.Color              { return z.NewColor(uint8(0), uint8(255), uint8(0), 255) }
func LIMEGREEN() z.Color         { return z.NewColor(uint8(50), uint8(205), uint8(50), 255) }
func LIGHTGREEN() z.Color        { return z.NewColor(uint8(144), uint8(238), uint8(144), 255) }
func PALEGREEN() z.Color         { return z.NewColor(uint8(152), uint8(251), uint8(152), 255) }
func DARKSEAGREEN() z.Color      { return z.NewColor(uint8(143), uint8(188), uint8(143), 255) }
func MEDIUMSPRINGGREEN() z.Color { return z.NewColor(uint8(0), uint8(250), uint8(154), 255) }
func SPRINGGREEN() z.Color       { return z.NewColor(uint8(0), uint8(255), uint8(127), 255) }
func SEAGREEN() z.Color          { return z.NewColor(uint8(46), uint8(139), uint8(87), 255) }
func MEDIUMAQUAMARINE() z.Color  { return z.NewColor(uint8(102), uint8(205), uint8(170), 255) }
func MEDIUMSEAGREEN() z.Color    { return z.NewColor(uint8(60), uint8(179), uint8(113), 255) }
func LIGHTSEAGREEN() z.Color     { return z.NewColor(uint8(32), uint8(178), uint8(170), 255) }
func DARKSLATEGRAY() z.Color     { return z.NewColor(uint8(47), uint8(79), uint8(79), 255) }
func TEAL() z.Color              { return z.NewColor(uint8(0), uint8(128), uint8(128), 255) }
func DARKCYAN() z.Color          { return z.NewColor(uint8(0), uint8(139), uint8(139), 255) }
func AQUA() z.Color              { return z.NewColor(uint8(0), uint8(255), uint8(255), 255) }
func CYAN() z.Color              { return z.NewColor(uint8(0), uint8(255), uint8(255), 255) }
func LIGHTCYAN() z.Color         { return z.NewColor(uint8(224), uint8(255), uint8(255), 255) }
func DARKTURQUOISE() z.Color     { return z.NewColor(uint8(0), uint8(206), uint8(209), 255) }
func TURQUOISE() z.Color         { return z.NewColor(uint8(64), uint8(224), uint8(208), 255) }
func MEDIUMTURQUOISE() z.Color   { return z.NewColor(uint8(72), uint8(209), uint8(204), 255) }
func PALETURQUOISE() z.Color     { return z.NewColor(uint8(175), uint8(238), uint8(238), 255) }
func AQUAMARINE() z.Color        { return z.NewColor(uint8(127), uint8(255), uint8(212), 255) }
func POWDERBLUE() z.Color        { return z.NewColor(uint8(176), uint8(224), uint8(230), 255) }
func CADETBLUE() z.Color         { return z.NewColor(uint8(95), uint8(158), uint8(160), 255) }
func STEELBLUE() z.Color         { return z.NewColor(uint8(70), uint8(130), uint8(180), 255) }
func CORNFLOWERBLUE() z.Color    { return z.NewColor(uint8(100), uint8(149), uint8(237), 255) }
func DEEPSKYBLUE() z.Color       { return z.NewColor(uint8(0), uint8(191), uint8(255), 255) }
func DODGERBLUE() z.Color        { return z.NewColor(uint8(30), uint8(144), uint8(255), 255) }
func LIGHTBLUE() z.Color         { return z.NewColor(uint8(173), uint8(216), uint8(230), 255) }
func SKYBLUE() z.Color           { return z.NewColor(uint8(135), uint8(206), uint8(235), 255) }
func LIGHTSKYBLUE() z.Color      { return z.NewColor(uint8(135), uint8(206), uint8(250), 255) }
func MIDNIGHTBLUE() z.Color      { return z.NewColor(uint8(25), uint8(25), uint8(112), 255) }
func NAVY() z.Color              { return z.NewColor(uint8(0), uint8(0), uint8(128), 255) }
func DARKBLUE() z.Color          { return z.NewColor(uint8(0), uint8(0), uint8(139), 255) }
func MEDIUMBLUE() z.Color        { return z.NewColor(uint8(0), uint8(0), uint8(205), 255) }
func BLUE() z.Color              { return z.NewColor(uint8(0), uint8(0), uint8(255), 255) }
func ROYALBLUE() z.Color         { return z.NewColor(uint8(65), uint8(105), uint8(225), 255) }
func BLUEVIOLET() z.Color        { return z.NewColor(uint8(138), uint8(43), uint8(226), 255) }
func INDIGO() z.Color            { return z.NewColor(uint8(75), uint8(0), uint8(130), 255) }
func DARKSLATEBLUE() z.Color     { return z.NewColor(uint8(72), uint8(61), uint8(139), 255) }
func SLATEBLUE() z.Color         { return z.NewColor(uint8(106), uint8(90), uint8(205), 255) }
func MEDIUMSLATEBLUE() z.Color   { return z.NewColor(uint8(123), uint8(104), uint8(238), 255) }
func MEDIUMPURPLE() z.Color      { return z.NewColor(uint8(147), uint8(112), uint8(219), 255) }
func DARKMAGENTA() z.Color       { return z.NewColor(uint8(139), uint8(0), uint8(139), 255) }
func DARKVIOLET() z.Color        { return z.NewColor(uint8(148), uint8(0), uint8(211), 255) }
func DARKORCHID() z.Color        { return z.NewColor(uint8(153), uint8(50), uint8(204), 255) }
func MEDIUMORCHID() z.Color      { return z.NewColor(uint8(186), uint8(85), uint8(211), 255) }
func PURPLE() z.Color            { return z.NewColor(uint8(128), uint8(0), uint8(128), 255) }
func THISTLE() z.Color           { return z.NewColor(uint8(216), uint8(191), uint8(216), 255) }
func PLUM() z.Color              { return z.NewColor(uint8(221), uint8(160), uint8(221), 255) }
func VIOLET() z.Color            { return z.NewColor(uint8(238), uint8(130), uint8(238), 255) }
func MAGENTA() z.Color           { return z.NewColor(uint8(255), uint8(0), uint8(255), 255) }
func ORCHID() z.Color            { return z.NewColor(uint8(218), uint8(112), uint8(214), 255) }
func MEDIUMVIOLETRED() z.Color   { return z.NewColor(uint8(199), uint8(21), uint8(133), 255) }
func PALEVIOLETRED() z.Color     { return z.NewColor(uint8(219), uint8(112), uint8(147), 255) }
func DEEPPINK() z.Color          { return z.NewColor(uint8(255), uint8(20), uint8(147), 255) }
func HOTPINK() z.Color           { return z.NewColor(uint8(255), uint8(105), uint8(180), 255) }
func LIGHTPINK() z.Color         { return z.NewColor(uint8(255), uint8(182), uint8(193), 255) }
func PINK() z.Color              { return z.NewColor(uint8(255), uint8(192), uint8(203), 255) }
func ANTIQUEWHITE() z.Color      { return z.NewColor(uint8(250), uint8(235), uint8(215), 255) }
func BEIGE() z.Color             { return z.NewColor(uint8(245), uint8(245), uint8(220), 255) }
func BISQUE() z.Color            { return z.NewColor(uint8(255), uint8(228), uint8(196), 255) }
func BLANCHEDALMOND() z.Color    { return z.NewColor(uint8(255), uint8(235), uint8(205), 255) }
func WHEAT() z.Color             { return z.NewColor(uint8(245), uint8(222), uint8(179), 255) }
func CORNSILK() z.Color          { return z.NewColor(uint8(255), uint8(248), uint8(220), 255) }
func LEMONCHIFFON() z.Color      { return z.NewColor(uint8(255), uint8(250), uint8(205), 255) }
func LIGHTGOLDENRODYELLOW() z.Color {
	return z.NewColor(uint8(250), uint8(250), uint8(210), 255)
}
func LIGHTYELLOW() z.Color    { return z.NewColor(uint8(255), uint8(255), uint8(224), 255) }
func SADDLEBROWN() z.Color    { return z.NewColor(uint8(139), uint8(69), uint8(19), 255) }
func SIENNA() z.Color         { return z.NewColor(uint8(160), uint8(82), uint8(45), 255) }
func CHOCOLATE() z.Color      { return z.NewColor(uint8(210), uint8(105), uint8(30), 255) }
func PERU() z.Color           { return z.NewColor(uint8(205), uint8(133), uint8(63), 255) }
func SANDYBROWN() z.Color     { return z.NewColor(uint8(244), uint8(164), uint8(96), 255) }
func BURLYWOOD() z.Color      { return z.NewColor(uint8(222), uint8(184), uint8(135), 255) }
func TAN() z.Color            { return z.NewColor(uint8(210), uint8(180), uint8(140), 255) }
func ROSYBROWN() z.Color      { return z.NewColor(uint8(188), uint8(143), uint8(143), 255) }
func MOCCASIN() z.Color       { return z.NewColor(uint8(255), uint8(228), uint8(181), 255) }
func NAVAJOWHITE() z.Color    { return z.NewColor(uint8(255), uint8(222), uint8(173), 255) }
func PEACHPUFF() z.Color      { return z.NewColor(uint8(255), uint8(218), uint8(185), 255) }
func MISTYROSE() z.Color      { return z.NewColor(uint8(255), uint8(228), uint8(225), 255) }
func LAVENDERBLUSH() z.Color  { return z.NewColor(uint8(255), uint8(240), uint8(245), 255) }
func LINEN() z.Color          { return z.NewColor(uint8(250), uint8(240), uint8(230), 255) }
func OLDLACE() z.Color        { return z.NewColor(uint8(253), uint8(245), uint8(230), 255) }
func PAPAYAWHIP() z.Color     { return z.NewColor(uint8(255), uint8(239), uint8(213), 255) }
func SEASHELL() z.Color       { return z.NewColor(uint8(255), uint8(245), uint8(238), 255) }
func MINTCREAM() z.Color      { return z.NewColor(uint8(245), uint8(255), uint8(250), 255) }
func SLATEGRAY() z.Color      { return z.NewColor(uint8(112), uint8(128), uint8(144), 255) }
func LIGHTSLATEGRAY() z.Color { return z.NewColor(uint8(119), uint8(136), uint8(153), 255) }
func LIGHTSTEELBLUE() z.Color { return z.NewColor(uint8(176), uint8(196), uint8(222), 255) }
func LAVENDER() z.Color       { return z.NewColor(uint8(230), uint8(230), uint8(250), 255) }
func FLORALWHITE() z.Color    { return z.NewColor(uint8(255), uint8(250), uint8(240), 255) }
func ALICEBLUE() z.Color      { return z.NewColor(uint8(240), uint8(248), uint8(255), 255) }
func GHOSTWHITE() z.Color     { return z.NewColor(uint8(248), uint8(248), uint8(255), 255) }
func HONEYDEW() z.Color       { return z.NewColor(uint8(240), uint8(255), uint8(240), 255) }
func IVORY() z.Color          { return z.NewColor(uint8(255), uint8(255), uint8(240), 255) }
func AZURE() z.Color          { return z.NewColor(uint8(240), uint8(255), uint8(255), 255) }
func SNOW() z.Color           { return z.NewColor(uint8(255), uint8(250), uint8(250), 255) }
func BLACK() z.Color          { return z.NewColor(uint8(0), uint8(0), uint8(0), 255) }
func DIMGREY() z.Color        { return z.NewColor(uint8(105), uint8(105), uint8(105), 255) }
func GREY() z.Color           { return z.NewColor(uint8(128), uint8(128), uint8(128), 255) }
func DARKGREY() z.Color       { return z.NewColor(uint8(169), uint8(169), uint8(169), 255) }
func SILVER() z.Color         { return z.NewColor(uint8(192), uint8(192), uint8(192), 255) }
func LIGHTGREY() z.Color      { return z.NewColor(uint8(211), uint8(211), uint8(211), 255) }
func GAINSBORO() z.Color      { return z.NewColor(uint8(220), uint8(220), uint8(220), 255) }
func WHITESMOKE() z.Color     { return z.NewColor(uint8(245), uint8(245), uint8(245), 255) }
func WHITE() z.Color          { return z.NewColor(uint8(255), uint8(255), uint8(255), 255) }

// MARK: RANDOM COLORS
func RANCLR() z.Color {
	return z.NewColor(uint8(RINT(0, 256)), uint8(RINT(0, 256)), uint8(RINT(0, 256)), 255)
}
func RANDARKGREEN() z.Color {
	return z.NewColor(uint8(RINT(0, 30)), uint8(RINT(40, 90)), uint8(RINT(0, 40)), 255)
}
func RANGREEN() z.Color {
	return z.NewColor(uint8(RINT(0, 60)), uint8(RINT(140, 256)), uint8(RINT(0, 60)), 255)
}
func RANRED() z.Color {
	return z.NewColor(uint8(RINT(140, 256)), uint8(RINT(0, 60)), uint8(RINT(0, 60)), 255)
}
func RANPINK() z.Color {
	return z.NewColor(uint8(RINT(200, 256)), uint8(RINT(10, 110)), uint8(RINT(130, 180)), 255)
}
func RANBLUE() z.Color {
	return z.NewColor(uint8(RINT(0, 180)), uint8(RINT(0, 180)), uint8(RINT(140, 256)), 255)
}
func RANDARKBLUE() z.Color {
	return z.NewColor(uint8(RINT(0, 20)), uint8(RINT(0, 20)), uint8(RINT(100, 160)), 255)
}
func RANCYAN() z.Color {
	return z.NewColor(uint8(RINT(0, 120)), uint8(RINT(200, 256)), uint8(RINT(150, 256)), 255)
}
func RANYELLOW() z.Color {
	return z.NewColor(uint8(RINT(245, 256)), uint8(RINT(200, 256)), uint8(RINT(0, 100)), 255)
}
func RANORANGE() z.Color {
	return z.NewColor(uint8(255), uint8(RINT(70, 170)), uint8(RINT(0, 50)), 255)
}
func RANBROWN() z.Color {
	return z.NewColor(uint8(RINT(100, 150)), uint8(RINT(50, 120)), uint8(RINT(30, 90)), 255)
}
func RANGREY() z.Color {
	return z.NewColor(uint8(RINT(170, 220)), uint8(RINT(170, 220)), uint8(RINT(170, 220)), 255)
}
func RANDARKGREY() z.Color {
	return z.NewColor(uint8(RINT(90, 120)), uint8(RINT(90, 120)), uint8(RINT(90, 120)), 255)
}

// MARK: MAKE
func mCOLORS() {
	COLORS = append(COLORS, z.NewColor(uint8(0), uint8(0), uint8(0), 255),
		z.NewColor(uint8(0), uint8(0), uint8(128), 255),
		z.NewColor(uint8(0), uint8(0), uint8(139), 255),
		z.NewColor(uint8(0), uint8(0), uint8(205), 255),
		z.NewColor(uint8(0), uint8(0), uint8(255), 255),
		z.NewColor(uint8(0), uint8(100), uint8(0), 255),
		z.NewColor(uint8(0), uint8(128), uint8(0), 255),
		z.NewColor(uint8(0), uint8(128), uint8(128), 255),
		z.NewColor(uint8(0), uint8(139), uint8(139), 255),
		z.NewColor(uint8(0), uint8(191), uint8(255), 255),
		z.NewColor(uint8(0), uint8(206), uint8(209), 255),
		z.NewColor(uint8(0), uint8(250), uint8(154), 255),
		z.NewColor(uint8(0), uint8(255), uint8(0), 255),
		z.NewColor(uint8(0), uint8(255), uint8(127), 255),
		z.NewColor(uint8(0), uint8(255), uint8(255), 255),
		z.NewColor(uint8(0), uint8(255), uint8(255), 255),
		z.NewColor(uint8(100), uint8(149), uint8(237), 255),
		z.NewColor(uint8(102), uint8(205), uint8(170), 255),
		z.NewColor(uint8(105), uint8(105), uint8(105), 255),
		z.NewColor(uint8(106), uint8(90), uint8(205), 255),
		z.NewColor(uint8(107), uint8(142), uint8(35), 255),
		z.NewColor(uint8(112), uint8(128), uint8(144), 255),
		z.NewColor(uint8(119), uint8(136), uint8(153), 255),
		z.NewColor(uint8(123), uint8(104), uint8(238), 255),
		z.NewColor(uint8(124), uint8(252), uint8(0), 255),
		z.NewColor(uint8(127), uint8(255), uint8(0), 255),
		z.NewColor(uint8(127), uint8(255), uint8(212), 255),
		z.NewColor(uint8(128), uint8(0), uint8(0), 255),
		z.NewColor(uint8(128), uint8(0), uint8(128), 255),
		z.NewColor(uint8(128), uint8(128), uint8(0), 255),
		z.NewColor(uint8(128), uint8(128), uint8(128), 255),
		z.NewColor(uint8(135), uint8(206), uint8(235), 255),
		z.NewColor(uint8(135), uint8(206), uint8(250), 255),
		z.NewColor(uint8(138), uint8(43), uint8(226), 255),
		z.NewColor(uint8(139), uint8(0), uint8(0), 255),
		z.NewColor(uint8(139), uint8(0), uint8(139), 255),
		z.NewColor(uint8(139), uint8(69), uint8(19), 255),
		z.NewColor(uint8(143), uint8(188), uint8(143), 255),
		z.NewColor(uint8(144), uint8(238), uint8(144), 255),
		z.NewColor(uint8(147), uint8(112), uint8(219), 255),
		z.NewColor(uint8(148), uint8(0), uint8(211), 255),
		z.NewColor(uint8(152), uint8(251), uint8(152), 255),
		z.NewColor(uint8(153), uint8(50), uint8(204), 255),
		z.NewColor(uint8(154), uint8(205), uint8(50), 255),
		z.NewColor(uint8(160), uint8(82), uint8(45), 255),
		z.NewColor(uint8(165), uint8(42), uint8(42), 255),
		z.NewColor(uint8(169), uint8(169), uint8(169), 255),
		z.NewColor(uint8(173), uint8(216), uint8(230), 255),
		z.NewColor(uint8(173), uint8(255), uint8(47), 255),
		z.NewColor(uint8(175), uint8(238), uint8(238), 255),
		z.NewColor(uint8(176), uint8(196), uint8(222), 255),
		z.NewColor(uint8(176), uint8(224), uint8(230), 255),
		z.NewColor(uint8(178), uint8(34), uint8(34), 255),
		z.NewColor(uint8(184), uint8(134), uint8(11), 255),
		z.NewColor(uint8(186), uint8(85), uint8(211), 255),
		z.NewColor(uint8(188), uint8(143), uint8(143), 255),
		z.NewColor(uint8(189), uint8(183), uint8(107), 255),
		z.NewColor(uint8(192), uint8(192), uint8(192), 255),
		z.NewColor(uint8(199), uint8(21), uint8(133), 255),
		z.NewColor(uint8(205), uint8(133), uint8(63), 255),
		z.NewColor(uint8(205), uint8(92), uint8(92), 255),
		z.NewColor(uint8(210), uint8(105), uint8(30), 255),
		z.NewColor(uint8(210), uint8(180), uint8(140), 255),
		z.NewColor(uint8(211), uint8(211), uint8(211), 255),
		z.NewColor(uint8(216), uint8(191), uint8(216), 255),
		z.NewColor(uint8(218), uint8(112), uint8(214), 255),
		z.NewColor(uint8(218), uint8(165), uint8(32), 255),
		z.NewColor(uint8(219), uint8(112), uint8(147), 255),
		z.NewColor(uint8(220), uint8(20), uint8(60), 255),
		z.NewColor(uint8(220), uint8(220), uint8(220), 255),
		z.NewColor(uint8(221), uint8(160), uint8(221), 255),
		z.NewColor(uint8(222), uint8(184), uint8(135), 255),
		z.NewColor(uint8(224), uint8(255), uint8(255), 255),
		z.NewColor(uint8(230), uint8(230), uint8(250), 255),
		z.NewColor(uint8(233), uint8(150), uint8(122), 255),
		z.NewColor(uint8(238), uint8(130), uint8(238), 255),
		z.NewColor(uint8(238), uint8(232), uint8(170), 255),
		z.NewColor(uint8(240), uint8(128), uint8(128), 255),
		z.NewColor(uint8(240), uint8(230), uint8(140), 255),
		z.NewColor(uint8(240), uint8(248), uint8(255), 255),
		z.NewColor(uint8(240), uint8(255), uint8(240), 255),
		z.NewColor(uint8(240), uint8(255), uint8(255), 255),
		z.NewColor(uint8(244), uint8(164), uint8(96), 255),
		z.NewColor(uint8(245), uint8(222), uint8(179), 255),
		z.NewColor(uint8(245), uint8(245), uint8(220), 255),
		z.NewColor(uint8(245), uint8(245), uint8(245), 255),
		z.NewColor(uint8(245), uint8(255), uint8(250), 255),
		z.NewColor(uint8(248), uint8(248), uint8(255), 255),
		z.NewColor(uint8(25), uint8(25), uint8(112), 255),
		z.NewColor(uint8(250), uint8(128), uint8(114), 255),
		z.NewColor(uint8(250), uint8(235), uint8(215), 255),
		z.NewColor(uint8(250), uint8(240), uint8(230), 255),
		z.NewColor(uint8(250), uint8(250), uint8(210), 255),
		z.NewColor(uint8(253), uint8(245), uint8(230), 255),
		z.NewColor(uint8(255), uint8(0), uint8(0), 255),
		z.NewColor(uint8(255), uint8(0), uint8(255), 255),
		z.NewColor(uint8(255), uint8(105), uint8(180), 255),
		z.NewColor(uint8(255), uint8(127), uint8(80), 255),
		z.NewColor(uint8(255), uint8(140), uint8(0), 255),
		z.NewColor(uint8(255), uint8(160), uint8(122), 255),
		z.NewColor(uint8(255), uint8(165), uint8(0), 255),
		z.NewColor(uint8(255), uint8(182), uint8(193), 255),
		z.NewColor(uint8(255), uint8(192), uint8(203), 255),
		z.NewColor(uint8(255), uint8(20), uint8(147), 255),
		z.NewColor(uint8(255), uint8(215), uint8(0), 255),
		z.NewColor(uint8(255), uint8(218), uint8(185), 255),
		z.NewColor(uint8(255), uint8(222), uint8(173), 255),
		z.NewColor(uint8(255), uint8(228), uint8(181), 255),
		z.NewColor(uint8(255), uint8(228), uint8(196), 255),
		z.NewColor(uint8(255), uint8(228), uint8(225), 255),
		z.NewColor(uint8(255), uint8(235), uint8(205), 255),
		z.NewColor(uint8(255), uint8(239), uint8(213), 255),
		z.NewColor(uint8(255), uint8(240), uint8(245), 255),
		z.NewColor(uint8(255), uint8(245), uint8(238), 255),
		z.NewColor(uint8(255), uint8(248), uint8(220), 255),
		z.NewColor(uint8(255), uint8(250), uint8(205), 255),
		z.NewColor(uint8(255), uint8(250), uint8(240), 255),
		z.NewColor(uint8(255), uint8(250), uint8(250), 255),
		z.NewColor(uint8(255), uint8(255), uint8(0), 255),
		z.NewColor(uint8(255), uint8(255), uint8(224), 255),
		z.NewColor(uint8(255), uint8(255), uint8(240), 255),
		z.NewColor(uint8(255), uint8(255), uint8(255), 255),
		z.NewColor(uint8(255), uint8(69), uint8(0), 255),
		z.NewColor(uint8(255), uint8(99), uint8(71), 255),
		z.NewColor(uint8(30), uint8(144), uint8(255), 255),
		z.NewColor(uint8(32), uint8(178), uint8(170), 255),
		z.NewColor(uint8(34), uint8(139), uint8(34), 255),
		z.NewColor(uint8(46), uint8(139), uint8(87), 255),
		z.NewColor(uint8(47), uint8(79), uint8(79), 255),
		z.NewColor(uint8(50), uint8(205), uint8(50), 255),
		z.NewColor(uint8(60), uint8(179), uint8(113), 255),
		z.NewColor(uint8(64), uint8(224), uint8(208), 255),
		z.NewColor(uint8(65), uint8(105), uint8(225), 255),
		z.NewColor(uint8(70), uint8(130), uint8(180), 255),
		z.NewColor(uint8(72), uint8(209), uint8(204), 255),
		z.NewColor(uint8(72), uint8(61), uint8(139), 255),
		z.NewColor(uint8(75), uint8(0), uint8(130), 255),
		z.NewColor(uint8(85), uint8(107), uint8(47), 255),
		z.NewColor(uint8(95), uint8(158), uint8(160), 255),
	)

}
