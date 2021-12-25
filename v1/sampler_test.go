package v1

import (
	"bufio"
	"image"
	"image/color"
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

			image.SetGray16(0, 0, color.Gray16{1})
			image.SetGray16(0, 180, color.Gray16{8})

			writer := bufio.NewWriter(file)
			tiff.Encode(writer, image, nil)
			writer.Flush()
		})

		sampler, err := Sampler(path)
		if err != nil {
			t.Fatal(err)
		}

		AssertUInt16(
			t,
			sampler(0, 0),
			1,
		)

		AssertUInt16(
			t,
			sampler(0, 180),
			8,
		)
	})
}

func AssertUInt16(t *testing.T, got, want uint16) {
	t.Helper()
	if got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}
