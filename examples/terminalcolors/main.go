package main

import (
	"fmt"
	"time"

	"github.com/Alvesafk/scolor/termc"
)

func main() {
	fmt.Println(termc.Blue("Hello, world! Using the terminal Blue!", "none", 0))
	time.Sleep(1 * time.Second / 2)

	fmt.Println(termc.Cyan("Hello, world! Using the terminal Cyan!", "none", 0))
	time.Sleep(1 * time.Second / 2)

	fmt.Println(termc.Green("Hello, world! Using the terminal Green!", "none", 0))
	time.Sleep(1 * time.Second / 2)

	fmt.Println(termc.Purple("Hello, world! Using the terminal Purple!", "none", 0))
	time.Sleep(1 * time.Second / 2)

	fmt.Println(termc.Red("Hello, world! Using the terminal Red!", "none", 0))
	time.Sleep(1 * time.Second / 2)

	fmt.Println(termc.Yellow("Hello, world! Using the terminal Yellow!", "none", 0))
	time.Sleep(1 * time.Second / 2)

	fmt.Println(termc.White("Hello, world! Using the terminal White!", "none", 0))
	time.Sleep(1 * time.Second / 2)

	fmt.Println(termc.Rainbow("Hello, world! Rainbow string with the terminal colors!", "none", 0))
	time.Sleep(1 * time.Second / 2)
}
