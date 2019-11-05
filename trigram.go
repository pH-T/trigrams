package trigrams

// Trigram (single) three uint8 --> uint32 (shifted)
type Trigram uint32

// Trigrams list of Trigrams
type Trigrams []Trigram

// CommonAndUnique returns the amount of unique trigrams and amount of common trigrams
func (t Trigrams) CommonAndUnique(other Trigrams) (int, int) {
	var found = make(map[Trigram]uint8) //map for performance
	unique := 0
	common := 0

	// add all unique trigrams to the found map and count the uniques in this process
	for _, i := range t {
		_, ok := found[i]
		if !ok {
			found[i] = 0
			unique++
		}
	}

	for _, i := range other {
		_, ok := found[i]
		if ok {
			found[i]++
		} else {
			found[i] = 0
			unique++
		}
	}

	for _, i := range found {
		if i > 0 {
			common++
		}
	}

	return common, unique

}

// JaccardCompare returns the amount of unique trigrams / amount of common trigrams.
// If the size difference is too big, 0.0 is returned (e.g. 1.2 --> < 20% size difference is okay)
// if acceptedSizeDifference is -1 the size difference is not considered
func (t Trigrams) JaccardCompare(other Trigrams, acceptedSizeDifference float64) float64 {
	// check for size difference as Jaccard does not consider size
	if float64(len(t))*acceptedSizeDifference < float64(len(other)) {
		return 0.0
	}

	if float64(len(other))*acceptedSizeDifference < float64(len(t)) {
		return 0.0
	}

	c, u := t.CommonAndUnique(other)

	// Jaccard coefficient
	return float64(c) / float64(u)

}

// ToTrigrams pads the input with \x00 and generates Trigrams, returning them
func ToTrigrams(input []byte) Trigrams {
	var trigrams Trigrams

	padded := []byte{0, 0} // pad it
	padded = append(padded, input...)
	padded = append(padded, 0) // pad it
	padded = append(padded, 0) // pad it

	for i := 0; i < len(padded)-2; i++ {
		x := Trigram(uint32(padded[i])<<16 | uint32(padded[i+1])<<8 | uint32(padded[i+2]))
		trigrams = append(trigrams, x)
	}

	return trigrams
}

// ToByte returns the given Trigram as bytes
func ToByte(input Trigram) []byte {
	return []byte{byte(input >> 16), byte(input >> 8), byte(input)}
}
