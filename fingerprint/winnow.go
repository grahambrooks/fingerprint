package fingerprint

import "math"

func Winnow(g int, kgrams []uint32) (winnow [][]uint32) {
	winnow = make([][]uint32, 0)

	len := len(kgrams)
	if len == 0 {
		return winnow
	}

	for i := g; i <= len; i++ {
		winnow = append(winnow, kgrams[i-g:i])
	}

	return winnow
}

func WinnowFingerprint(g int, kgrams []uint32) (finger Finger) {
	values := Winnow(g, kgrams)

	for _, value := range values {
		finger = append(finger, RightmostLowestValue(value))
	}

	return finger
}

type Window struct {
	Min   uint32
	Index int
}

type Finger []Window

func RightmostLowestValue(values []uint32) (w Window) {
	w.Min = math.MaxUint32
	for i, v := range values {
		if v <= w.Min {
			w.Min = v
			w.Index = i
		}
	}
	return w
}
