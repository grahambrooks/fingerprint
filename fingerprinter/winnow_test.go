package fingerprinter

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestWinnow(t *testing.T) {
	tests := []struct {
		name     string
		g        int
		expected [][]uint32
		values   []uint32
	}{
		{name: "k-gram less that g", expected: [][]uint32{}, g: 0, values: []uint32{}},
		{name: "returns a single k-gram if length equal to g", g: 3, expected: [][]uint32{{1, 2, 3}}, values: []uint32{1, 2, 3}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, Winnow(test.g, test.values))
		})
	}
}

func TestWinnowFingerprinting(t *testing.T) {
	tests := []struct {
		name        string
		fingerprint Fingerprint
		g           int
		kgrams      []uint32
	}{
		{name: "empty k-grams", fingerprint: Fingerprint{NewMark(0x2, 2), NewMark(0x1, 2)}, g: 3, kgrams: []uint32{4, 3, 2, 1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.fingerprint, WinnowFingerprint(test.g, test.kgrams))
		})
	}
}

func TestRightmostLowestValue(t *testing.T) {
	tests := []struct {
		name     string
		expected Mark
		values   []uint32
	}{
		{"empty", NewMark(math.MaxUint32, 0), []uint32{}},
		{"single entry", NewMark(1, 0), []uint32{1}},
		{"min in the right most position", NewMark(1, 2), []uint32{100, 10, 1}},
		{"min in the right most position repeated", NewMark(1, 4), []uint32{1, 100, 10, 1, 1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, RightmostLowestValue(test.values))
		})
	}
}
