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
	result := FgAnsi("hello", ARed)

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
	result := BgAnsi("hello", ARed)

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
	result := FgAnsi("", ABlue)
	if !strings.HasSuffix(result, "\033[0m") {
		t.Errorf("FgAnsi empty string: expected reset suffix, got %q", result)
	}
}

func TestBgAnsi_EmptyString(t *testing.T) {
	result := FgAnsi("", AYellow)
	if !strings.HasSuffix(result, "\033[0m") {
		t.Errorf("BgAnsi empty string: expected reset suffix, got %q", result)
	}
}

// Preset ANSI colros - code values

func TestPresetAnsiColors_FgCodes(t *testing.T) {
	tests := []struct{
		name string
		color ansiColor
		fg string
	}{
		{"ARed", ARed, "\033[31m"},
		{"AGreen", AGreen, "\033[32m"},
		{"AYellow", AYellow, "\033[33m"},
		{"ABlue", ABlue, "\033[34m"},
		{"APurple", APurple, "\033[35m"},
		{"ACyan", ACyan, "\033[36m"},
		{"AWhite", AWhite, "\033[37m"},
	}

	for _, tt := range tests {
		t.Run(tt.name+"_fg", func(t *testing.T) {
			result := FgAnsi("x", tt.color)
			if !strings.HasPrefix(result, tt.fg) {
				t.Errorf("%s fg: expected prefix %q, got %q", tt.name, tt.fg, result)
			}
		})
	}
}

func TestPresetAnsiColors_BgCodes(t *testing.T) {
	tests := []struct{
		name string
		color ansiColor
		bg string
	}{
		{"ARed", ARed, "\033[41m"},
		{"AGreen", AGreen, "\033[42m"},
		{"AYellow", AYellow, "\033[43m"},
		{"ABlue", ABlue, "\033[44m"},
		{"APurple", APurple, "\033[45m"},
		{"ACyan", ACyan, "\033[46m"},
		{"AWhite", AWhite, "\033[47m"},
	}

	for _, tt := range tests {
		t.Run(tt.name+"_bg", func(t *testing.T) {
			result := BgAnsi("x", tt.color)
			if !strings.HasPrefix(result, tt.bg) {
				t.Errorf("%s bg: expected prefix %q, got %q", tt.name, tt.bg, result)
			}
		})
	}
}

// AnsiTemplate

func TestCreateAnsiTemplate(t *testing.T) {
	tmpl := CreateAnsiTemplate(ABlue, AWhite)

	if tmpl == nil {
		t.Fatal("CreateAnsiTemplate returned nil")
	}

	if tmpl.Bg != ABlue {
		t.Errorf("Bg = %v, want ABlue", tmpl.Bg)
	}

	if tmpl.Fg != AWhite {
		t.Errorf("Fg = %v, want AWhite", tmpl.Fg)
	}
}

func TestAnsiTemplate_FormatString(t *testing.T) {
	tmpl := CreateAnsiTemplate(ABlue, AWhite)
	result := tmpl.FormatString("test")

	bgIdx := strings.Index(result, ABlue.bg)
	fgIdx := strings.Index(result, AWhite.fg)

	if bgIdx == -1 {
		t.Fatalf("FormatString: bg escape %q not found in %q", ABlue.bg, result)
	}

	if fgIdx == -1 {
		t.Fatalf("FormatString: fg escape %q not found in %q", ABlue.fg, result)
	}

	if bgIdx > fgIdx {
		t.Errorf("FormatString: expected bg to wrap fg (bg at %d should be before fg at %d)", bgIdx, fgIdx)
	}

	if !strings.Contains(result, "test") {
		t.Errorf("FormatString: original text missing in %q", result)
	}

	if !strings.Contains(result, "\033[0m") {
		t.Errorf("FormatString: reset escape sequence is missing in %q", result)
	}
}

// Rainbow

func TestRainbow_ContainsInput(t *testing.T) {
	result := Rainbow("hello", "", 0)
	for _, c := range "hello" {
		if !strings.Contains(result, string(c)) {
			t.Errorf("Rainbow: character %q missing from result", c)
		}
	}
}

func TestRainbow_SpacesUncolored(t *testing.T) {
	result := Rainbow("a b", "", 0)

	if !strings.Contains(result, " ") {
		t.Error("Rainbow: space character missing from the result")
	}
}

func TestRainbow_Escape_AddsNewLines(t *testing.T) {
	result := Rainbow("hi", "", 2)
	count := strings.Count(result, "\n")
	if count != 2 {
		t.Errorf("Rainbow escape=2: expected 2 newlines, got %d", count)
	}
}

func TestRainbow_ZeroEscape_NoNewLines(t *testing.T) {
	result := Rainbow("hi", "", 0)
	if strings.Contains(result, "\n") {
		t.Errorf("Rainbow escape=0: unexpected newline in result %q", result)
	}
}

func TestRainbow_Mod_Bold(t *testing.T) {
	result := Rainbow("hi", "bold", 0)
	if !strings.Contains(result, "\033[1m") {
		t.Errorf("Rainbow bold: expected bold escape code, got %q", result)
	}
}

func TestRainbow_Mod_Unknown_DoesNotPanic(t *testing.T) {
	defer func(){
		if r := recover(); r != nil {
			t.Errorf("Rainbow panicked with unknown mod: %v", r)
		}
	}()
	Rainbow("hi", "unknown", 0)
}

func TestRainbow_ColorCycling(t *testing.T) {
	long := "abcdefghij"
	result := Rainbow(long, "", 0)
	for _, c := range long {
		if !strings.Contains(result, string(c)) {
			t.Errorf("Rainbow: char %q missing after color cycling", c)
		}
	}
}

func TestRainbow_EmptyString(t *testing.T) {
	result := Rainbow("", "", 0)
	_ = result
}
