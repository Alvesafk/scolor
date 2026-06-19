package main

import (
	"fmt"
	"time"

	"github.com/Alvesafk/scolor"
)

func main() {
	fmt.Println(scolor.TBlue("Hello, world! Using the terminal Blue!", "none", 0))
	time.Sleep(1 * time.Second / 2)

	fmt.Println(scolor.TCyan("Hello, world! Using the terminal Cyan!", "none", 0))
	time.Sleep(1 * time.Second / 2)

	fmt.Println(scolor.TGreen("Hello, world! Using the terminal Green!", "none", 0))
	time.Sleep(1 * time.Second / 2)

	fmt.Println(scolor.TPurple("Hello, world! Using the terminal Purple!", "none", 0))
	time.Sleep(1 * time.Second / 2)

	fmt.Println(scolor.TRed("Hello, world! Using the terminal Red!", "none", 0))
	time.Sleep(1 * time.Second / 2)

	fmt.Println(scolor.TYellow("Hello, world! Using the terminal Yellow!", "none", 0))
	time.Sleep(1 * time.Second / 2)

	fmt.Println(scolor.TWhite("Hello, world! Using the terminal White!", "none", 0))
	time.Sleep(1 * time.Second / 2)

	fmt.Println(scolor.TRainbow("Hello, world! Rainbow string with the terminal colors!", "none", 0))
	time.Sleep(1 * time.Second / 2)
}
