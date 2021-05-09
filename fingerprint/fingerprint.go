package fingerprint

import (
	text2 "fingerprint/text"
)

type FingerPrintOptions struct {
	GuaranteeThreshold int // t
	NoiseThreshold     int // k
}

func (o FingerPrintOptions) IsValid() (bool, error) {
	return o.NoiseThreshold <= o.GuaranteeThreshold, nil
}

func Record(k int, text string) (fingerprints []uint32) {
	return KGramHash(KGram(k, text2.Clean(text)))
}

func SomeFingerPrintFunction(input string, options FingerPrintOptions) Finger {
	return WinnowFingerprint(options.GuaranteeThreshold, Record(options.NoiseThreshold, text2.Clean(input)))
}
