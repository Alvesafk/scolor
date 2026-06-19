/*
Author © 2026 alvesafk <migueldealmeidaalves55@gmail.com>
*/

package scolor

import (
	"fmt"
	"os"
)

type Color struct {
	Red, Green, Blue uint8
}

func RGB(red, green, blue uint8) Color {
	return Color{Red: red, Green: green, Blue: blue}
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

func FgRGB(s string, color Color) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", color.Red, color.Green, color.Blue, s)
}

func BgRGB(s string, color Color) string {
	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm%s\x1b[0m", color.Red, color.Green, color.Blue, s)
}

var (
	BLACK  = Color{Red: 0, Green: 0, Blue: 0}
	BLUE   = Color{Red: 0, Green: 0, Blue: 200}
	BROWN  = Color{Red: 110, Green: 20, Blue: 20}
	CYAN   = Color{Red: 0, Green: 200, Blue: 200}
	GREEN  = Color{Red: 0, Green: 200, Blue: 0}
	ORANGE = Color{Red: 200, Green: 115, Blue: 0}
	PINK   = Color{Red: 200, Green: 140, Blue: 150}
	PURPLE = Color{Red: 80, Green: 0, Blue: 80}
	RED    = Color{Red: 200, Green: 0, Blue: 0}
	YELLOW = Color{Red: 200, Green: 200, Blue: 0}
	WHITE  = Color{Red: 255, Green: 255, Blue: 255}
)

type RgbTemplate struct {
	Bg, Fg Color
}

func CreateRgbTemplate(bg, fg Color) *RgbTemplate {
	return &RgbTemplate{Bg: bg, Fg: fg}
}

func (template RgbTemplate) Println(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = BgRGB(FgRGB(fmt.Sprint(v), template.Fg), template.Bg)
	}

	return fmt.Fprintln(os.Stdout, colored...)
}

func (template RgbTemplate) Print(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = BgRGB(FgRGB(fmt.Sprint(v), template.Fg), template.Bg)
	}

	return fmt.Fprint(os.Stdout, colored...)
}

func (template RgbTemplate) Printf(format string, a ...any) (n int, err error) {
	return fmt.Fprintf(os.Stdout, BgRGB(FgRGB(format, template.Fg), template.Bg), a...)
}

/*
INDEX:
type Color struct
func (color Color) BgPrintln(a ...any) (n int, err error)
func (color Color) BgPrint(a ...any) (n int, err error)
func (color Color) BgPrintf(format string, a ...any) (n int, err error)
func (color Color) FgPrintln(a ...any) (n int, err error)
func (color Color) FgPrint(a ...any) (n int, err error)
func (color Color) FgPrintf(format string, a ...any) (n int, err error)
func FgRGB(s string, color Color) string
func BgRGB(s string, color Color) string
var BLACK Color
var BLUE Color
var BROWN Color
var CYAN Color
var GREEN Color
var ORANGE Color
var PINK Color
var PURPLE Color
var RED Color
var YELLOW Color
var WHITE Color
*/
