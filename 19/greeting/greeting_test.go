package greeting

import (
	"testing"
)

func TestHello_English(t *testing.T) {

	got := Hello("Aaron", "en")
	want := "Hello, Aaron"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestHello_Spanish(t *testing.T) {
	got := Hello("Aaron", "es")
	want := "Hola, Aaron"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestHello_French(t *testing.T) {
	got := Hello("Aaron", "fr")
	want := "Bonjour, Aaron"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestIsEmpty(t *testing.T) {
	got := IsEmpty("")
	want := true

	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

// Table-driven tests
type test struct {
	name     string
	lang     string
	expected string
}

func TestHello_Table(t *testing.T) {

	tests := []test{
		{name: "Aaron", lang: "en", expected: "Hello, Aaron"},
		{name: "Aaron", lang: "es", expected: "Hola, Aaron"},
		{name: "Aaron", lang: "fr", expected: "Bonjour, Aaron"},
	}

	for _, tt := range tests {
		t.Run(tt.name+"_"+tt.lang, func(t *testing.T) {
			if got := Hello(tt.name, tt.lang); got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}
