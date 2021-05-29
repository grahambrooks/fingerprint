package similarity

import (
	"github.com/grahambrooks/fingerprint/fingerprinter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringSimilarity(t *testing.T) {
	t.Run("should return 1 for identical strings", func(t *testing.T) {
		text := "the quick brown fox jumped over the lazy dog"
		assert.Equal(t, 1.0, StringSimilarity(text, text, fingerprinter.Options{}))
	})
}

func Test_intersect(t *testing.T) {
	t.Run("intersect of two empty fingerprints is empty", func(t *testing.T) {
		s1 := make(fingerprinter.Fingerprint, 0)
		s2 := make(fingerprinter.Fingerprint, 0)
		set := intersect(s1, s2)
		assert.Empty(t, set)
	})

	t.Run("intersect of a fingerprinter and an empty fingerprinter is empty", func(t *testing.T) {
		s1 := make(fingerprinter.Fingerprint, 0)
		s1 = append(s1, fingerprinter.Mark{MinValue: 1, Index: 2})
		s2 := make(fingerprinter.Fingerprint, 0)
		set := intersect(s1, s2)
		assert.Len(t, set, 0)

	})
	t.Run("intersect of two fingerprints with the same mark is not empty", func(t *testing.T) {
		s1 := make(fingerprinter.Fingerprint, 0)
		s1 = append(s1, fingerprinter.Mark{MinValue: 1, Index: 2})
		s1 = append(s1, fingerprinter.Mark{MinValue: 3, Index: 4})
		s2 := make(fingerprinter.Fingerprint, 0)
		s2 = append(s2, fingerprinter.Mark{MinValue: 1, Index: 2})
		set := intersect(s1, s2)
		assert.Len(t, set, 1)
		assert.Contains(t, set, fingerprinter.Mark{MinValue: 1, Index: 2})

	})
}

func Test_union(t *testing.T) {
	t.Run("union of two empty fingerprints is empty", func(t *testing.T) {
		s1 := make(fingerprinter.Fingerprint, 0)
		s2 := make(fingerprinter.Fingerprint, 0)
		set := union(s1, s2)
		assert.Empty(t, set)
	})

	t.Run("union of a fingerprinter and an empty fingerprinter is not empty", func(t *testing.T) {
		s1 := make(fingerprinter.Fingerprint, 0)
		s1 = append(s1, fingerprinter.Mark{MinValue: 1, Index: 2})
		s2 := make(fingerprinter.Fingerprint, 0)
		set := union(s1, s2)
		assert.Len(t, set, 1)
		assert.Contains(t, set, fingerprinter.Mark{MinValue: 1, Index: 2})

	})
	t.Run("intersect of two fingerprints with the same mark is not empty", func(t *testing.T) {
		s1 := make(fingerprinter.Fingerprint, 0)
		s1 = append(s1, fingerprinter.Mark{MinValue: 1, Index: 2})
		s1 = append(s1, fingerprinter.Mark{MinValue: 3, Index: 4})
		s2 := make(fingerprinter.Fingerprint, 0)
		s2 = append(s2, fingerprinter.Mark{MinValue: 1, Index: 2})
		set := union(s1, s2)
		assert.Len(t, set, 2)
		assert.Contains(t, set, fingerprinter.Mark{MinValue: 1, Index: 2})
		assert.Contains(t, set, fingerprinter.Mark{MinValue: 3, Index: 4})

	})
}
