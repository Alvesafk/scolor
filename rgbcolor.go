package scolor

import (
	"fmt"
)

type Color struct {
	Red, Green, Blue uint8
}

func FgRGB(s string, color Color) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", color.Red, color.Green, color.Blue, s)
}

func BgRGB(s string, color Color) string {
	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm%s\x1b[0m", color.Red, color.Green, color.Blue, s)
}
