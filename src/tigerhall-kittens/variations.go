package tigerhall

import (
	"fmt"
	"image"
)

// A image with a specific size denotes variation.
type Variation struct {
	Size      Size
	Extension string
	Data      image.Image
	Image     *Image
}

//
// Stores Pixel Size of variation.
//
type Size struct {
	Width  int
	Height int
}

//
// Convert To String format.
//
func (s *Size) ToText() string {
	return fmt.Sprintf("%dx%d", s.Width, s.Height)
}
