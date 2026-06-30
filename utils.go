package scolor

import "regexp"

// Mod "enum", it's used on a switch statement for the AddMod function.
const (
	Bold = iota + 1
	Underline
	Strike
	Italic
)

// func AddMod receives a string to be modified and a mod, its recomended the use of the
// "enum" defined within this package, it returns the modified string, the mods are: Bold,
// Underline, Strike, Italic, if the mod string is different than this the function returns
// the string to be modified.
//
// Usage:
//
// boldString := AddMod("Hello, world!", Bold)
func AddMod(s string, mod int) string {
	switch mod {
	case Bold:
		s = "\033[1m" + s
	case Underline:
		s = "\033[4m" + s
	case Strike:
		s = "\033[9m" + s
	case Italic:
		s = "\033[3m" + s
	default:
		return s
	}

	return s
}

// func RemoveEscapeSequence receives a string and returns the same string without the escape
// sequences if any.
//
// Usage:
//
// boldString := AddMod("Hello, world!", Bold)
// cleanString := RemoveEscapeSequence(boldString)
func RemoveEscapeSequence(s string) string {
	escape := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return escape.ReplaceAllString(s, "")
}

/*
INDEX:
func AddMod(s string, mod int) string
func RemoveEscapeSequence(s string) string
*/
