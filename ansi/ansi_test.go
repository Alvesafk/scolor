/*
Author © 2026 alvesafk <migueldealmeidaalves55@gmail.com>

Tests for ansi.go
*/
package ansi

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func captureAnsiStdout(t *testing.T, fn func() (int, error)) (string, int, error) {
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

func TestAnsiColorValues(t *testing.T) {
	want := []AnsiColor{Black, Red, Green, Yellow, Blue, Purple, Cyan, White}
	for i, got := range want {
		if got != AnsiColor(i) {
			t.Fatalf("color at index %d = %d, want %d", i, got, i)
		}
	}
}

func TestFgAnsiAndBgAnsi(t *testing.T) {
	if got, want := FgAnsi("hello", Cyan), "\033[36mhello\033[0m"; got != want {
		t.Fatalf("FgAnsi() = %q, want %q", got, want)
	}

	if got, want := BgAnsi("hello", Purple), "\033[45mhello\033[0m"; got != want {
		t.Fatalf("BgAnsi() = %q, want %q", got, want)
	}
}

func TestAnsiColorForegroundPrintMethods(t *testing.T) {
	color := Green

	t.Run("FgPrint", func(t *testing.T) {
		got, n, err := captureAnsiStdout(t, func() (int, error) {
			return color.FgPrint("hello", 123)
		})
		want := FgAnsi("hello", color) + FgAnsi("123", color)
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
		got, n, err := captureAnsiStdout(t, func() (int, error) {
			return color.FgPrintln("hello", "world")
		})
		want := FgAnsi("hello", color) + " " + FgAnsi("world", color) + "\n"
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
		got, n, err := captureAnsiStdout(t, func() (int, error) {
			return color.FgPrintf("hello %s", "gopher")
		})
		want := FgAnsi("hello gopher", color)
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

func TestAnsiColorBackgroundPrintMethods(t *testing.T) {
	color := Yellow

	t.Run("BgPrint", func(t *testing.T) {
		got, n, err := captureAnsiStdout(t, func() (int, error) {
			return color.BgPrint("hello", 123)
		})
		want := BgAnsi("hello", color) + BgAnsi("123", color)
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
		got, n, err := captureAnsiStdout(t, func() (int, error) {
			return color.BgPrintln("hello", "world")
		})
		want := BgAnsi("hello", color) + " " + BgAnsi("world", color) + "\n"
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
		got, n, err := captureAnsiStdout(t, func() (int, error) {
			return color.BgPrintf("hello %s", "gopher")
		})
		want := BgAnsi("hello gopher", color)
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

func TestCreateAnsiTemplateAndFormatString(t *testing.T) {
	template := CreateAnsiTemplate(White, Black)
	if template == nil {
		t.Fatal("CreateAnsiTemplate() returned nil")
	}
	if template.Bg != White || template.Fg != Black {
		t.Fatalf("CreateAnsiTemplate() = %#v, want Bg White and Fg Black", template)
	}

	want := BgAnsi(FgAnsi("hello", Black), White)
	if got := template.FormatString("hello"); got != want {
		t.Fatalf("AnsiTemplate.FormatString() = %q, want %q", got, want)
	}
}

func TestAnsiTemplatePrintMethods(t *testing.T) {
	template := AnsiTemplate{Bg: White, Fg: Black}

	t.Run("Print", func(t *testing.T) {
		got, n, err := captureAnsiStdout(t, func() (int, error) {
			return template.Print("hello", 123)
		})
		want := template.FormatString("hello") + template.FormatString("123")
		if err != nil {
			t.Fatalf("AnsiTemplate.Print() error = %v", err)
		}
		if got != want {
			t.Fatalf("AnsiTemplate.Print() output = %q, want %q", got, want)
		}
		if n != len(want) {
			t.Fatalf("AnsiTemplate.Print() n = %d, want %d", n, len(want))
		}
	})

	t.Run("Println", func(t *testing.T) {
		got, n, err := captureAnsiStdout(t, func() (int, error) {
			return template.Println("hello", "world")
		})
		want := template.FormatString("hello") + " " + template.FormatString("world") + "\n"
		if err != nil {
			t.Fatalf("AnsiTemplate.Println() error = %v", err)
		}
		if got != want {
			t.Fatalf("AnsiTemplate.Println() output = %q, want %q", got, want)
		}
		if n != len(want) {
			t.Fatalf("AnsiTemplate.Println() n = %d, want %d", n, len(want))
		}
	})

	t.Run("Printf", func(t *testing.T) {
		got, n, err := captureAnsiStdout(t, func() (int, error) {
			return template.Printf("hello %s", "gopher")
		})
		want := template.FormatString("hello gopher")
		if err != nil {
			t.Fatalf("AnsiTemplate.Printf() error = %v", err)
		}
		if got != want {
			t.Fatalf("AnsiTemplate.Printf() output = %q, want %q", got, want)
		}
		if n != len(want) {
			t.Fatalf("AnsiTemplate.Printf() n = %d, want %d", n, len(want))
		}
	})
}

func TestFgRainbow(t *testing.T) {
	got := FgRainbow("ab cdefgh")
	want := "\033[34ma\033[36mb \033[32mc\033[35md\033[31me\033[33mf\033[37mg\033[34mh"
	if got != want {
		t.Fatalf("FgRainbow() = %q, want %q", got, want)
	}
}

func TestBgRainbow(t *testing.T) {
	got := BgRainbow("ab cdefgh")
	want := FgAnsi("\033[44ma\033[46mb \033[42mc\033[45md\033[41me\033[43mf\033[47mg\033[44mh", Black)
	if got != want {
		t.Fatalf("BgRainbow() = %q, want %q", got, want)
	}
}
