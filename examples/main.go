package main

import (
	"fmt"
	"time"

	"github.com/Alvesafk/scolor"
)

func main() {
	fmt.Println(color.Blue("Hello, world! Using the terminal Blue!", "none", 0))
	time.Sleep(1 * time.Second)

	fmt.Println(color.Cyan("Hello, world! Using the terminal Cyan!", "none", 0))
	time.Sleep(1 * time.Second)

	fmt.Println(color.Green("Hello, world! Using the terminal Green!", "none", 0))
	time.Sleep(1 * time.Second)

	fmt.Println(color.Purple("Hello, world! Using the terminal Purple!", "none", 0))
	time.Sleep(1 * time.Second)

	fmt.Println(color.Red("Hello, world! Using the terminal Red!", "none", 0))
	time.Sleep(1 * time.Second)

	fmt.Println(color.Yellow("Hello, world! Using the terminal Yellow!", "none", 0))
	time.Sleep(1 * time.Second)

	fmt.Println(color.White("Hello, world! Using the terminal White!", "none", 0))
	time.Sleep(1 * time.Second)
}
