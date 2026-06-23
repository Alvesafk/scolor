package main

import (
	"fmt"
	"time"

	"github.com/Alvesafk/scolor"
)

func main() {
	blue := scolor.BLUE
	cyan := scolor.CYAN
	green := scolor.GREEN
	purple := scolor.PURPLE
	red := scolor.RED
	yellow := scolor.YELLOW

	blue.FgPrintln("Hello, world! Using the RGB Blue!")
	blue.BgPrintln("Hello, world! Using the RGB Blue!")
	time.Sleep(1 * time.Second / 2)

	cyan.FgPrintln("Hello, world! Using the RGB Cyan!")
	cyan.BgPrintln("Hello, world! Using the RGB Cyan!")
	time.Sleep(1 * time.Second / 2)

	green.FgPrintln("Hello, world! Using the RGB Green!")
	green.BgPrintln("Hello, world! Using the RGB Green!")
	time.Sleep(1 * time.Second / 2)

	purple.FgPrintln("Hello, world! Using the RGB Purple!")
	purple.BgPrintln("Hello, world! Using the RGB Purple!")
	time.Sleep(1 * time.Second / 2)

	red.FgPrintln("Hello, world! Using the RGB Red!")
	red.BgPrintln("Hello, world! Using the RGB Red!")
	time.Sleep(1 * time.Second / 2)

	yellow.FgPrintln("Hello, world! Using the RGB Yellow!")
	yellow.BgPrintln("Hello, world! Using the RGB Yellow!")
	time.Sleep(1 * time.Second / 2)

	tmpl1 := scolor.CreateRgbTemplate(scolor.RGB(215, 106, 151), scolor.RGB(0, 0, 0))
	tmpl2 := scolor.CreateRgbTemplate(scolor.RGB(104, 150, 214), scolor.RGB(0, 0, 0))

	fmt.Println(scolor.TmplGradient("Any color you want!", *tmpl1, *tmpl2))
}
