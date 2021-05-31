package imaging

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"strings"
)

//
// Mime Types
//
const MIME_TYPE_JPEG = "image/jpeg"
const MIME_TYPE_PNG = "image/png"
const MIME_TYPE_GIF = "image/gif"

//
// Extensions
//
const EXT_JPEG = "jpeg"
const EXT_JPG = "jpg"
const EXT_PNG = "png"
const EXT_GIF = "gif"

var byteMimeTable = map[string]string{
	"\xff\xd8\xff":      MIME_TYPE_JPEG,
	"\x89PNG\r\n\x1a\n": MIME_TYPE_PNG,
	"GIF87a":            MIME_TYPE_GIF,
	"GIF89a":            MIME_TYPE_GIF,
}

var mimeExtTable = map[string]string{
	MIME_TYPE_JPEG: EXT_JPEG,
	MIME_TYPE_PNG:  EXT_PNG,
	MIME_TYPE_GIF:  MIME_TYPE_GIF,
}

var extMIMETable = map[string]string{
	EXT_JPEG: MIME_TYPE_JPEG,
	EXT_JPG:  MIME_TYPE_JPEG,
	EXT_PNG:  MIME_TYPE_PNG,
	EXT_GIF:  MIME_TYPE_GIF,
}

func GetMime(data []byte) (string, error) {
	dataStr := string(data)

	for raw, mime := range byteMimeTable {
		if strings.HasPrefix(dataStr, raw) {
			return mime, nil
		}
	}
	return "", fmt.Errorf("Could not deduce mime of image.")
}

func GetExtension4mMime(mime string) (string, error) {
	if val, ok := mimeExtTable[mime]; ok {
		return val, nil
	}
	return "", fmt.Errorf("Not a Valid Mime Type")
}

func GetMime4mExt(ext string) (string, error) {
	if val, ok := extMIMETable[ext]; ok {
		return val, nil
	}
	return "", fmt.Errorf("Not a Valid Mime Type")
}

func GetBytes4mImage(im image.Image, mime string) ([]byte, error) {

	if im == nil {
		return nil, nil
	}
	switch mime {
	case MIME_TYPE_PNG:
		return ImageToBytesPng(im)
	case MIME_TYPE_JPEG:
		return ImageToBytesJpeg(im)
	default:
		return nil, fmt.Errorf("This Image extension is not supported.")
	}
}

func GetCoreImage(dataBytes []byte, mime string) (image.Image, error) {

	if len(dataBytes) == 0 {
		return nil, fmt.Errorf("GetCoreImage()->Empty DataBytes supplied")
	}
	var imF image.Image
	var err error

	switch mime {
	case MIME_TYPE_PNG:
		imF, err = PngToImage(dataBytes)
	case MIME_TYPE_JPEG:
		imF, err = JpegToImage(dataBytes)

	default:
		return nil, fmt.Errorf("GetCoreImage()->This Image Mime is not supported.")
	}
	return imF, err
}

func PngToImage(dataBytes []byte) (image.Image, error) {
	return png.Decode(bytes.NewReader(dataBytes))
}

func ImageToBytesPng(im image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, im)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func JpegToImage(dataBytes []byte) (image.Image, error) {
	return jpeg.Decode(bytes.NewReader(dataBytes))
}

func ImageToBytesJpeg(im image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, im, nil)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
