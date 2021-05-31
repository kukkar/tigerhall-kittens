package tigerhall

import (
	"encoding/base64"
	"fmt"
	"image"
	"math"

	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/sanksons/gowraps/imaging"
)

func distance(lat1 float64, lng1 float64,
	lat2 float64, lng2 float64, unit ...string) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) +
		math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}

//
// Generate a random name for the image.
//
func GenerateRandomName() string {
	return randomdata.SillyName() + randomdata.StringNumber(1, "")
}

//
// Stores Actual Image content
//
type ImageData struct {
	Data   string
	Format string
}

func (this *ImageData) PrepareExtension() (string, error) {
	databytes, err := this.toBytes()
	if err != nil {
		return "", err
	}
	mime, err := imaging.GetMime(databytes)
	if err != nil {
		return "", err
	}
	ext, err := imaging.GetExtension4mMime(mime)
	if err != nil {
		return "", err
	}
	return ext, nil
}

//
// Convert the supplied data bytes to core image.Image
//
func (this *ImageData) ToCoreImage(extension string) (*image.Image, error) {

	databytes, err := this.toBytes()
	if err != nil {
		return nil, err
	}
	mime, _ := imaging.GetMime4mExt(extension)
	imF, err := imaging.GetCoreImage(databytes, mime)
	if err != nil {
		return nil, err
	}
	return &imF, nil
}

// Convert Image data to bytes.
func (this *ImageData) toBytes() (databytes []byte, err error) {

	switch this.Format {
	case IMAGE_ENCODE_FORMAT:
		databytes, err = base64.StdEncoding.DecodeString(this.Data)
	case IMAGE_ENCODE_FORMAT_PLAIN:
		databytes, err = []byte(this.Data), nil
	default:
		databytes, err = nil, fmt.Errorf("Not a Valid format specified")
	}
	return databytes, err
}
