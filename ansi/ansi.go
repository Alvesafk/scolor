/*
Author © 2026 alvesafk <migueldealmeidaalves55@gmail.com>

ansi package is a separte package of scolor, the package uses the ansi colors of the user
terminal instead of the RGB colors of the main package.
*/
package ansi

import (
	"fmt"
	"os"
	"strings"

	"github.com/Alvesafk/scolor"
)

// const declaration for the reset ansi escape sequence.
const (
	reset = "\033[0m"
)

// ansiColor struct is a private struct with the ansi codes for background and foreground
// colors.
type ansiColor struct {
	bg, fg string
}

// The default ansi terminal colors.
var (
	ABlack  = ansiColor{bg: "\033[40m", fg: "\033[30m"}
	ARed    = ansiColor{bg: "\033[41m", fg: "\033[31m"}
	AGreen  = ansiColor{bg: "\033[42m", fg: "\033[32m"}
	AYellow = ansiColor{bg: "\033[43m", fg: "\033[33m"}
	ABlue   = ansiColor{bg: "\033[44m", fg: "\033[34m"}
	APurple = ansiColor{bg: "\033[45m", fg: "\033[35m"}
	ACyan   = ansiColor{bg: "\033[46m", fg: "\033[36m"}
	AWhite  = ansiColor{bg: "\033[47m", fg: "\033[37m"}
)

// func BgPrintln is a ansiColor method for printing text with a colored background on the
// terminal, the use is identical to the fmt Println function.
func (color ansiColor) BgPrintln(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = BgAnsi(fmt.Sprint(v), color)
	}

	return fmt.Fprintln(os.Stdout, colored...)
}

// func BgPrint is a ansiColor method for printing text with a colored background on the
// terminal, the use is identical to the fmt Print function.
func (color ansiColor) BgPrint(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = BgAnsi(fmt.Sprint(v), color)
	}

	return fmt.Fprint(os.Stdout, colored...)
}

// func BgPrintf is a ansiColor method for printing text with a colored background on the
// terminal, the use is identical to the fmt Printf function.
func (color ansiColor) BgPrintf(format string, a ...any) (n int, err error) {
	return fmt.Fprintf(os.Stdout, BgAnsi(format, color), a...)
}

// func FgPrintln is a ansiColor method for printing text with a colored foreground on the
// terminal, the use is identical to the fmt Println function.
func (color ansiColor) FgPrintln(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = FgAnsi(fmt.Sprint(v), color)
	}

	return fmt.Fprintln(os.Stdout, colored...)
}

// func FgPrint is a ansiColor method for printing text with a colored foreground on the
// terminal, the use is identical to the fmt Print function.
func (color ansiColor) FgPrint(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = FgAnsi(fmt.Sprint(v), color)
	}

	return fmt.Fprint(os.Stdout, colored...)
}

// func FgPrintf is a ansiColor method for printing text with a colored foreground on the
// terminal, the use is identical to the fmt Printf function.
func (color ansiColor) FgPrintf(format string, a ...any) (n int, err error) {
	return fmt.Fprintf(os.Stdout, FgAnsi(format, color), a...)
}

// func FgAnsi receives a string and a color, it returns a string which it's foreground is
// colored.
func FgAnsi(s string, color ansiColor) string {
	return fmt.Sprintf("%s%s%s", color.fg, s, reset)
}

// func BgAnsi receives a string and a color, it returns a string which it's foreground is
// colored.
func BgAnsi(s string, color ansiColor) string {
	return fmt.Sprintf("%s%s%s", color.bg, s, reset)
}

// AnsiTemplate accepts a background color and a foreground color, making it easier to print
// strings with colored background and foreground.
type AnsiTemplate struct {
	Bg, Fg ansiColor
}

// func CreateAnsiTemplate receives a background color and a foreground color, it returns a
// initialized AnsiTemplate struct.
func CreateAnsiTemplate(bg, fg ansiColor) *AnsiTemplate {
	return &AnsiTemplate{Bg: bg, Fg: fg}
}

// func Println is a AnsiTemplate method for printing text with the background and foreground
// of the template onto the terminal, the use is identical to the fmt Println function.
func (template AnsiTemplate) Println(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = BgAnsi(FgAnsi(fmt.Sprint(v), template.Fg), template.Bg)
	}

	return fmt.Fprintln(os.Stdout, colored...)
}

// func Print is a AnsiTemplate method for printing text with the background and foreground
// of the template onto the terminal, the use is identical to the fmt Print function.
func (template AnsiTemplate) Print(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = BgAnsi(FgAnsi(fmt.Sprint(v), template.Fg), template.Bg)
	}

	return fmt.Fprint(os.Stdout, colored...)
}

// func Printf is a AnsiTemplate method for printing text with the background and foreground
// of the template onto the terminal, the use is identical to the fmt Printf function.
func (template AnsiTemplate) Printf(format string, a ...any) (n int, err error) {
	return fmt.Fprintf(os.Stdout, BgAnsi(FgAnsi(format, template.Fg), template.Bg), a...)
}

// func FormatString is a AnsiTemplate method, it receives a string and returns a formatted
// string with the colors of the template.
func (template AnsiTemplate) FormatString(s string) string {
	return BgAnsi(FgAnsi(s, template.Fg), template.Bg)
}

// func Rainbow receives a string to modify, a mod string and the amount of new lines, it
// returns a "rainbow" string based on the ansi colors.
func FgRainbow(s, mod string, escape int) string {
	all_term_colors := []string{ABlue.fg, ACyan.fg, AGreen.fg, APurple.fg, ARed.fg, AYellow.fg, AWhite.fg}

	var result string

	var cc int
	for _, c := range s {
		if cc >= 7 {
			cc = 0
		}

		if c == ' ' {
			result += " "
			continue
		}

		result += all_term_colors[cc] + string(c)

		cc++
	}

	result = scolor.AddMod(result, mod)

	if escape > 0 {
		result += strings.Repeat("\n", escape)
	}

	return result
}

func BgRainbow(s, mod string, escape int) string {
	all_term_colors := []string{ABlue.bg, ACyan.bg, AGreen.bg, APurple.bg, ARed.bg, AYellow.bg, AWhite.bg}

	var result string

	var cc int
	for _, c := range s {
		if cc >= 7 {
			cc = 0
		}

		result += all_term_colors[cc] + string(c)

		cc++
	}

	result = scolor.AddMod(result, mod)

	if escape > 0 {
		result += strings.Repeat("\n", escape)
	}

	result = FgAnsi(result, ABlack)

	return result
}

/*
INDEX:
const reset
type ansiColor struct
func (color ansiColor) BgPrintln(a ...any) (n int, err error)
func (color ansiColor) BgPrint(a ...any) (n int, err error)
func (color ansiColor) BgPrintf(a ...any) (n int, err error)
func (color ansiColor) FgPrintln(a ...any) (n int, err error)
func (color ansiColor) FgPrint(a ...any) (n int, err error)
func (color ansiColor) FgPrintf(a ...any) (n int, err error)
func FgAnsi(s string, color ansiColor) string
func BgAnsi(s string, color ansiColor) string
type AnsiTemplate struct
func CreateAnsiTemplate(bg, fg ansiColor) *AnsiTemplate
func (template AnsiTemplate) Println(a ...any) (n int, err error)
func (template AnsiTemplate) Print(a ...any) (n int, err error)
func (template AnsiTemplate) Printf(a ...any) (n int, err error)
func (template AnsiTemplate) FormatString(s string) string
func FgRainbow(s, mod string, escape int) string
func BgRainbow(s, mod string, escape int) string

var
	ARed
	AGreen
	AYellow
	ABlue
	APurple
	ACyan
	AWhite
*/
