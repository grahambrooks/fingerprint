package similarity

import (
	"github.com/grahambrooks/fingerprint/fingerprinter"
)

func StringSimilarity(s1, s2 string, options fingerprinter.Options) float64 {
	f1 := fingerprinter.TextFingerprint(s1, options)
	f2 := fingerprinter.TextFingerprint(s1, options)

	return Compare(f1, f2)
}

func Compare(s1, s2 fingerprinter.Fingerprint) float64 {
	intersectCardinality := len(intersect(s1, s2))
	unionCardinality := len(union(s1, s2))
	return float64(intersectCardinality) / float64(unionCardinality)
}

func intersect(f1, f2 fingerprinter.Fingerprint) (set fingerprinter.MarkSet) {
	s1 := f1.AsSet()
	s2 := f2.AsSet()
	set = make(fingerprinter.MarkSet, 0)
	for mark := range s1 {
		if s2[mark] {
			set[mark] = true
		}
	}
	return set
}
func union(f1, f2 fingerprinter.Fingerprint) (set fingerprinter.MarkSet) {
	s1 := f1.AsSet()
	s2 := f2.AsSet()
	set = make(fingerprinter.MarkSet, 0)
	for mark := range s1 {
		set[mark] = true
	}
	for mark := range s2 {
		set[mark] = true
	}
	return set
}
