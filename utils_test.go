/*
Author © 2026 alvesafk <migueldealmeidaalves55@gmail.com>

Tests for utils.go
*/

package scolor

import "testing"

func TestAddMod(t *testing.T) {
	tests := []struct {
		name string
		mod  int
		want string
	}{
		{name: "bold", mod: Bold, want: "\033[1mhello"},
		{name: "underline", mod: Underline, want: "\033[4mhello"},
		{name: "strike", mod: Strike, want: "\033[9mhello"},
		{name: "italic", mod: Italic, want: "\033[3mhello"},
		{name: "unknown mod returns original string", mod: 999, want: "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddMod("hello", tt.mod); got != tt.want {
				t.Fatalf("AddMod() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestRemoveEscapeSequence(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "plain string",
			in:   "hello",
			want: "hello",
		},
		{
			name: "single modifier",
			in:   "\033[1mhello",
			want: "hello",
		},
		{
			name: "foreground RGB sequence with reset",
			in:   "\x1b[38;2;1;2;3mhello\x1b[0m",
			want: "hello",
		},
		{
			name: "multiple sequences",
			in:   "\033[1mhello \x1b[48;2;4;5;6mworld\x1b[0m",
			want: "hello world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveEscapeSequence(tt.in); got != tt.want {
				t.Fatalf("RemoveEscapeSequence() = %q, want %q", got, tt.want)
			}
		})
	}
}
