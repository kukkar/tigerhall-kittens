package tigerhall

import (
	"context"
	"fmt"

	"github.com/kukkar/common-golang/pkg/utils/queryparser"
	"github.com/sanksons/gowraps/imaging"
)

type tigherhall struct {
	stAdapter storageAdapter
}

//
// CreateTiger Register a new Tiger
//
func (this *tigherhall) CreateTiger(req ReqCreateTiger) error {

	ctx := context.TODO()
	err := this.stAdapter.createNewTiger(ctx, TigerCollection{
		Name:                req.Name,
		DOB:                 req.DOB,
		LastSeenAt:          req.SeenAt,
		LastSeenCoordinates: req.Coordinates,
	})
	if err != nil {
		return ErrDBExecution
	}
	return nil
}

//
// ListTigers List all tigers registered previously on query basis
//
func (this *tigherhall) ListTigers(q queryparser.QueryParamsList,
	limit, offset int) ([]ResListTiger, error) {

	return nil, fmt.Errorf("todo")
}

//
// SightATiger Register a particular sight of a tiger
//
func (this *tigherhall) SightATiger(req ReqSightOfATiger) (*int, error) {
	return nil, fmt.Errorf("todo")
}

//
// ListSigntsOfTiger List All Signts of a tiger
//
func (this *tigherhall) ListSigntsOfTiger(q queryparser.QueryParamsList,
	limit, offset int) (*ResListSigntsOfTiger, error) {

	return nil, fmt.Errorf("todo")
}

//
// Get the Raw Image Bytes based on the conversion defined.
//
func (this *Image) GetDataBytes() []byte {

	mime, _ := imaging.GetMime4mExt(this.Extension)
	databytes, _ := imaging.GetBytes4mImage(this.Data, mime)
	return databytes
}

//
// Fetch gets the image data in []byte
//
func (this *Image) Fetch() ([]byte, error) {
	if !this.IsStorageAdapterDefined() {
		return nil, ErrStorageAdapterNotSupplied
	}
	return this.storageAdapter.GetImage(this)
}

//
// Check if storage Adapter Defined.
//
func (im *Image) IsStorageAdapterDefined() bool {
	if im.storageAdapter != nil {
		return true
	}
	return false
}

//
// get Image Name.
//
func (this *Image) GetName(withExtension bool) string {

	if this.Name == "" {
		this.Name = GenerateRandomName()
	}
	//Check if Extension needs to be appended.
	if !withExtension {
		return this.Name
	} else {
		return fmt.Sprintf("%s.%s", this.Name, this.Extension)
	}
}

//
//get Image Resource.
//
func (this *Image) GetResource() string {
	if this.Resource == DEFAULT_RESOURCE {
		return ""
	}
	// fmt.Println("Image Resource : ", this.Resource)
	return this.Resource
}
