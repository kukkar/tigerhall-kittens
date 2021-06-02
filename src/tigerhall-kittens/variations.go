package tigerhall

import (
	"fmt"
	"image"

	"github.com/nfnt/resize"
	"github.com/sanksons/gowraps/imaging"
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

//
// Converts databytes to core image.Image
//
func (this *Variation) GetDataBytes() []byte {
	if this.Data == nil {
		return nil
	}
	mime, _ := imaging.GetMime4mExt(this.Extension)
	databytes, e := imaging.GetBytes4mImage(this.Data, mime)
	if e != nil {
		fmt.Printf(e.Error())
	}
	return databytes
}

//
// Tailor the variation and store using adapter.
// update in mongo
//
func (this *Variation) Tailor() error {

	//Chop Image to correct size
	this.Data = this.tailor(this.Size, this.Image.Data)

	if this.Image == nil {
		return fmt.Errorf("unable to get image from instance")
	}
	//Store Variation
	err := this.Image.storageAdapter.CreateVariation(this)
	if err != nil {
		return fmt.Errorf("Could not tailor variation [%s], Error: %s", this.Size.ToText(), err.Error())
	}
	return nil
}

// Chop variation to needed size.
func (this *Variation) tailor(size Size, im image.Image) image.Image {
	if im == nil {
		return nil
	}
	return resize.Thumbnail(
		uint(size.Width), uint(size.Height), im, resize.Lanczos3,
	)
}

//
// Fetch data bytes for the variation from storage.
//
func (this *Variation) Fetch() ([]byte, error) {
	return this.Image.storageAdapter.GetVariation(this)
}
