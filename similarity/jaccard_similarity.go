package similarity

import (
	"fingerprint/fingerprint"
	"github.com/deckarep/golang-set"
)

func Jaccard(s1, s2 fingerprint.Fingerprint) float64 {
	s1set := convertFingerToSet(s1)
	s2set := convertFingerToSet(s2)
	intersectCardinality := s1set.Intersect(s2set).Cardinality()
	unionCardinality := s1set.Union(s2set).Cardinality()
	return float64(intersectCardinality) / float64(unionCardinality)
}

func convertFingerToSet(s fingerprint.Fingerprint) mapset.Set {
	set := mapset.NewSet()
	for _, token := range s {
		set.Add(token)
	}
	return set
}
