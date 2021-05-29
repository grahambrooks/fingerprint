package fingerprinter

import "hash/fnv"

func KGramHash(kgrams []string) (hashes []uint32) {
	hashes = make([]uint32, 0)
	for _, kgram := range kgrams {
		hashes = append(hashes, hash(kgram))
	}
	return hashes
}

func hash(s string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return h.Sum32()
}
