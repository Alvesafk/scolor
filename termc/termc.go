/*
Author © 2026 alvesafk <migueldealmeidaalves55@gmail.com>
*/

package termc

import (
	"fmt"
	"os"
	"strings"

	"github.com/Alvesafk/scolor"
)

const (
	reset   = "\033[0m"
)

type ansiColor struct {
	bg, fg string
}

var (
	ARed    = ansiColor{bg: "\033[41m", fg: "\033[31m"}
	AGreen  = ansiColor{bg: "\033[42m", fg: "\033[32m"}
	AYellow = ansiColor{bg: "\033[43m", fg: "\033[33m"}
	ABlue   = ansiColor{bg: "\033[44m", fg: "\033[34m"}
	APurple = ansiColor{bg: "\033[45m", fg: "\033[35m"}
	ACyan   = ansiColor{bg: "\033[46m", fg: "\033[36m"}
	AWhite  = ansiColor{bg: "\033[47m", fg: "\033[37m"}
)

// bg
func (color ansiColor) BgPrintln(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = BgAnsi(fmt.Sprint(v), color)
	}

	return fmt.Fprintln(os.Stdout, colored...)
}

func (color ansiColor) BgPrint(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = BgAnsi(fmt.Sprint(v), color)
	}

	return fmt.Fprint(os.Stdout, colored...)
}

func (color ansiColor) BgPrintf(format string, a ...any) (n int, err error) {
	return fmt.Fprintf(os.Stdout, BgAnsi(format, color), a...)
}

// fg
func (color ansiColor) FgPrintln(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = FgAnsi(fmt.Sprint(v), color)
	}

	return fmt.Fprintln(os.Stdout, colored...)
}

func (color ansiColor) FgPrint(a ...any) (n int, err error) {
	colored := make([]any, len(a))
	for i, v := range a {
		colored[i] = FgAnsi(fmt.Sprint(v), color)
	}

	return fmt.Fprint(os.Stdout, colored...)
}

func (color ansiColor) FgPrintf(format string, a ...any) (n int, err error) {
	return fmt.Fprintf(os.Stdout, FgAnsi(format, color), a...)
}

func FgAnsi(s string, color ansiColor) string {
	return fmt.Sprintf("%s%s%s", color.fg, s, reset)
}

func BgAnsi(s string, color ansiColor) string {
	return fmt.Sprintf("%s%s%s", color.bg, s, reset)
}

func Rainbow(s, mod string, escape int) string {
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

/*
func AddMod(s, mod string) string
func Red(s string, mod string, escape int) string
func Green(s string, mod string, escape int) string
func Yellow(s string, mod string, escape int) string
func Blue(s string, mod string, escape int) string
func Purple(s string, mod string, escape int) string
func Cyan(s string, mod string, escape int) string
func White(s string, mod string, escape int) string
func Rainbow(s, mod string, escape int) string
*/
