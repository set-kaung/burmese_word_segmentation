package blitter

import "strings"

const (
	TALL_AA   rune = 'ါ'
	AA        rune = 'ာ'
	I         rune = 'ိ'
	II        rune = 'ီ'
	U         rune = 'ု'
	UU        rune = 'ူ'
	E         rune = 'ေ'
	AI        rune = 'ဲ'
	ANUSVARA  rune = 'ံ'
	DOT_BELOW rune = '့'
	VISARGA   rune = 'း'
	VIRAMA    rune = '္'
	ASAT      rune = '်'
	MEDIAL_YA rune = 'ျ'
	MEDIAL_RA rune = 'ြ'
	MEDIAL_WA rune = 'ွ'
	MEDIAL_HA rune = 'ှ'
)

func isDiacritic(r rune) bool {
	switch r {
	case ANUSVARA, VIRAMA, MEDIAL_YA, MEDIAL_HA, UU, AI, ASAT,
		MEDIAL_RA, DOT_BELOW, VISARGA, MEDIAL_WA, TALL_AA, AA, I, II, U, E:
		return true
	}
	return false
}

// Segment returns the Burmese syllables of a sentence
func SegmentSyllables(sentence string) ([]string, error) {
	sRunes := []rune(sentence)
	return segment(sRunes)
}

// splitting a burmese sentence into each syllable,
// Ignore spaces but include punctuation
func segment(sRunes []rune) (words []string, err error) {
	words = make([]string, 0, len(sRunes))
	index := 0
	builder := strings.Builder{}
	var nextRune rune
	for i := 0; i < len(sRunes); i++ {
		r := sRunes[i]

		if r == '\r' || r == ' ' || r == '\n' {
			continue
		}

		//checking if end
		if r == '။' || r == '၊' {
			_, err = builder.WriteRune(r)
			if err != nil {
				return nil, err
			}
			word := builder.String()
			words = append(words, word)
			index++
			builder.Reset()
			continue
		}
		//checking whether index out of bounds
		//for nextRune.
		//if out of bound current rune
		//and next is the same
		//this is current rune is the last one
		if i != len(sRunes)-1 {
			nextRune = sRunes[i+1]
			if nextRune == '။' || nextRune == '၊' {
				_, err = builder.WriteRune(r)
				if err != nil {
					return nil, err
				}
				word := builder.String()
				words = append(words, word)
				index++
				builder.Reset()
				continue
			}

		} else {
			nextRune = r
		}
		if isDiacritic(r) {
			//we skipping checking if next rune is diacritic
			//if currnent rune is ္
			if r != VIRAMA {
				if !isDiacritic(nextRune) {
					_, err = builder.WriteRune(r)
					if err != nil {
						return nil, err
					}
					if i+2 <= len(sRunes)-1 {
						//we check if the next rune is
						//something like တ်
						//if it is, then current word in the buffer is
						//something like န(တ်)
						n2 := sRunes[i+2]
						if n2 == ASAT || n2 == DOT_BELOW {
							continue
						}
						word := builder.String()
						words = append(words, word)
						index++
						builder.Reset()
						continue

					} else {
						word := builder.String()
						words = append(words, word)
						index++
						builder.Reset()
						continue
					}
				}
			}
		}
		// if all above procedures isn't executed
		//we can safely assume that current rune
		//is part of a word
		_, err = builder.WriteRune(r)
		if err != nil {
			return nil, err
		}

		//if currnent rune is not ္
		//and the next rune is not a diacritics
		//or if the current rune is the last one
		//we do the following
		switch {
		case (r != VIRAMA && !isDiacritic(nextRune)) || nextRune == r:
			// Again checking for something like နတ်
			if i+2 <= len(sRunes)-1 {
				switch sRunes[i+2] {
				case ASAT, DOT_BELOW:
					continue
				}
			}
			word := builder.String()
			words = append(words, word)
			index++
			builder.Reset()
		}
	}
	return words, nil
}
