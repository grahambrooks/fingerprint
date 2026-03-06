package similarity

import (
	"github.com/grahambrooks/fingerprint/fingerprinter"
)

func StringSimilarity(s1, s2 string, options fingerprinter.Options) float64 {
	f1 := fingerprinter.TextFingerprint(s1, options)
	f2 := fingerprinter.TextFingerprint(s2, options)

	return Compare(f1, f2)
}

func Compare(f1, f2 fingerprinter.Fingerprint) float64 {
	set1 := f1.AsSet()
	set2 := f2.AsSet()

	unionCount := len(set1)
	intersectCount := 0
	for mark := range set2 {
		if set1[mark] {
			intersectCount++
		} else {
			unionCount++
		}
	}
	if unionCount == 0 {
		return 0
	}
	return float64(intersectCount) / float64(unionCount)
}
