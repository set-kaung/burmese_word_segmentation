package tests

import (
	"testing"

	"github.com/Set-Kaung/blitter"
)

func TestSplitIntoWords(t *testing.T) {
	sentence := "ခေတ်အလိုက်ပြဋ္ဌာန်းပြီးဥပဒေများ"
	blitter.Splitter(sentence)
	sentence = "နတ်"
	sl := blitter.CreateWordsSlice(blitter.Splitter(sentence))
	if sl[0] != sentence {
		t.Errorf("want %s, got %s", sentence, sl[0])
	}
}
