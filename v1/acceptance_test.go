package v1

import (
	"testing"
)

func TestAcceptance(t *testing.T) {
	err := main("wrong_aspect_ratio.tif")
	want := ErrWrongAspectRatio
	if err != want {
		t.Fatalf("got %s want %s", err, want)
	}
}
