package controllers

import (
	"fmt"
	"github.com/kukkar/common-golang/pkg/utils/queryparser"
	appConf "github.com/kukkar/tigerhall-kittens/conf"
	"github.com/kukkar/tigerhall-kittens/src/globalconst"
	tigerhall "github.com/kukkar/tigerhall-kittens/src/tigerhall-kittens"
	validator "gopkg.in/go-playground/validator.v9"
	"regexp"
	"strings"
)

func parseQuery(q string) (queryparser.QueryParamsList, error) {

	queryParamList, queryParamErr := queryparser.Parse(q)
	if queryParamErr != nil {
		return nil, queryParamErr
	}

	validateParamErr := queryParamList.RemoveInvalid(globalconst.TigerHallQueryMap)
	if validateParamErr != nil {
		return nil, validateParamErr
	}
	return queryParamList, nil
}

//
// Convert The Request Data to tigerhall.Image format.
//
func (r *ReqUploadImage) toImage(conf *appConf.AppConfig) (*tigerhall.Image, error) {

	im := new(tigerhall.Image)

	im.Name = r.getName()
	im.Resource = r.getResource()
	im.Type = r.getType()

	imdata := tigerhall.ImageData{
		Format: r.Image.Format,
		Data:   r.Image.Data,
	}
	ext, err := imdata.PrepareExtension()
	if err != nil {
		return nil, fmt.Errorf("RequestData#toImage()->%s", err.Error())
	}
	im.Extension = ext

	//Set Image Data.
	imCore, err := (imdata).ToCoreImage(im.Extension)

	if err != nil {
		return nil, fmt.Errorf("Could not convert to core Image, Error:%s", err.Error())
	}
	im.Data = *imCore

	return im, nil
}

// Check if the name is supplied, if not generate a random name.
func (r *ReqUploadImage) getResource() string {
	if r.Resource == "" {
		return tigerhall.DEFAULT_RESOURCE
	}
	return strings.ToLower(r.Resource)
}

// Check if the name is supplied, if not generate a random name.
func (r *ReqUploadImage) getName() string {
	//process name
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil || r.Name == "" {
		r.Name = tigerhall.GenerateRandomName()
	}
	//make sure we remove extension from name.
	strArr := strings.Split(r.Name, ".")
	if len(strArr) > 1 {
		r.Name = strings.Join(strArr[0:len(strArr)-1], "-")
	}
	r.Name = reg.ReplaceAllString(r.Name, "")
	return r.Name
}

// Check if the Type is supplied, if not assign default type.
func (r *ReqUploadImage) getType() string {
	if r.Type == "" {
		return tigerhall.DEFAULT_TYPE
	}
	return strings.ToLower(r.Type)
}

//
// validate the Request Object.
//
func (r *ReqUploadImage) validate(validate *validator.Validate) error {
	if !SpecialCharCheck(r.Resource) {
		return fmt.Errorf("Resource does not confronts with rules ")
	}

	if r.Type != "" && !SpecialCharCheck(r.Type) {
		return fmt.Errorf("Type does not confronts with rules")
	}

	// Remove Special char check as we are already removing such chars lateron.
	// if r.Name != "" {
	// 	if err := util.CheckForValidFileName(r.Name); err != nil {
	// 		return fmt.Errorf("Name does not confronts with rules. %s", err.Error())
	// 	}
	// }
	validateErr := validate.Struct(r)
	if validateErr == nil {
		return nil
	}
	var errors []string
	for _, err := range validateErr.(validator.ValidationErrors) {
		if err.Tag() == "required" {
			errors = append(errors, fmt.Sprintf("%s could not be left empty",
				err.Field(),
			))
		} else {
			errors = append(errors, fmt.Sprintf("%s does not confronts with rule: %s",
				err.StructNamespace(), err.Tag(),
			))
		}
	}
	return fmt.Errorf(strings.Join(errors, "|"))
}

func SpecialCharCheck(resource string) bool {

	matched, err := regexp.MatchString(`^[a-zA-Z][a-zA-Z0-9]{0,}$`, resource)
	if !matched || err != nil {
		return false
	}
	return true
}
