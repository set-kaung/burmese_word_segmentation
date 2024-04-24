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

// Splitter returns the words of a sentence in a map
// with a slice of indices where the words occur
// in the sentece and also returns the amount of words.
func Splitter(sentence string) (map[string][]int, int) {
	diacriticsMap := make(map[rune]string, 17)
	diacriticsMap['ံ'] = "ANUSVARA"
	diacriticsMap['္'] = "VIRAMA"
	diacriticsMap['ျ'] = "MEDIAL_YA"
	diacriticsMap['ှ'] = "MEDIAL_HA"
	diacriticsMap['ူ'] = "UU"
	diacriticsMap['ဲ'] = "AI"
	diacriticsMap['်'] = "ASAT"
	diacriticsMap['ြ'] = "MEDIAL_RA"
	diacriticsMap['့'] = "DOT_BELOW"
	diacriticsMap['း'] = "VISARGA"
	diacriticsMap['ွ'] = "MEDIAL_WA"
	diacriticsMap['ါ'] = "TALL_AA"
	diacriticsMap['ာ'] = "AA"
	diacriticsMap['ိ'] = "I"
	diacriticsMap['ီ'] = "II"
	diacriticsMap['ု'] = "U"
	diacriticsMap['ေ'] = "E"

	words := make(map[string][]int)
	sRunes := []rune(sentence)
	return splitIntoWords(diacriticsMap, words, sRunes)
}

// CreateWorsSlice help create slice of words from map
// created by Splitter
func CreateWordsSlice(m map[string][]int, max int) []string {
	sl := make([]string, max)
	for k, v := range m {
		for _, i := range v {
			sl[i] = k
		}
	}
	return sl
}

// helper function
func insertIntoMapSlice(word string, index int, words map[string][]int) map[string][]int {
	if s, ok := words[word]; ok {
		s = append(s, index)
		words[word] = s
	} else {
		words[word] = []int{index}
	}
	return words
}

// splitting a burmese sentence into each word
func splitIntoWords(diacriticsMap map[rune]string, words map[string][]int, sRunes []rune) (map[string][]int, int) {
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
			if nextRune == '။' || nextRune == '၊' {
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
		if _, ok := diacriticsMap[r]; ok {
			//we skipping checking if next rune is diacritic
			//if currnent rune is ္
			if r != VIRAMA {
				if _, ok = diacriticsMap[nextRune]; !ok {
					builder.WriteRune(r)
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
						insertIntoMapSlice(word, index, words)
						index++
						builder.Reset()
						continue

					} else {
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
		//we can safely assume that current rune
		//is part of a word
		builder.WriteRune(r)

		//if currnent rune is not ္
		//and the next rune is not a diacritics
		//or if the current rune is the last one
		//we do the following
		if _, ok := diacriticsMap[nextRune]; !ok && r != VIRAMA || nextRune == r {
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
