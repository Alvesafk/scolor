/*
Author © 2026 alvesafk <migueldealmeidaalves55@gmail.com>

Tests for utils.go
*/
package scolor

import (
	"strings"
	"testing"
)

// AddMod tests

func TestAddMod_Bold(t *testing.T) {
	result := AddMod("hello", "bold")
	if !strings.Contains(result, "\033[1m") {
		t.Errorf("AddMod bold: expected bold escape, got %q", result)
	}
	if !strings.Contains(result, "hello") {
		t.Errorf("AddMod bold: original string missing, got %q", result)
	}
}

func TestAddMod_Underine(t *testing.T) {
	result := AddMod("hello", "underline")
	if !strings.Contains(result, "\033[4m") {
		t.Errorf("AddMod underline: expected underline escape, got %q", result)
	}
	if !strings.Contains(result, "hello") {
		t.Errorf("AddMod underline: original string missing, got %q", result)
	}
}

func TestAddMod_Strike(t *testing.T) {
	result := AddMod("hello", "strike")
	if !strings.Contains(result, "\033[9m") {
		t.Errorf("AddMod strike: expected strike escape, got %q", result)
	}
	if !strings.Contains(result, "hello") {
		t.Errorf("AddMod strike: original string missing, got %q", result)
	}
}

func TestAddMod_Italic(t *testing.T) {
	result := AddMod("hello", "italic")
	if !strings.Contains(result, "\033[3m") {
		t.Errorf("AddMod italic: expected italic escape, got %q", result)
	}
	if !strings.Contains(result, "hello") {
		t.Errorf("AddMod italic: original string missing, got %q", result)
	}
}

func TestAddMod_Unknown_ReturnsUnchanged(t *testing.T) {
	input := "hello"
	result := AddMod(input, "rainbow")
	if result != input {
		t.Errorf("AddMod unknown mod: expected %q unchanged, got %q", input, result)
	}
}

func TestAddMod_EmptyMod_ReturnsUnchanged(t *testing.T) {
	input := "hello"
	result := AddMod(input, "")
	if result != input {
		t.Errorf("AddMod empty mod: expected %q unchanged, got %q", input, result)
	}
}

func TestAddMod_EmptyString(t *testing.T) {
	mods := []string{"bold", "underline", "strike", "italic"}
	for _, mod := range mods {
		t.Run(mod, func(t *testing.T) {
			result := AddMod("", mod)
			if result == "" {
				t.Errorf("AddMod(%q, %q): result should contain escape code, got empty string", "", mod)
			}
		})
	}
}

func TestAddMod_PreservesText(t *testing.T) {
	mods := []string{"bold", "underline", "strike", "italic"}
	text := "scolor"

	for _, mod := range mods {
		t.Run(mod, func(t *testing.T) {
			result := AddMod(text, mod)
			if !strings.Contains(result, text) {
				t.Errorf("AddMod(%q, %q): original text %q not found in result %q", text, mod, text, result)
			}
		})
	}
}

func TestAddMod_PrefixPosition(t *testing.T) {
	result := AddMod("x", "bold")
	boldIdx := strings.Index(result, "\033[1m")
	xIdx := strings.Index(result, "x")
	if boldIdx == -1 {
		t.Fatal("Bold escape not found")
	}

	if boldIdx > xIdx {
		t.Errorf("Bold escape (%d) should precede text (%d)", boldIdx, xIdx)
	}
}
