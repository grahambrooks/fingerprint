package fingerprint

import (
	text2 "fingerprint/text"
)

type Options struct {
	GuaranteeThreshold int // t
	NoiseThreshold     int // k
}

func (o Options) IsValid() (bool, error) {
	return o.NoiseThreshold <= o.GuaranteeThreshold, nil
}

func Record(k int, text string) (fingerprints []uint32) {
	return KGramHash(KGram(k, text2.Clean(text)))
}

func SomeFingerPrintFunction(input string, options Options) Fingerprint {
	return WinnowFingerprint(options.GuaranteeThreshold, Record(options.NoiseThreshold, text2.Clean(input)))
}
