package v1

import (
	"testing"
)

func TestAcceptance(t *testing.T) {
	t.Run("wrong aspect ratio", func(t *testing.T) {
		err := main("fixtures/wrong_aspect_ratio.tif")
		want := ErrWrongAspectRatio
		if err != want {
			t.Fatalf("got %s want %s", err, want)
		}
	})

	t.Run("right aspect ratio", func(t *testing.T) {
		err := main("fixtures/right_aspect_ratio.tif")
		if err != nil {
			t.Fatal(err)
		}
	})
}
