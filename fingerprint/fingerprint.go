package fingerprint

import (
	"github.com/grahambrooks/fingerprint/text"
)

type Options struct {
	GuaranteeThreshold int // t
	NoiseThreshold     int // k
}

func (o Options) IsValid() (bool, error) {
	return o.NoiseThreshold <= o.GuaranteeThreshold, nil
}

func Record(k int, sourceText string) (fingerprints []uint32) {
	return KGramHash(KGram(k, text.Clean(sourceText)))
}

func SomeFingerPrintFunction(input string, options Options) Fingerprint {
	return WinnowFingerprint(options.GuaranteeThreshold, Record(options.NoiseThreshold, text.Clean(input)))
}
