package text

import "testing"
import "github.com/stretchr/testify/assert"

func TestTextCleaning(t *testing.T) {
	t.Run("should return an empty string for an empty string", func(t *testing.T) {
		assert.Equal(t, "", Clean(``))
	})

	t.Run("should return an empty string for all punctuation strings", func(t *testing.T) {
		assert.Equal(t, "", Clean(`,/[] `))
	})

	t.Run("should return significant characters from string", func(t *testing.T) {
		assert.Equal(t, "adorunrunrunadorunrun", Clean(`A do run run run, a do run run`))

	})
}

