package data_test

import (
	"testing"

	"github.com/Tesohh/xlearn/data"
)

func TestTagify(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{"giura fra", "giura-fra"},
		{"ZesTY MILLIANNNN__ewwfj", "zesty-milliannnn--ewwfj"},
		{"ken carson 67", "ken-carson-67"},
		{"brodie _ Cast", "brodie---cast"},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			got := data.Tagify(tC.input, false)
			if got != tC.want {
				t.Fatalf("want: %v, got: %v", tC.want, got)
			}
		})
	}
}

func TestHexString(t *testing.T) {
	t.Run("check if hexstring is always of same length", func(t *testing.T) {
		for i := 0; i < 100000; i++ {
			got := data.HexString()
			if len(got) != len("000000") {
				t.Fatalf("len(got) == %v at iteration #%v", len(got), i)
			}
		}
	})
}
