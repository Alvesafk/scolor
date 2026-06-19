package main

import (
	"fmt"
	"time"

	"github.com/Alvesafk/scolor/termc"
)

func main() {
	termc.ABlue.FgPrintln("Hello, world! Using the terminal Blue!")
	time.Sleep(1 * time.Second / 2)

	termc.ACyan.FgPrintln("Hello, world! Using the terminal Cyan!")
	time.Sleep(1 * time.Second / 2)

	termc.AGreen.FgPrintln("Hello, world! Using the terminal Green!")
	time.Sleep(1 * time.Second / 2)

	termc.APurple.FgPrintln("Hello, world! Using the terminal Purple!")
	time.Sleep(1 * time.Second / 2)

	termc.ARed.FgPrintln("Hello, world! Using the terminal Red!")
	time.Sleep(1 * time.Second / 2)

	termc.AYellow.FgPrintln("Hello, world! Using the terminal Yellow!")
	time.Sleep(1 * time.Second / 2)

	termc.AWhite.FgPrintln("Hello, world! Using the terminal White!")
	time.Sleep(1 * time.Second / 2)

	fmt.Println(termc.Rainbow("Hello, world! Rainbow string with the terminal colors!", "none", 0))
	time.Sleep(1 * time.Second / 2)
}
