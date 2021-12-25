package v1

import (
	"bufio"
	"image"
	"os"
	"testing"

	"golang.org/x/image/tiff"
)

func TestSampler(t *testing.T) {
	dir := t.TempDir()

	t.Run("wrong aspect ratio", func(t *testing.T) {
		path := dir + "/wrong_aspect_ratio.tif"

		t.Run("create test heightmap", func(t *testing.T) {
			file, err := os.Create(path)
			if err != nil {
				t.Fatal(err)
			}
			defer file.Close()

			image := image.NewGray16(
				image.Rectangle{
					image.Point{0, 0},
					image.Point{23, 34},
				},
			)

			writer := bufio.NewWriter(file)
			tiff.Encode(writer, image, nil)
			writer.Flush()
		})

		_, got := Sampler(path)
		want := ErrWrongAspectRatio
		if got != want {
			t.Fatalf("got %s want %s", got, want)
		}
	})

	t.Run("right aspect ratio", func(t *testing.T) {
		path := dir + "/right_aspect_ratio.tif"

		t.Run("create test heightmap", func(t *testing.T) {
			file, err := os.Create(path)
			if err != nil {
				t.Fatal(err)
			}
			defer file.Close()

			image := image.NewGray16(
				image.Rectangle{
					image.Point{0, 0},
					image.Point{2847, 942},
				},
			)

			writer := bufio.NewWriter(file)
			tiff.Encode(writer, image, nil)
			writer.Flush()
		})

		sampler, err := Sampler(path)
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
