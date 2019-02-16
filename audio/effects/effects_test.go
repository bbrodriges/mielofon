package effects

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		effect Effect
		text   string
		expect string
	}{
		{Hamster, "squeek", `<speaker effect="hamster">squeek<speaker effect="-">`},
		{BehindTheWall, "psst", `<speaker effect="behind_the_wall">psst<speaker effect="-">`},
		{Megaphone, "heeey", `<speaker effect="megaphone">heeey<speaker effect="-">`},
		{PitchDown, "looool", `<speaker effect="pitch_down">looool<speaker effect="-">`},
		{Psychodelic, "high", `<speaker effect="psychodelic">high<speaker effect="-">`},
		{Pulse, "beep", `<speaker effect="pulse">beep<speaker effect="-">`},
		{TrainAnnounce, "here it comes", `<speaker effect="train_announce">here it comes<speaker effect="-">`},
	}

	for _, tc := range testCases {
		t.Run(string(tc.effect), func(t *testing.T) {
			assert.Equal(t, tc.expect, Add(tc.effect, tc.text))
		})
	}
}
