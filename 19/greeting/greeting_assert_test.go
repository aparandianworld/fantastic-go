package greeting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello_With_Assert(t *testing.T) {
	assert.Equal(t, "Hello, Aaron", Hello("Aaron", "en"))
	assert.Equal(t, "Hola, Aaron", Hello("Aaron", "es"))
	assert.Equal(t, "Bonjour, Aaron", Hello("Aaron", "fr"))
}

func TestHello_Table_With_Assert(t *testing.T) {
	tests := []test{
		{name: "Aaron", lang: "en", expected: "Hello, Aaron"},
		{name: "Aaron", lang: "es", expected: "Hola, Aaron"},
		{name: "Aaron", lang: "fr", expected: "Bonjour, Aaron"},
	}

	for _, tt := range tests {
		t.Run(tt.name+"_"+tt.lang, func(t *testing.T) {
			got := Hello(tt.name, tt.lang)
			assert.Equal(t, tt.expected, got)
		})
	}
}
