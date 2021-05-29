package fingerprinter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKGramHash(t *testing.T) {
	t.Run("returns an empty hash vector for an empty kgram vector", func(t *testing.T) {
		assert.Equal(t, []uint32{}, KGramHash([]string{}))
	})
	t.Run("returns a single hash for a single kgram", func(t *testing.T) {
		assert.Equal(t, []uint32{440920331}, KGramHash([]string{"abc"}))
	})
	t.Run("returns a hash for a kgram", func(t *testing.T) {
		given := []string{
			"adoru", "dorun", "orunr", "runru", "unrun", "nrunr", "runru", "unrun", "nruna", "runad", "unado", "nador", "adoru", "dorun", "orunr", "runru", "unrun",
		}
		expected := []uint32{0xf765d270, 0xce399891, 0x9d4bf9df, 0x88294ce7, 0x2f4243db, 0x7116aba4, 0x88294ce7, 0x2f4243db, 0x8416c98d, 0x8908facf, 0x43870890, 0x1a820a81, 0xf765d270, 0xce399891, 0x9d4bf9df, 0x88294ce7, 0x2f4243db}

		assert.Equal(t, expected, KGramHash(given))
	})
}
