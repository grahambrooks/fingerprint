package fingerprint

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
		fingerprint Finger
		g           int
		kgrams      []uint32
	}{
		{name: "empty k-grams", fingerprint: Finger{Window{Min: 0x2, Index: 2}, Window{Min: 0x1, Index: 2}}, g: 3, kgrams: []uint32{4, 3, 2, 1}},
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
		expected Window
		values   []uint32
	}{
		{"empty", Window{Min: math.MaxUint32, Index: 0}, []uint32{}},
		{"single entry", Window{Min: 1, Index: 0}, []uint32{1}},
		{"min in the right most position", Window{Min: 1, Index: 2}, []uint32{100, 10, 1}},
		{"min in the right most position repeated", Window{Min: 1, Index: 4}, []uint32{1, 100, 10, 1, 1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, RightmostLowestValue(test.values))
		})
	}
}
