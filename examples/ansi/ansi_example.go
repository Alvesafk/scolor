package main

import (
	"fmt"
	"time"

	"github.com/Alvesafk/scolor/ansi"
)

func main() {
	ansi.Blue.BgPrintln("Hello, world! Using the terminal Blue!")
	time.Sleep(1 * time.Second / 2)

	ansi.Cyan.FgPrintln("Hello, world! Using the terminal Cyan!")
	time.Sleep(1 * time.Second / 2)

	ansi.Green.BgPrintln("Hello, world! Using the terminal Green!")
	time.Sleep(1 * time.Second / 2)

	ansi.Purple.FgPrintln("Hello, world! Using the terminal Purple!")
	time.Sleep(1 * time.Second / 2)

	ansi.Red.BgPrintln("Hello, world! Using the terminal Red!")
	time.Sleep(1 * time.Second / 2)

	ansi.Yellow.FgPrintln("Hello, world! Using the terminal Yellow!")
	time.Sleep(1 * time.Second / 2)

	ansi.White.BgPrintln("Hello, world! Using the terminal White!")
	time.Sleep(1 * time.Second / 2)

	fmt.Println(ansi.FgRainbow("Hello, world! Rainbow string with the terminal colors!"))
	fmt.Println(ansi.BgRainbow("Hello, world! Rainbow string with the terminal colors!"))
	time.Sleep(1 * time.Second / 2)
}
