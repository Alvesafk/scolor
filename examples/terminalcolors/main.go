package main

import (
	"fmt"
	"time"

	"github.com/Alvesafk/scolor/ansi"
)

func main() {
	ansi.ABlue.FgPrintln("Hello, world! Using the terminal Blue!")
	time.Sleep(1 * time.Second / 2)

	ansi.ACyan.FgPrintln("Hello, world! Using the terminal Cyan!")
	time.Sleep(1 * time.Second / 2)

	ansi.AGreen.FgPrintln("Hello, world! Using the terminal Green!")
	time.Sleep(1 * time.Second / 2)

	ansi.APurple.FgPrintln("Hello, world! Using the terminal Purple!")
	time.Sleep(1 * time.Second / 2)

	ansi.ARed.FgPrintln("Hello, world! Using the terminal Red!")
	time.Sleep(1 * time.Second / 2)

	ansi.AYellow.FgPrintln("Hello, world! Using the terminal Yellow!")
	time.Sleep(1 * time.Second / 2)

	ansi.AWhite.FgPrintln("Hello, world! Using the terminal White!")
	time.Sleep(1 * time.Second / 2)

	fmt.Println(ansi.Rainbow("Hello, world! Rainbow string with the terminal colors!", "none", 0))
	time.Sleep(1 * time.Second / 2)
}
