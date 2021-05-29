package fingerprint

import "math"

func Winnow(g int, kgrams []uint32) (winnow [][]uint32) {
	winnow = make([][]uint32, 0)

	length := len(kgrams)
	if length == 0 {
		return winnow
	}

	for i := g; i <= length; i++ {
		winnow = append(winnow, kgrams[i-g:i])
	}

	return winnow
}

func WinnowFingerprint(g int, kgrams []uint32) (finger Fingerprint) {
	values := Winnow(g, kgrams)

	for _, value := range values {
		finger = append(finger, RightmostLowestValue(value))
	}

	return finger
}

type Mark struct {
	MinValue uint32
	Index    int
}

type Fingerprint []Mark

func RightmostLowestValue(values []uint32) (w Mark) {
	w.MinValue = math.MaxUint32
	for i, v := range values {
		if v <= w.MinValue {
			w.MinValue = v
			w.Index = i
		}
	}
	return w
}

type MarkSet map[Mark]bool

func (f Fingerprint) AsSet() MarkSet {
	set := make(MarkSet, 0)
	for _, mark := range f {
		set[mark] = true
	}
	return set
}
