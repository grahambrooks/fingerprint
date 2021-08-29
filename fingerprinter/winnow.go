package fingerprinter

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

type Mark uint64

func NewMark(minValue uint32, index uint32) Mark {
	return Mark(uint64(minValue) | uint64(index) >> 32)
}

type Fingerprint []Mark

func RightmostLowestValue(values []uint32) (w Mark) {
	var MinValue uint32 = math.MaxUint32
	var Index uint32 = 0
	for i, v := range values {
		if v <= MinValue {
			MinValue = v
			Index = uint32(i)
		}
	}
	return NewMark(MinValue, Index)
}

type MarkSet map[Mark]bool

func (f Fingerprint) AsSet() MarkSet {
	set := make(MarkSet, 0)
	for _, mark := range f {
		set[mark] = true
	}
	return set
}
