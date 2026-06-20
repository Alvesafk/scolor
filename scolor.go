/*
Author © 2026 alvesafk <migueldealmeidaalves55@gmail.com>

scolor makes it easy to colorize strings in go, it supports ansi and rgb colors.
*/
package scolor

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

// IsRGBSupported tells if the user terminal has support to True Color in order to use the
// RGB colors.
var IsRGBSupported bool

// Checks if user is on a TTY and has access to True Color, if both of them are right,
// IsRGBSupported is True.
func init() {
	IsRGBSupported = isTTY() && hasTrueColor()
}

// func isTTY receives nothing and returns a boolean, it uses syscalls to check if program
// is being run in a TTY.
func isTTY() bool {
	var termios syscall.Termios
	_, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		os.Stdout.Fd(),
		syscall.TCGETS,
		uintptr(unsafe.Pointer(&termios)),
	)

	return errno == 0
}

// func hasTrueColor receives nothing and returns a boolean, it checks if the env of the
// TTY has access to True Colors.
func hasTrueColor() bool {
	if !isTTY() {
		return false
	}

	colorterm := strings.ToLower(os.Getenv("COLORTERM"))
	return colorterm == "truecolor" || colorterm == "24bit"
}

// Color struct, it defines a 24bit RGB color, it has Red, Green and Blue uint8 fields.
type Color struct {
	Red, Green, Blue uint8
}

// func RGB receives a red, green and blue uint8 and returns a instantiated Color struct.
func RGB(red, green, blue uint8) Color {
	return Color{Red: red, Green: green, Blue: blue}
}

// func BgPrintln is a color method for printing text with a colored background on the
// terminal, the use is identical to the fmt Println function.
func (color Color) BgPrintln(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = BgRGB(fmt.Sprint(v), color)
	}

	return fmt.Fprintln(os.Stdout, colored...)
}

// func BgPrint is a color method for printing text with a colored background on the
// terminal, the use is identical to the fmt Print function.
func (color Color) BgPrint(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = BgRGB(fmt.Sprint(v), color)
	}

	return fmt.Fprint(os.Stdout, colored...)
}

// func BgPrintf is a color method for printing text with a colored background on the
// terminal, the use is identical to the fmt Printf function.
func (color Color) BgPrintf(format string, a ...any) (n int, err error) {
	return fmt.Fprintf(os.Stdout, BgRGB(format, color), a...)
}

// func FgPrintln is a color method for printing text with a colored foreground on the
// terminal, the use is identical to the fmt Println function.
func (color Color) FgPrintln(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = FgRGB(fmt.Sprint(v), color)
	}

	return fmt.Fprintln(os.Stdout, colored...)
}

// func FgPrint is a color method for printing text with a colored foreground on the
// terminal, the use is identical to the fmt Print function.
func (color Color) FgPrint(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = FgRGB(fmt.Sprint(v), color)
	}

	return fmt.Fprint(os.Stdout, colored...)
}

// func FgPrintf is a color method for printing text with a colored foreground on the
// terminal, the use is identical to the fmt Printf function.
func (color Color) FgPrintf(format string, a ...any) (n int, err error) {
	return fmt.Fprintf(os.Stdout, FgRGB(format, color), a...)
}

// func FgRGB receives a string and a color, it returns a string which it's foreground is
// colored.
func FgRGB(s string, color Color) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", color.Red, color.Green, color.Blue, s)
}

// func BgRGB receives a string and a color, it returns a string which it's background is
// colored.
func BgRGB(s string, color Color) string {
	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm%s\x1b[0m", color.Red, color.Green, color.Blue, s)
}

// Some preset colors, they all can be changed by the user of the lib.
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

// RgbTemplate accepts a background color and a foreground color, making it easier to print
// strings with colored background and foreground.
type RgbTemplate struct {
	Bg, Fg Color
}

// func CreateRgbTemplate receives a background color and a foreground color, it returns a
// initialized RgbTemplate struct.
func CreateRgbTemplate(bg, fg Color) *RgbTemplate {
	return &RgbTemplate{Bg: bg, Fg: fg}
}

// func Println is a RgbTemplate method for printing text with the background and foreground
// of the template onto the terminal, the use is identical to the fmt Println function.
func (template RgbTemplate) Println(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = BgRGB(FgRGB(fmt.Sprint(v), template.Fg), template.Bg)
	}

	return fmt.Fprintln(os.Stdout, colored...)
}

// func Print is a RgbTemplate method for printing text with the background and foreground
// of the template onto the terminal, the use is identical to the fmt Print function.
func (template RgbTemplate) Print(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = BgRGB(FgRGB(fmt.Sprint(v), template.Fg), template.Bg)
	}

	return fmt.Fprint(os.Stdout, colored...)
}

// func Printf is a RgbTemplate method for printing text with the background and foreground
// of the template onto the terminal, the use is identical to the fmt Printf function.
func (template RgbTemplate) Printf(format string, a ...any) (n int, err error) {
	return fmt.Fprintf(os.Stdout, BgRGB(FgRGB(format, template.Fg), template.Bg), a...)
}

// func FormatString is a RgbTemplate method, it receives a string and returns a formatted
// string with the colors of the template.
func (template RgbTemplate) FormatString(s string) string {
	return BgRGB(FgRGB(s, template.Fg), template.Bg)
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
type RgbTemplate struct 
func CreateRgbTemplate(bg, fg Color) *RgbTemplate
func (template RgbTemplate) Println(a ...any) (n int, err error)
func (template RgbTemplate) Print(a ...any) (n int, err error)
func (template RgbTemplate) Printf(format string, a ...any) (n int, err error)
func (template RgbTemplate) FormatString(s string) string
*/
