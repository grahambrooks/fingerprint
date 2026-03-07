# Fingerprint

A fingerprinting and similarity module for Go.

This is an implementation of the algorithm described in
the [Winnowing: Local Algorithms for Document Fingerprinting](http://igm.univ-mlv.fr/~mac/ENS/DOC/sigmod03-1.pdf) paper.
It can be used to fingerprint document text and calculate a similarity score using
the [Jaccard index (similarity coefficient)](https://en.wikipedia.org/wiki/Jaccard_index).

## Key Features

- **Winnowing Algorithm**: Robust local fingerprinting that is insensitive to small changes in the document.
- **Jaccard Similarity**: Easy comparison of fingerprints to determine document similarity.
- **Text Cleaning**: Built-in text normalization (removes non-letters and converts to lowercase) for more accurate
  matching.

## Installation

```bash
go get github.com/grahambrooks/fingerprint
```

## Usage

The library provides a high-level `similarity` package for quick comparisons and a `fingerprinter` package for more
granular control.

### Simple Similarity Comparison

```go
package main

import (
	"fmt"
	"github.com/grahambrooks/fingerprint/fingerprinter"
	"github.com/grahambrooks/fingerprint/similarity"
)

func main() {
	text1 := "The quick brown fox jumped over the lazy dog"
	text2 := "The quick brown fox jumps over the lazy dog"

	// Options define the sensitivity of the fingerprinting
	options := fingerprinter.Options{
		GuaranteeThreshold: 4,
		NoiseThreshold:     4,
	}

	score := similarity.StringSimilarity(text1, text2, options)

	fmt.Printf("Similarity score: %f\n", score)
}
```

### Understanding Options

- `NoiseThreshold` (k): Any match shorter than `k` is considered noise and ignored.
- `GuaranteeThreshold` (t): Any match of length at least `t` is guaranteed to be detected.

A common configuration is to set `NoiseThreshold` and `GuaranteeThreshold` to the same value (e.g., 4 or 5).

## Performance Note

The library is designed for relatively small documents and is not currently optimized for use with large streams or
channels.
