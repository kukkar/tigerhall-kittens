package tigerhall

import (
	"fmt"
	"os"
	"strings"

	"github.com/sanksons/gowraps/filesystem"
)

//
// Initialize Local Storage Adapter.
//
func InitializeLocalStorage(conf LocalStorageConf) (*LocalStorage, error) {
	storage := new(LocalStorage)
	storage.Path = conf.Path
	return storage, nil
}

//
// Configuration for Local Storage.
//
type LocalStorageConf struct {
	Path string
}

//
// Local Storage Adapter.
//
type LocalStorage struct {
	// Base path of images.
	Path string
}

//
// Implementation of CreateImage() function
//
func (this *LocalStorage) CreateImage(im *Image) error {

	imagePath := this.getStoragePath(im.Resource, im.Type, ORIGINAL_IMAGE)
	// Create directory if required.
	if err := this.createDirectory(imagePath); err != nil {
		return err
	}
	fullImagePath := strings.Join([]string{imagePath, im.GetName(true)}, DS)
	saveerr := this.save(fullImagePath, im.GetDataBytes(), true)
	if saveerr != nil {
		return saveerr
	}
	return nil
}

//
// Implementation of GetImage() function
//
func (this *LocalStorage) GetImage(im *Image) ([]byte, error) {
	imagePath := this.getStoragePath(im.Resource, im.Type, ORIGINAL_IMAGE)
	fullImagePath := strings.Join([]string{imagePath, im.GetName(true)}, DS)
	dataBytes, err := this.getFile(fullImagePath)
	if err != nil {
		return nil, err
	}

	if err == nil && dataBytes == nil {
		return nil, ErrFileNotFound
	}
	return dataBytes, nil
}

// get storage path for image or variation.
func (this *LocalStorage) getStoragePath(resource, rtype, size string) string {
	var imagepath string
	imagepath = strings.Join([]string{resource, rtype, size}, DS)

	sliceStr := []string{this.Path, imagepath}
	return strings.Join(sliceStr, DS)
}

// get file data.
func (this *LocalStorage) getFile(filepath string) ([]byte, error) {
	databytes, err := filesystem.GetFile(filepath)
	if err == nil {
		return databytes, nil
	}
	switch err {
	case filesystem.ErrFileNotFound:
		return nil, ErrFileNotFound
	case filesystem.ErrPermissionDenied:
		return nil, ErrPermissionDenied
	}
	return nil, err
}

// create directory if required.
func (this *LocalStorage) createDirectory(imagePath string) error {
	return filesystem.CreateDirTree(imagePath, 0775)
}

// Save Image based on supplied path and data.
func (this *LocalStorage) save(imageFilePath string, data []byte, override bool) error {
	if !override {
		//check if a file with name already exists.
		//if so its a error.
		alreadyExists, err := this.checkIfFileExists(imageFilePath)
		if err != nil {
			return err
		}
		if alreadyExists {
			return fmt.Errorf("A file with the name already exists")
		}
	}
	file, err := os.Create(imageFilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// check if file already exists.
func (this *LocalStorage) checkIfFileExists(filepath string) (bool, error) {
	return filesystem.CheckIfFileExists(filepath)
}

//
// Implementation of CreateVariation() function
//
func (this *LocalStorage) CreateVariation(v *Variation) error {

	im := v.Image
	imagePath := this.getStoragePath(v.Image.Resource, v.Image.Type, v.Size.ToText())
	// Create directory if required.
	if err := this.createDirectory(imagePath); err != nil {
		return err
	}
	fullImagePath := strings.Join([]string{imagePath, im.GetName(true)}, DS)
	saveerr := this.save(fullImagePath, v.GetDataBytes(), true)
	if saveerr != nil {
		return saveerr
	}
	return nil
}

//
// Implementation of GetVariation() function
//
func (this *LocalStorage) GetVariation(v *Variation) ([]byte, error) {
	imagePath := this.getStoragePath(v.Image.Resource, v.Image.Type, v.Size.ToText())
	fullImagePath := strings.Join([]string{imagePath, v.Image.GetName(true)}, DS)
	dataBytes, err := this.getFile(fullImagePath)
	if err != nil {
		return nil, err
	}
	if err == nil && dataBytes == nil {
		return nil, ErrFileNotFound
	}
	return dataBytes, nil
}
