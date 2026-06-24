package scolor

import "regexp"

// func AddMod receives a string to be modified and a mod string, it returns the modified
// string, the mods are: "bold", "underline", "strike", "italic", if the mod string is
// different than this the function returns the string to be modified.
//
// Usage:
//
// boldString := AddMod("Hello, world!", "bold")
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

// func RemoveEscapeSequence receives a string and returns the same string without the escape
// sequences if any.
//
// Usage:
//
// boldString := AddMod("Hello, world!", "bold")
// cleanString := RemoveEscapeSequence(boldString)
func RemoveEscapeSequence(s string) string {
	escape := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return escape.ReplaceAllString(s, "")
}

/*
INDEX:
func AddMod(s, mod string) string
func RemoveEscapeSequence(s string) string
*/
