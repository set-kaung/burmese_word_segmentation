package blitter

import (
	"testing"
)

func checkIfSliceEqual(sl1, sl2 []string) bool {
	if len(sl1) != len(sl2) {
		return false
	}
	for i := 0; i < len(sl1); i++ {
		if sl1[i] != sl2[i] {
			return false
		}
	}
	return true
}

func TestSplitIntoWords(t *testing.T) {
	tests := []struct {
		sentence string
		want     []string
	}{
		{sentence: "လက္ခဏ", want: []string{"လ", "က္ခ", "ဏ"}},
		{sentence: "ဆေးရုံ", want: []string{"ဆေး", "ရုံ"}},
		{sentence: "အနောက်တိုင်းဆေး ", want: []string{"အ", "နောက်", "တိုင်း", "ဆေး"}},
		{sentence: "စမ်းသပ်၊စစ်ဆေးချက်များ", want: []string{"စမ်း", "သပ်", "၊", "စစ်", "ဆေး", "ချက်", "များ"}},
		{sentence: "ရောဂါကိုရှာဖွေရသည်။", want: []string{"ရော", "ဂါ", "ကို", "ရှာ", "ဖွေ", "ရ", "သည်", "။"}},
		{sentence: "ကုသ", want: []string{"ကု", "သ"}},
		{sentence: "ကကက", want: []string{"က", "က", "က"}},
	}

	for _, tt := range tests {
		got := CreateWordsSlice(Splitter(tt.sentence))
		if !checkIfSliceEqual(got, tt.want) {
			t.Errorf("want: %+v, got %+v for %s", tt.want, got, tt.sentence)
		}
	}
}
