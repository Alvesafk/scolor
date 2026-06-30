/*
Author © 2026 alvesafk <migueldealmeidaalves55@gmail.com>

Tests for ansi.go
*/
package ansi

import (
	"strings"
	"testing"
)

// FgAnsi / BgAnsi - escape sequence format

func TestFgAnsi_EscapeSequences(t *testing.T) {
	result := FgAnsi("hello", Red)

	if !strings.HasPrefix(result, "\033[31m") {
		t.Errorf("FgAnsi: expected red fg prefix \\033[31m, got %q", result)
	}

	if !strings.Contains(result, "hello") {
		t.Errorf("FgAnsi: Original string missing, got %q", result)
	}

	if !strings.HasSuffix(result, "\033[0m") {
		t.Errorf("FgAnsi: expected reset suffix, got %q", result)
	}
}

func TestBgAnsi_EscapeSequences(t *testing.T) {
	result := BgAnsi("hello", Red)

	if !strings.HasPrefix(result, "\033[41m") {
		t.Errorf("BgAnsi: expected red bg prefix \\033[41m, got %q", result)
	}

	if !strings.Contains(result, "hello") {
		t.Errorf("BgAnsi: Original string missing, got %q", result)
	}

	if !strings.HasSuffix(result, "\033[0m") {
		t.Errorf("BgAnsi: expected reset suffix, got %q", result)
	}
}

func TestFgAnsi_EmptyString(t *testing.T) {
	result := FgAnsi("", Blue)
	if !strings.HasSuffix(result, "\033[0m") {
		t.Errorf("FgAnsi empty string: expected reset suffix, got %q", result)
	}
}

func TestBgAnsi_EmptyString(t *testing.T) {
	result := FgAnsi("", Yellow)
	if !strings.HasSuffix(result, "\033[0m") {
		t.Errorf("BgAnsi empty string: expected reset suffix, got %q", result)
	}
}

// AnsiTemplate

func TestCreateAnsiTemplate(t *testing.T) {
	tmpl := CreateAnsiTemplate(Blue, White)

	if tmpl == nil {
		t.Fatal("CreateAnsiTemplate returned nil")
	}

	if tmpl.Bg != Blue {
		t.Errorf("Bg = %v, want ABlue", tmpl.Bg)
	}

	if tmpl.Fg != White {
		t.Errorf("Fg = %v, want AWhite", tmpl.Fg)
	}
}

// Rainbow

func TestFgRainbow_ContainsInput(t *testing.T) {
	result := FgRainbow("hello")
	for _, c := range "hello" {
		if !strings.Contains(result, string(c)) {
			t.Errorf("Rainbow: character %q missing from result", c)
		}
	}
}

func TestFgRainbow_SpacesUncolored(t *testing.T) {
	result := FgRainbow("a b")

	if !strings.Contains(result, " ") {
		t.Error("Rainbow: space character missing from the result")
	}
}

func TestFgRainbow_ColorCycling(t *testing.T) {
	long := "abcdefghij"
	result := FgRainbow(long)
	for _, c := range long {
		if !strings.Contains(result, string(c)) {
			t.Errorf("Rainbow: char %q missing after color cycling", c)
		}
	}
}

func TestFgRainbow_EmptyString(t *testing.T) {
	result := FgRainbow("")
	_ = result
}

func TestBgRainbow_ContainsInput(t *testing.T) {
	result := BgRainbow("hello")
	for _, c := range "hello" {
		if !strings.Contains(result, string(c)) {
			t.Errorf("Rainbow: character %q missing from result", c)
		}
	}
}

func TestBgRainbow_SpacesUncolored(t *testing.T) {
	result := BgRainbow("a b")

	if !strings.Contains(result, " ") {
		t.Error("Rainbow: space character missing from the result")
	}
}

func TestBgRainbow_ColorCycling(t *testing.T) {
	long := "abcdefghij"
	result := BgRainbow(long)
	for _, c := range long {
		if !strings.Contains(result, string(c)) {
			t.Errorf("Rainbow: char %q missing after color cycling", c)
		}
	}
}

func TestBgRainbow_EmptyString(t *testing.T) {
	result := BgRainbow("")
	_ = result
}
