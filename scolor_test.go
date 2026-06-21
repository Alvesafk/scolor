/*
Author © 2026 alvesafk <migueldealmeidaalves55@gmail.com>

Tests for scolor.go
*/
package scolor

import (
	"strings"
	"testing"
)

// Color constructor

func TestRGB(t *testing.T) {
	c := RGB(10, 20, 30)
	if c.Red != 10 || c.Green != 20 || c.Blue != 30 {
		t.Errorf("RGB(10,20,30) = {%d,%d,%d}, want {10,20,30}",
		c.Red, c.Green, c.Blue)
	}
}

func TestRGB_Boundaries(t *testing.T) {
	black := RGB(0,0,0)
	white := RGB(255,255,255)

	if black.Red != 0 || black.Green != 0 || black.Blue != 0 {
		t.Errorf("RGB(0,0,0): unexpected values")
	}
	if white.Red != 255 || white.Green != 255 || white.Blue != 255 {
		t.Errorf("RGB(255,255,255): unexpected values")
	}
}

// FgRGB / BgRGB - escape sequence format

func TestFgRGB_EscapeSequence(t *testing.T) {
	c := RGB(200, 100, 50)
	result := FgRGB("hello", c)

	if !strings.HasPrefix(result, "\x1b[38;2") {
		t.Errorf("FgRGB: expected foreground escape prefix, got %q", result)
	}

	if !strings.Contains(result, "200;100;50m") {
		t.Errorf("FgRGB: RGB values not found in escape sequence, got %q", result)
	}

	if !strings.Contains(result, "hello") {
		t.Errorf("FgRGB: original string not present in result, got %q", result)
	}

	if !strings.HasSuffix(result, "\x1b[0m") {
		t.Errorf("FgRGB: expected reset suffix, got %q", result)
	}
}

func TestBgRGB_EscapeSequence(t *testing.T) {
	c := RGB(10, 20, 30)
	result := BgRGB("world", c)

	if !strings.HasPrefix(result, "\x1b[48;2") {
		t.Errorf("BgRGB: expected foreground escape prefix, got %q", result)
	}

	if !strings.Contains(result, "10;20;30m") {
		t.Errorf("BgRGB: RGB values not found in escape sequence, got %q", result)
	}

	if !strings.Contains(result, "world") {
		t.Errorf("BgRGB: original string not present in result, got %q", result)
	}

	if !strings.HasSuffix(result, "\x1b[0m") {
		t.Errorf("BgRGB: expected reset suffix, got %q", result)
	}
}

func TestFgRGB_EmptyString(t *testing.T) {
	result := FgRGB("", RGB(255, 0, 0))
	if !strings.HasSuffix(result, "\x1b[0m") {
		t.Errorf("FgRGB with empty string: expected reset suffix, got %q", result)
	}
}

func TestBgRGB_EmptyString(t *testing.T) {
	result := BgRGB("", RGB(255, 0, 0))
	if !strings.HasSuffix(result, "\x1b[0m") {
		t.Errorf("BgRGB with empty string: expected reset suffix, got %q", result)
	}
}

// Preset colors

func TestPresetColors(t *testing.T) {
	tests := []struct {
		name string
		color Color
		r, g, b uint8
	}{
		{"BLACK", BLACK, 0, 0, 0},
		{"WHITE", WHITE, 255, 255, 255},
		{"RED", RED, 200, 0, 0},
		{"GREEN", GREEN, 0, 200, 0},
		{"BLUE", BLUE, 0, 0, 200},
		{"CYAN", CYAN, 0, 200, 200},
		{"YELLOW", YELLOW, 200, 200, 0},
		{"ORANGE", ORANGE, 200, 115, 0},
		{"PINK", PINK, 200, 140, 150},
		{"PURPLE", PURPLE, 80, 0, 80},
		{"BROWN", BROWN, 110, 20, 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.color.Red != tt.r || tt.color.Green != tt.g || tt.color.Blue != tt.b {
				t.Errorf("%s = {%d,%d,%d}, want {%d,%d,%d}",
					tt.name, tt.color.Red, tt.color.Green, tt.color.Blue,
					tt.r, tt.g, tt.b)
			}
		})
	}
}

// RgbTemplate

func TestCreateRgbTemplate(t *testing.T) {
	bg := RGB(10, 20, 30)
	fg := RGB(200, 200, 200)
	tmpl := CreateRgbTemplate(bg, fg)

	if tmpl == nil {
		t.Fatal("CreateRgbTemplate returned nil")
	}

	if tmpl.Bg != bg {
		t.Errorf("Bg = %v, want %v", tmpl.Bg, bg)
	}

	if tmpl.Fg != fg {
		t.Errorf("Fg = %v, want %v", tmpl.Fg, bg)
	}
}

func TestRgbTemplate_FormatString(t *testing.T) {
	bg := RGB(0, 0, 0)
	fg := RGB(255, 255, 255)
	tmpl := CreateRgbTemplate(bg, fg)

	result := tmpl.FormatString("test")

	if !strings.Contains(result, "\x1b[48;2;") {
		t.Errorf("FormatString: background escape sequence missing, got %q", result)
	}

	if !strings.Contains(result, "\x1b[38;2;") {
		t.Errorf("FormatString: foreground escape sequence missing, got %q", result)
	}

	if !strings.Contains(result, "test") {
		t.Errorf("FormatString: original text missing, got %q", result)
	}
}

func TestRgbTemplate_FormatString_LayerOrder(t *testing.T) {
	bg := RGB(10, 20, 30)
	fg := RGB(200, 100, 50)
	tmpl := CreateRgbTemplate(bg, fg)

	result := tmpl.FormatString("x")

	bgIdx := strings.Index(result, "\x1b[48;2;")
	fgIdx := strings.Index(result, "\x1b[38;2;")

	if bgIdx == -1 || fgIdx == -1 {
		t.Fatal("missing escape sequences in FormatString output")
	}

	if bgIdx > fgIdx {
		t.Errorf("expected BgRGB to wrap FgRGB (bg before fg), got bg at %d, fg at %d", bgIdx, fgIdx)
	}
}

// enviroment awarness 

func TestIsTTY_ReturnsBool(t *testing.T) {
	result := isTTY()
	_ = result
}

func TestHasTrueColor_ReturnsBool(t *testing.T) {
	result := hasTrueColor()
	_ = result
}

func TestIsRGBSupport_ConsistentWithComponents(t *testing.T) {
	expected := isTTY() && hasTrueColor()
	if IsRGBSupported != expected {
		t.Errorf("IsRGBSupported = %v, want %v (isTTY=%v, hasTrueColor=%v)",
			IsRGBSupported, expected, isTTY(), hasTrueColor())
	}
}
