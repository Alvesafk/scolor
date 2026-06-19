/*
Author © 2026 alvesafk <migueldealmeidaalves55@gmail.com>
*/

package scolor

import (
	"fmt"
	"strings"
)

const (
	reset   = "\033[0m"
	cRed    = "\033[31m"
	cGreen  = "\033[32m"
	cYellow = "\033[33m"
	cBlue   = "\033[34m"
	cPurple = "\033[35m"
	cCyan   = "\033[36m"
	cWhite  = "\033[37m"
)

func AddMod(s, mod string) string {
	switch mod {
	case "bold":
		s = "\033[1m" + s
	case "underline":
		s = "\033[4m" + s
	case "strike":
		s = "\033[9m" + s
	case "italic":
		s = "\033[3m" + s
	default:
		return s
	}

	return s
}

func colorize(s, mod, code string, escape int) string {
	r := fmt.Sprintf("%s%s%s", code, AddMod(s, mod), reset)
	if escape > 0 {
		r += strings.Repeat("\n", escape)
	}

	return r
}

func TRed(s, mod string, escape int) string    { return colorize(s, mod, cRed, escape) }
func TGreen(s, mod string, escape int) string  { return colorize(s, mod, cGreen, escape) }
func TYellow(s, mod string, escape int) string { return colorize(s, mod, cYellow, escape) }
func TBlue(s, mod string, escape int) string   { return colorize(s, mod, cBlue, escape) }
func TPurple(s, mod string, escape int) string { return colorize(s, mod, cPurple, escape) }
func TCyan(s, mod string, escape int) string   { return colorize(s, mod, cCyan, escape) }
func TWhite(s, mod string, escape int) string  { return colorize(s, mod, cWhite, escape) }

func TRainbow(s, mod string, escape int) string {
	all_term_colors := []string{cBlue, cCyan, cGreen, cPurple, cRed, cYellow, cWhite}

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

	result = AddMod(result, mod)

	if escape > 0 {
		result += strings.Repeat("\n", escape)
	}

	return result
}

/*
func AddMod(s, mod string) string
func TRed(s string, mod string, escape int) string
func TGreen(s string, mod string, escape int) string
func TYellow(s string, mod string, escape int) string
func TBlue(s string, mod string, escape int) string
func TPurple(s string, mod string, escape int) string
func TCyan(s string, mod string, escape int) string
func TWhite(s string, mod string, escape int) string
func TRainbow(s, mod string, escape int) string
*/
