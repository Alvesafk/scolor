/*
Author © 2026 alvesafk <migueldealmeidaalves55@gmail.com>

ansi package is a separate package of scolor, the package uses the ansi colors of the user
terminal instead of the RGB colors of the main package.
*/
package ansi

import (
	"fmt"
	"os"
	"strings"
)

// const declaration for the reset ansi escape sequence.
const reset = "\033[0m"

// Custom int type for the Ansi Colors, they are just a enum from 0 to 7.
type AnsiColor int

// Main ansi colors.
const (
	Black AnsiColor = iota
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan
	White
)

// func BgPrintln is a AnsiColor method for printing text with a colored background on the
// terminal, the use is identical to the fmt Println function.
//
// Usage:
//
// Cyan.BgPrintln("Hello, ", "World!")
func (color AnsiColor) BgPrintln(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = BgAnsi(fmt.Sprint(v), color)
	}

	return fmt.Fprintln(os.Stdout, colored...)
}

// func BgPrint is a AnsiColor method for printing text with a colored background on the
// terminal, the use is identical to the fmt Print function.
//
// Usage:
//
// Cyan.BgPrintl"Hello, ", "World!")
func (color AnsiColor) BgPrint(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = BgAnsi(fmt.Sprint(v), color)
	}

	return fmt.Fprint(os.Stdout, colored...)
}

// func BgPrintf is a AnsiColor method for printing text with a colored background on the
// terminal, the use is identical to the fmt Printf function.
//
// Usage:
//
// Cyan.BgPrintf"Hello, ", "World!")
func (color AnsiColor) BgPrintf(format string, a ...any) (n int, err error) {
	colorString := BgAnsi(format, color)
	return fmt.Fprintf(os.Stdout, colorString, a...)
}

// func FgPrintln is a AnsiColor method for printing text with a colored foreground on the
// terminal, the use is identical to the fmt Println function.
//
// Usage:
//
// Cyan.FgPrintln"Hello, ", "World!")
func (color AnsiColor) FgPrintln(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = FgAnsi(fmt.Sprint(v), color)
	}

	return fmt.Fprintln(os.Stdout, colored...)
}

// func FgPrint is a AnsiColor method for printing text with a colored foreground on the
// terminal, the use is identical to the fmt Print function.
//
// Usage:
//
// Cyan.FgPrint("Hello, ", "World!")
func (color AnsiColor) FgPrint(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = FgAnsi(fmt.Sprint(v), color)
	}

	return fmt.Fprint(os.Stdout, colored...)
}

// func FgPrintf is a AnsiColor method for printing text with a colored foreground on the
// terminal, the use is identical to the fmt Printf function.
//
// Usage:
//
// Cyan.FgPrintf"Hello, ", "World!")
func (color AnsiColor) FgPrintf(format string, a ...any) (n int, err error) {
	colorString := FgAnsi(format, color)
	return fmt.Fprintf(os.Stdout, colorString, a...)
}

// func FgAnsi receives a string and a color, it returns a string which it's foreground is
// colored.
//
// Usage:
//
// cyanFgString := FgAnsi("Hello, World!", Cyan)
func FgAnsi(s string, color AnsiColor) string {
	return fmt.Sprintf("\033[3%vm%s%s", color, s, reset)
}

// func BgAnsi receives a string and a color, it returns a string which it's foreground is
// colored.
//
// Usage:
//
// cyanBgString := BgAnsi("Hello, World!", Cyan)
func BgAnsi(s string, color AnsiColor) string {
	return fmt.Sprintf("\033[4%vm%s%s", color, s, reset)
}

// AnsiTemplate accepts a background color and a foreground color, making it easier to print
// strings with colored background and foreground.
type AnsiTemplate struct {
	Bg, Fg AnsiColor
}

// func CreateAnsiTemplate receives a background color and a foreground color, it returns a
// initialized AnsiTemplate struct.
//
// Usage:
//
// whiteBgWithBlackFg := CreateAnsiTemplate(White, Black)
func CreateAnsiTemplate(bg, fg AnsiColor) *AnsiTemplate {
	return &AnsiTemplate{Bg: bg, Fg: fg}
}

// func Println is a AnsiTemplate method for printing text with the background and foreground
// of the template onto the terminal, the use is identical to the fmt Println function.
//
// Usage:
//
// whiteBgWithBlackFg.Println("Hello, ", "World!")
func (template AnsiTemplate) Println(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = BgAnsi(FgAnsi(fmt.Sprint(v), template.Fg), template.Bg)
	}

	return fmt.Fprintln(os.Stdout, colored...)
}

// func Print is a AnsiTemplate method for printing text with the background and foreground
// of the template onto the terminal, the use is identical to the fmt Print function.
//
// Usage:
//
// whiteBgWithBlackFg.Print("Hello, ", "World!")
func (template AnsiTemplate) Print(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = BgAnsi(FgAnsi(fmt.Sprint(v), template.Fg), template.Bg)
	}

	return fmt.Fprint(os.Stdout, colored...)
}

// func Printf is a AnsiTemplate method for printing text with the background and foreground
// of the template onto the terminal, the use is identical to the fmt Printf function.
//
// Usage:
//
// whiteBgWithBlackFg.Printf("Hello, ", "World!")
func (template AnsiTemplate) Printf(format string, a ...any) (n int, err error) {
	return fmt.Fprintf(os.Stdout, BgAnsi(FgAnsi(format, template.Fg), template.Bg), a...)
}

// func FormatString is a AnsiTemplate method, it receives a string and returns a formatted
// string with the colors of the template.
//
// Usage:
//
// stringWhiteBgBlackFg := whiteBgWithBlackFg.FormatString("Hello, world!")
func (template AnsiTemplate) FormatString(s string) string {
	return BgAnsi(FgAnsi(s, template.Fg), template.Bg)
}

// func FgRainbow receives a string, it returns a string whose foreground is like a rainbow
// based on the ansi colors.
//
// Usage:
//
// rainbowFgString := FgRainbow("Hello, world!")
func FgRainbow(s string) string {
	all_term_colors := []AnsiColor{Blue, Cyan, Green, Purple, Red, Yellow, White}

	var sb strings.Builder

	var cc int
	for _, c := range s {
		if cc >= 7 {
			cc = 0
		}

		if c == ' ' {
			sb.WriteString(" ")
			continue
		}

		sb.WriteString(fmt.Sprintf("\033[3%vm", all_term_colors[cc]))
		sb.WriteRune(c)

		cc++
	}

	return sb.String()
}

// func BgRainbow receives a string, it returns a string whose background is like a rainbow
// based on the ansi colors.
//
// Usage:
//
// rainbowFgString := FgRainbow("Hello, world!")
func BgRainbow(s string) string {
	all_term_colors := []AnsiColor{Blue, Cyan, Green, Purple, Red, Yellow, White}

	var sb strings.Builder

	var cc int
	for _, c := range s {
		if cc >= 7 {
			cc = 0
		}

		if c == ' ' {
			sb.WriteString(" ")
			continue
		}

		sb.WriteString(fmt.Sprintf("\033[4%vm", all_term_colors[cc]))
		sb.WriteRune(c)

		cc++
	}

	result := FgAnsi(sb.String(), Black)

	return result
}

/*
INDEX:
const reset
type AnsiColor int
const
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan
	White

func (color AnsiColor) BgPrintln(a ...any) (n int, err error)
func (color AnsiColor) BgPrint(a ...any) (n int, err error)
func (color AnsiColor) BgPrintf(a ...any) (n int, err error)
func (color AnsiColor) FgPrintln(a ...any) (n int, err error)
func (color AnsiColor) FgPrint(a ...any) (n int, err error)
func (color AnsiColor) FgPrintf(a ...any) (n int, err error)
func FgAnsi(s string, color AnsiColor) string
func BgAnsi(s string, color AnsiColor) string
type AnsiTemplate struct
func CreateAnsiTemplate(bg, fg AnsiColor) *AnsiTemplate
func (template AnsiTemplate) Println(a ...any) (n int, err error)
func (template AnsiTemplate) Print(a ...any) (n int, err error)
func (template AnsiTemplate) Printf(a ...any) (n int, err error)
func (template AnsiTemplate) FormatString(s string) string
func FgRainbow(s string) string
func BgRainbow(s string) string
*/
