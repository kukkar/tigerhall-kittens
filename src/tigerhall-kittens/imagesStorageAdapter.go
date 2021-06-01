package tigerhall

type ImageStorageAdapter interface {
	// Save Image data to storage.
	CreateImage(im *Image) error
	// Get Image data from storage.
	GetImage(im *Image) ([]byte, error)
	// create variations of image
	CreateVariation(v *Variation) error
	//get variations of image
	GetVariation(v *Variation) ([]byte, error)
}
