package fingerprinter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFingerprinting(t *testing.T) {
	t.Run("returns empty fingerprinter for empty text", func(t *testing.T) {
		assert.Equal(t, []uint32{}, record(4, ""))
	})

	t.Run("returns fingerprinter vector for given text", func(t *testing.T) {
		expected := []uint32{0xf765d270, 0xce399891, 0x9d4bf9df, 0x88294ce7, 0x2f4243db, 0x7116aba4, 0x88294ce7, 0x2f4243db, 0x8416c98d, 0x8908facf, 0x43870890, 0x1a820a81, 0xf765d270, 0xce399891, 0x9d4bf9df, 0x88294ce7, 0x2f4243db}
		assert.Equal(t, expected, record(5, "A do run run run, a do run run"))
	})
}

func TestFingerprintOptions(t *testing.T) {
	tests := []struct {
		name   string
		option Options
		result bool
		error  error
	}{
		{name: "k > t", option: Options{GuaranteeThreshold: 0, NoiseThreshold: 1}, result: false, error: nil},
		{name: "k == t", option: Options{GuaranteeThreshold: 1, NoiseThreshold: 1}, result: true, error: nil},
		{name: "k < t", option: Options{GuaranteeThreshold: 2, NoiseThreshold: 1}, result: true, error: nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			valid := test.option.IsValid()
			assert.Equal(t, test.result, valid)
		})
	}
}

func TestFingerprint_AsSet(t *testing.T) {
	t.Run("Empty fingerprinter empty set", func(t *testing.T) {
		fingerprint := make(Fingerprint, 0)
		set := fingerprint.AsSet()
		assert.Len(t, set, 0)
	})

	t.Run("set contains all marks from fingerprinter", func(t *testing.T) {
		fingerprint := make(Fingerprint, 0)
		fingerprint = append(fingerprint, NewMark(1, 2))
		fingerprint = append(fingerprint, NewMark(3, 4))
		set := fingerprint.AsSet()
		assert.Len(t, set, 2)
		assert.Contains(t, fingerprint, NewMark(1, 2))
		assert.Contains(t, fingerprint, NewMark(3, 4))
	})
}
