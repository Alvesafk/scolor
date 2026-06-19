package scolor

import (
	"fmt"
	"os"
)

type Color struct {
	Red, Green, Blue uint8
}

// Bg Print
func (color Color) BgPrintln(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = BgRGB(fmt.Sprint(v), color)
	}

	return fmt.Fprintln(os.Stdout, colored...)
}

func (color Color) BgPrint(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = BgRGB(fmt.Sprint(v), color)
	}

	return fmt.Fprint(os.Stdout, colored...)
}

func (color Color) BgPrintf(format string, a ...any) (n int, err error) {
	return fmt.Fprintf(os.Stdout, BgRGB(format, color), a...)
}

// FG print
func (color Color) FgPrintln(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = FgRGB(fmt.Sprint(v), color)
	}

	return fmt.Fprintln(os.Stdout, colored...)
}

func (color Color) FgPrint(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = FgRGB(fmt.Sprint(v), color)
	}

	return fmt.Fprint(os.Stdout, colored...)
}

func (color Color) FgPrintf(format string, a ...any) (n int, err error) {
	return fmt.Fprintf(os.Stdout, FgRGB(format, color), a...)
}

var (
	RED    = Color{Red: 255, Green: 0, Blue: 0}
	GREEN  = Color{Red: 0, Green: 255, Blue: 0}
	BLUE   = Color{Red: 0, Green: 0, Blue: 255}
	YELLOW = Color{Red: 255, Green: 255, Blue: 0}
	ORANGE = Color{Red: 255, Green: 165, Blue: 0}
	PURPLE = Color{Red: 128, Green: 0, Blue: 128}
	PINK   = Color{Red: 255, Green: 192, Blue: 203}
	BROWN  = Color{Red: 164, Green: 42, Blue: 42}
	BLACK  = Color{Red: 0, Green: 0, Blue: 0}
	WHITE  = Color{Red: 255, Green: 255, Blue: 255}
	CYAN   = Color{Red: 0, Green: 255, Blue: 255}
)

func FgRGB(s string, color Color) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", color.Red, color.Green, color.Blue, s)
}

func BgRGB(s string, color Color) string {
	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm%s\x1b[0m", color.Red, color.Green, color.Blue, s)
}
