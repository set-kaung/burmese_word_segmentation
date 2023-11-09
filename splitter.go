package splitter

import "strings"

const (
	TALL_AA   rune = 'ါ'
	AA             = 'ာ'
	I              = 'ိ'
	II             = 'ီ'
	U              = 'ု'
	UU             = 'ူ'
	E              = 'ေ'
	AI             = 'ဲ'
	ANUSVARA       = 'ံ'
	DOT_BELOW      = '့'
	VISARGA        = 'း'
	VIRAMA         = '္'
	ASAT           = '်'
	MEDIAL_YA      = 'ျ'
	MEDIAL_RA      = 'ြ'
	MEDIAL_WA      = 'ွ'
	MEDIAL_HA      = 'ှ'
)

// Splitter returns the words of a sentence in a map
// with a slice of indices where the words occur
// in the sentece and also returns the amount of words.
func Splitter(sentence string) (map[string][]int, int) {
	diacritics_map := make(map[rune]string, 17)
	diacritics_map['ံ'] = "ANUSVARA"
	diacritics_map['္'] = "VIRAMA"
	diacritics_map['ျ'] = "MEDIAL_YA"
	diacritics_map['ှ'] = "MEDIAL_HA"
	diacritics_map['ူ'] = "UU"
	diacritics_map['ဲ'] = "AI"
	diacritics_map['်'] = "ASAT"
	diacritics_map['ြ'] = "MEDIAL_RA"
	diacritics_map['့'] = "DOT_BELOW"
	diacritics_map['း'] = "VISARGA"
	diacritics_map['ွ'] = "MEDIAL_WA"
	diacritics_map['ါ'] = "TALL_AA"
	diacritics_map['ာ'] = "AA"
	diacritics_map['ိ'] = "I"
	diacritics_map['ီ'] = "II"
	diacritics_map['ု'] = "U"
	diacritics_map['ေ'] = "E"

	words := make(map[string][]int)
	sRunes := []rune(sentence)
	return splitIntoWords(diacritics_map, words, sRunes)

}

func insertIntoMapSlice(word string, index int, words map[string][]int) map[string][]int {
	if s, ok := words[word]; ok {
		s = append(s, index)
		words[word] = s
	} else {
		words[word] = []int{index}
	}
	return words
}

func splitIntoWords(diacritics_map map[rune]string, words map[string][]int, sRunes []rune) (map[string][]int, int) {
	index := 0
	builder := strings.Builder{}
	var nextRune rune
	for i := 0; i < len(sRunes); i++ {
		r := sRunes[i]
		if r == '။' {
			builder.WriteRune(r)
			word := builder.String()
			words = insertIntoMapSlice(word, index, words)
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
			if nextRune == '။' {
				builder.WriteRune(r)
				word := builder.String()
				words = insertIntoMapSlice(word, index, words)
				index++
				builder.Reset()
				continue
			}

		} else {
			nextRune = r
		}
		if _, ok := diacritics_map[r]; ok {
			//we skipping checking if next rune is diacritic
			//if currnent rune is ္
			if r != VIRAMA {
				if _, ok = diacritics_map[nextRune]; !ok {
					builder.WriteRune(r)
					if i+2 <= len(sRunes)-1 {
						//we check if the next rune is
						//something like တ်
						//if it is, then current word in the buffer is
						//something like နတ်
						n2 := sRunes[i+2]
						if n2 == ASAT || n2 == DOT_BELOW {
							continue
						}
						word := builder.String()
						insertIntoMapSlice(word, index, words)
						index++
						builder.Reset()
						continue
					}
				}
			}
		}
		// if all above procedures isn't executed
		//we can safe to assume that current rune
		//is part of a word
		builder.WriteRune(r)

		//if currnent rune is not ္
		//and the next rune is not a diacritics
		//or if the current rune is the last one
		//we do the following
		if _, ok := diacritics_map[nextRune]; !ok && r != VIRAMA || nextRune == r {
			//again checking for something like နတ်
			if i+2 <= len(sRunes)-1 {
				if sRunes[i+2] == ASAT || sRunes[i+2] == DOT_BELOW {
					continue
				}
			}
			word := builder.String()
			insertIntoMapSlice(word, index, words)
			index++
			builder.Reset()
		}
	}
	return words, index
}
