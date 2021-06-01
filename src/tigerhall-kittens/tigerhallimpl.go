package tigerhall

import (
	"context"
	"fmt"

	"github.com/kukkar/common-golang/pkg/logger"
	"github.com/kukkar/common-golang/pkg/utils/queryparser"
	"github.com/sanksons/gowraps/imaging"
)

type tigherhall struct {
	stAdapter      storageAdapter
	imageStAdapter ImageStorageAdapter
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
	limit, offset int) ([]ResListTiger, int, error) {
	ctx := context.TODO()
	res := make([]ResListTiger, 0)
	dbData, count, err := this.stAdapter.getTigers(ctx, q, limit, offset, "", "asc")
	if err != nil {
		return nil, 0, err
	}
	for _, eachDbData := range dbData {
		res = append(res, ResListTiger{
			ID:                  eachDbData.ID,
			Name:                eachDbData.Name,
			DOB:                 eachDbData.DOB,
			SeenAt:              eachDbData.LastSeenAt,
			LastSeenCoordinates: eachDbData.LastSeenCoordinates,
		})
	}
	logger.Info(fmt.Sprintf("dbData %v count %d", dbData, count))
	return res, count, nil
}

//
// SightATiger Register a particular sight of a tiger
//
func (this *tigherhall) SightATiger(tigerID string, req ReqSightOfATiger) error {

	tigerData, err := this.stAdapter.getTigerData(context.Background(), tigerID)
	if err != nil {
		return err
	}

	dist := distance(tigerData.LastSeenCoordinates.Lat, tigerData.LastSeenCoordinates.Long,
		req.Coordinates.Lat, req.Coordinates.Long, DISTANCE_IN_KILOMETER)

	if dist <= VALIDATE_KILOMETER_ADD_SIGHT {
		return ErrKilometerValidation
	}

	return this.stAdapter.addTigerSight(context.Background(), tigerID, SightData{
		ImagePath:   req.ImagePath,
		TimeStamp:   req.TimeStamp,
		Coordinates: req.Coordinates,
	})
}

//
// ListSigntsOfTiger List All Signts of a tiger
//
func (this *tigherhall) ListSigntsOfTiger(id string,
	limit, offset int) (*ResListSigntsOfTiger, error) {

	var res ResListSigntsOfTiger
	tigerData, err := this.stAdapter.getTigerData(context.Background(), id)
	if err != nil {
		return nil, err
	}
	res.Name = tigerData.Name
	res.DOB = tigerData.DOB
	res.ID = tigerData.ID

	for _, eachSight := range tigerData.TigerLastSeenSights {
		res.TigerSights = append(res.TigerSights, SightData{

			Coordinates: Coordinates{
				Lat:  eachSight.Coordinates.Lat,
				Long: eachSight.Coordinates.Long,
			},
			TimeStamp: eachSight.TimeStamp,
			ImagePath: eachSight.ImagePath,
		})
	}
	return &res, nil
}

func (this *tigherhall) CreateImage(im *Image) error {
	if im.Variations == nil {
		im.Variations = []Variation{{
			Size: Size{
				Width:  DEFAULT_VARIATION_WIDTH,
				Height: DEFAULT_VARIATION_HEIGHT,
			},
			Data:  im.Data,
			Image: im,
		},
		}
	}
	for _, eachVariation := range im.Variations {
		err := eachVariation.Tailor()
		if err != nil {
			return err
		}
	}
	return this.imageStAdapter.CreateImage(im)
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
