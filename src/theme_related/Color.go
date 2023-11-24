package themerelated

import (
	"fmt"
	"strconv"
	"strings"
)

type Color struct {
	R, G, B uint8
}

func NewColor(r, g, b uint8) Color {
	return Color{R: r, G: g, B: b}
}

func (c Color) ToHex() string {
	return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
}

func FromHex(hex string) Color {
	if strings.HasPrefix(hex, "#") {
		hex = hex[1:]
	}
	hex = strings.ToLower(hex)
	r, _ := strconv.ParseUint(hex[0:2], 16, 8)
	g, _ := strconv.ParseUint(hex[2:4], 16, 8)
	b, _ := strconv.ParseUint(hex[4:6], 16, 8)

	return NewColor(uint8(r), uint8(g), uint8(b))
}
