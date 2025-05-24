package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type colors struct {
	r uint32
	g uint32
	b uint32
	a uint32
}

const (
	C_BACKGROUND uint32 = 0xFF131112
	C_T1         uint32 = 0xFF31748F
	C_T2         uint32 = 0xFFF96900
	C_T3         uint32 = 0xFF08415C
	C_T4         uint32 = 0xFFA5BE00
	C_T5         uint32 = 0xFF3943B7
	C_T6         uint32 = 0xFF6B2737
	C_T7         uint32 = 0xFFB33F62
	C_T8         uint32 = 0xFFC4A7E7
	C_TEXT       uint32 = 0xFFE0DEF4
	C_BORDER     uint32 = 0xFFEB5160
	C_GRID       uint32 = 0xFFBDD4E7
)

func uint32ToRLColors(c uint32) rl.Color {
	col := rl.Color{}
	col.R = uint8((c >> 16) & 0xFF)
	col.G = uint8((c >> 8) & 0xFF)
	col.B = uint8(c & 0xFF)
	col.A = uint8((c >> 24) & 0xFF)

	return col
}
