/*
Author © 2026 alvesafk <migueldealmeidaalves55@gmail.com>

Tests for scolor.go
*/
package scolor

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func captureScolorStdout(t *testing.T, fn func() (int, error)) (string, int, error) {
	t.Helper()

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe() error = %v", err)
	}

	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	n, fnErr := fn()
	if closeErr := w.Close(); closeErr != nil {
		t.Fatalf("closing stdout pipe writer: %v", closeErr)
	}

	var buf bytes.Buffer
	if _, copyErr := io.Copy(&buf, r); copyErr != nil {
		t.Fatalf("reading captured stdout: %v", copyErr)
	}
	if closeErr := r.Close(); closeErr != nil {
		t.Fatalf("closing stdout pipe reader: %v", closeErr)
	}

	return buf.String(), n, fnErr
}

func TestRGB(t *testing.T) {
	t.Run("valid RGB values", func(t *testing.T) {
		got := RGB(12, 34, 56)
		want := Color{Red: 12, Green: 34, Blue: 56}
		if got != want {
			t.Fatalf("RGB(12, 34, 56) = %#v, want %#v", got, want)
		}
	})

	t.Run("rejects values greater than 255", func(t *testing.T) {
		cases := []struct {
			name             string
			red, green, blue int
		}{
			{name: "red", red: 256, green: 0, blue: 0},
			{name: "green", red: 0, green: 256, blue: 0},
			{name: "blue", red: 0, green: 0, blue: 256},
		}

		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				if got := RGB(tc.red, tc.green, tc.blue); got != (Color{}) {
					t.Fatalf("RGB(%d, %d, %d) = %#v, want zero Color", tc.red, tc.green, tc.blue, got)
				}
			})
		}
	})
}

func TestFgRGBAndBgRGB(t *testing.T) {
	color := Color{Red: 1, Green: 2, Blue: 3}

	if got, want := FgRGB("text", color), "\x1b[38;2;1;2;3mtext\x1b[0m"; got != want {
		t.Fatalf("FgRGB() = %q, want %q", got, want)
	}

	if got, want := BgRGB("text", color), "\x1b[48;2;1;2;3mtext\x1b[0m"; got != want {
		t.Fatalf("BgRGB() = %q, want %q", got, want)
	}
}

func TestColorForegroundPrintMethods(t *testing.T) {
	color := Color{Red: 9, Green: 8, Blue: 7}

	t.Run("FgPrint", func(t *testing.T) {
		got, n, err := captureScolorStdout(t, func() (int, error) {
			return color.FgPrint("hello", 123)
		})
		want := FgRGB("hello", color) + FgRGB("123", color)
		if err != nil {
			t.Fatalf("FgPrint() error = %v", err)
		}
		if got != want {
			t.Fatalf("FgPrint() output = %q, want %q", got, want)
		}
		if n != len(want) {
			t.Fatalf("FgPrint() n = %d, want %d", n, len(want))
		}
	})

	t.Run("FgPrintln", func(t *testing.T) {
		got, n, err := captureScolorStdout(t, func() (int, error) {
			return color.FgPrintln("hello", "world")
		})
		want := FgRGB("hello", color) + " " + FgRGB("world", color) + "\n"
		if err != nil {
			t.Fatalf("FgPrintln() error = %v", err)
		}
		if got != want {
			t.Fatalf("FgPrintln() output = %q, want %q", got, want)
		}
		if n != len(want) {
			t.Fatalf("FgPrintln() n = %d, want %d", n, len(want))
		}
	})

	t.Run("FgPrintf", func(t *testing.T) {
		got, n, err := captureScolorStdout(t, func() (int, error) {
			return color.FgPrintf("hello %s", "gopher")
		})
		want := FgRGB("hello gopher", color)
		if err != nil {
			t.Fatalf("FgPrintf() error = %v", err)
		}
		if got != want {
			t.Fatalf("FgPrintf() output = %q, want %q", got, want)
		}
		if n != len(want) {
			t.Fatalf("FgPrintf() n = %d, want %d", n, len(want))
		}
	})
}

func TestColorBackgroundPrintMethods(t *testing.T) {
	color := Color{Red: 7, Green: 8, Blue: 9}

	t.Run("BgPrint", func(t *testing.T) {
		got, n, err := captureScolorStdout(t, func() (int, error) {
			return color.BgPrint("hello", 123)
		})
		want := BgRGB("hello", color) + BgRGB("123", color)
		if err != nil {
			t.Fatalf("BgPrint() error = %v", err)
		}
		if got != want {
			t.Fatalf("BgPrint() output = %q, want %q", got, want)
		}
		if n != len(want) {
			t.Fatalf("BgPrint() n = %d, want %d", n, len(want))
		}
	})

	t.Run("BgPrintln", func(t *testing.T) {
		got, n, err := captureScolorStdout(t, func() (int, error) {
			return color.BgPrintln("hello", "world")
		})
		want := BgRGB("hello", color) + " " + BgRGB("world", color) + "\n"
		if err != nil {
			t.Fatalf("BgPrintln() error = %v", err)
		}
		if got != want {
			t.Fatalf("BgPrintln() output = %q, want %q", got, want)
		}
		if n != len(want) {
			t.Fatalf("BgPrintln() n = %d, want %d", n, len(want))
		}
	})

	t.Run("BgPrintf", func(t *testing.T) {
		got, n, err := captureScolorStdout(t, func() (int, error) {
			return color.BgPrintf("hello %s", "gopher")
		})
		want := BgRGB("hello gopher", color)
		if err != nil {
			t.Fatalf("BgPrintf() error = %v", err)
		}
		if got != want {
			t.Fatalf("BgPrintf() output = %q, want %q", got, want)
		}
		if n != len(want) {
			t.Fatalf("BgPrintf() n = %d, want %d", n, len(want))
		}
	})
}

func TestGradients(t *testing.T) {
	first := Color{Red: 0, Green: 0, Blue: 0}
	second := Color{Red: 30, Green: 60, Blue: 90}

	wantFg := FgRGB("a", Color{Red: 0, Green: 0, Blue: 0}) +
		FgRGB("b", Color{Red: 10, Green: 20, Blue: 30}) +
		FgRGB("c", Color{Red: 20, Green: 40, Blue: 60})
	if got := FgGradient("abc", first, second); got != wantFg {
		t.Fatalf("FgGradient() = %q, want %q", got, wantFg)
	}

	wantBg := BgRGB("a", Color{Red: 0, Green: 0, Blue: 0}) +
		BgRGB("b", Color{Red: 10, Green: 20, Blue: 30}) +
		BgRGB("c", Color{Red: 20, Green: 40, Blue: 60})
	if got := BgGradient("abc", first, second); got != wantBg {
		t.Fatalf("BgGradient() = %q, want %q", got, wantBg)
	}
}

func TestCheck8Bit(t *testing.T) {
	value := 10
	check8Bit(&value, 12)
	if value != 22 {
		t.Fatalf("check8Bit() value = %d, want 22", value)
	}

	check8Bit(&value, 300)
	if value != 255 {
		t.Fatalf("check8Bit() should clamp to 255, got %d", value)
	}
}

func TestRgbTemplate(t *testing.T) {
	bg := Color{Red: 1, Green: 2, Blue: 3}
	fg := Color{Red: 4, Green: 5, Blue: 6}

	template := CreateRgbTemplate(bg, fg)
	if template == nil {
		t.Fatal("CreateRgbTemplate() returned nil")
	}
	if template.Bg != bg || template.Fg != fg {
		t.Fatalf("CreateRgbTemplate() = %#v, want Bg %#v and Fg %#v", template, bg, fg)
	}

	want := BgRGB(FgRGB("hello", fg), bg)
	if got := template.FormatString("hello"); got != want {
		t.Fatalf("RgbTemplate.FormatString() = %q, want %q", got, want)
	}
}

func TestRgbTemplatePrintMethods(t *testing.T) {
	template := RgbTemplate{
		Bg: Color{Red: 1, Green: 2, Blue: 3},
		Fg: Color{Red: 4, Green: 5, Blue: 6},
	}

	t.Run("Print", func(t *testing.T) {
		got, n, err := captureScolorStdout(t, func() (int, error) {
			return template.Print("hello", 123)
		})
		want := template.FormatString("hello") + template.FormatString("123")
		if err != nil {
			t.Fatalf("RgbTemplate.Print() error = %v", err)
		}
		if got != want {
			t.Fatalf("RgbTemplate.Print() output = %q, want %q", got, want)
		}
		if n != len(want) {
			t.Fatalf("RgbTemplate.Print() n = %d, want %d", n, len(want))
		}
	})

	t.Run("Println", func(t *testing.T) {
		got, n, err := captureScolorStdout(t, func() (int, error) {
			return template.Println("hello", "world")
		})
		want := template.FormatString("hello") + " " + template.FormatString("world") + "\n"
		if err != nil {
			t.Fatalf("RgbTemplate.Println() error = %v", err)
		}
		if got != want {
			t.Fatalf("RgbTemplate.Println() output = %q, want %q", got, want)
		}
		if n != len(want) {
			t.Fatalf("RgbTemplate.Println() n = %d, want %d", n, len(want))
		}
	})

	t.Run("Printf", func(t *testing.T) {
		got, n, err := captureScolorStdout(t, func() (int, error) {
			return template.Printf("hello %s", "gopher")
		})
		want := template.FormatString("hello gopher")
		if err != nil {
			t.Fatalf("RgbTemplate.Printf() error = %v", err)
		}
		if got != want {
			t.Fatalf("RgbTemplate.Printf() output = %q, want %q", got, want)
		}
		if n != len(want) {
			t.Fatalf("RgbTemplate.Printf() n = %d, want %d", n, len(want))
		}
	})
}

func TestTmplGradientUsesTemplateColors(t *testing.T) {
	first := RgbTemplate{
		Bg: Color{Red: 10, Green: 20, Blue: 30},
		Fg: Color{Red: 40, Green: 50, Blue: 60},
	}
	second := RgbTemplate{
		Bg: Color{Red: 30, Green: 60, Blue: 90},
		Fg: Color{Red: 80, Green: 100, Blue: 120},
	}

	want := BgRGB(FgRGB("a", Color{Red: 40, Green: 50, Blue: 60}), Color{Red: 10, Green: 20, Blue: 30}) +
		BgRGB(FgRGB("b", Color{Red: 50, Green: 62, Blue: 75}), Color{Red: 15, Green: 30, Blue: 45}) +
		BgRGB(FgRGB("c", Color{Red: 60, Green: 74, Blue: 90}), Color{Red: 20, Green: 40, Blue: 60}) +
		BgRGB(FgRGB("d", Color{Red: 70, Green: 86, Blue: 105}), Color{Red: 25, Green: 50, Blue: 75})

	if got := TmplGradient("abcd", first, second); got != want {
		t.Fatalf("TmplGradient() = %q, want %q", got, want)
	}
}
