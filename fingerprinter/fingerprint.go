package fingerprinter

import (
	"github.com/grahambrooks/fingerprint/text"
)

type Options struct {
	GuaranteeThreshold int // t
	NoiseThreshold     int // k
}

func (o Options) IsValid() bool {
	return o.NoiseThreshold <= o.GuaranteeThreshold && o.NoiseThreshold > 0
}

func (o Options) VerifyOrDefault() Options {
	if !o.IsValid() {
		return Options{GuaranteeThreshold: 4, NoiseThreshold: 4}
	}
	return o
}

func record(k int, sourceText string) (fingerprints []uint32) {
	return KGramHash(KGram(k, text.Clean(sourceText)))
}

func TextFingerprint(input string, options Options) Fingerprint {
	options = options.VerifyOrDefault()
	return WinnowFingerprint(options.GuaranteeThreshold, record(options.NoiseThreshold, text.Clean(input)))
}
