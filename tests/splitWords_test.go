package tests

import (
	"testing"

	"github.com/Set-Kaung/blitter"
)

func TestSplitIntoWords(t *testing.T) {
	sentence := "စိုက်ပျိုးရေး၊မွေးမြူရေးနှင့်ဆည်မြောင်းဝန်ကြီးဌာန"
	blitter.Splitter(sentence)
	sentence = "နတ်"
	sl := blitter.CreateWordsSlice(blitter.Splitter(sentence))
	if sl[0] != sentence {
		t.Errorf("want %s, got %s", sentence, sl[0])
	}
}
