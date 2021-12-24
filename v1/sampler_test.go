package v1

import (
	"testing"
)

func TestSampler(t *testing.T) {
	t.Run("wrong aspect ratio", func(t *testing.T) {
		_, err := Sampler("fixtures/wrong_aspect_ratio.tif")
		want := ErrWrongAspectRatio
		if err != want {
			t.Fatalf("got %s want %s", err, want)
		}
	})

	t.Run("right aspect ratio", func(t *testing.T) {
		sampler, err := Sampler("fixtures/right_aspect_ratio.tif")
		if err != nil {
			t.Fatal(err)
		}
		got := sampler(0, 0)
		want := 0.0
		if got != want {
			t.Fatalf("got %f want %f", got, want)
		}
	})
}
