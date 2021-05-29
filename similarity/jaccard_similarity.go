package similarity

import (
	"github.com/grahambrooks/fingerprint/fingerprint"
)

func Jaccard(s1, s2 fingerprint.Fingerprint) float64 {
	intersectCardinality := len(intersect(s1, s2))
	unionCardinality := len(union(s1, s2))
	return float64(intersectCardinality) / float64(unionCardinality)
}

func intersect(f1, f2 fingerprint.Fingerprint) (set fingerprint.MarkSet) {
	s1 := f1.AsSet()
	s2 := f2.AsSet()
	set = make(fingerprint.MarkSet, 0)
	for mark := range s1 {
		if s2[mark] {
			set[mark] = true
		}
	}
	return set
}
func union(f1, f2 fingerprint.Fingerprint) (set fingerprint.MarkSet) {
	s1 := f1.AsSet()
	s2 := f2.AsSet()
	set = make(fingerprint.MarkSet, 0)
	for mark := range s1 {
		set[mark] = true
	}
	for mark := range s2 {
		set[mark] = true
	}
	return set
}
