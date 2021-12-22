package v1

import (
	"errors"
)

func main(file string) error {
	return ErrWrongAspectRatio
}

var ErrWrongAspectRatio = errors.New("aspect ratio must be 284.7 x 94.2")
