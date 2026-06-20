/*
Author © 2026 alvesafk <migueldealmeidaalves55@gmail.com>

utils.go is for functions that work in generic form, meaning they work for both the scolor
main package (RGB) and the ansi package.
*/
package scolor

// func AddMod receives a string to be modified and a mod string, it returns the modified
// string, the mods are: "bold", "underline", "strike", "italic", if the mod string is
// different than this the function returns the string to be modified.
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

/*
INDEX:
func AddMod(s, mod string) string
*/
