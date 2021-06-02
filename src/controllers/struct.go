package controllers

import (
	"time"
)

//@TODO
type ResListTigers struct {
	TigerData  []TigerData `json:"tigers"`
	TotalCount int         `json:"totalCount"`
}

type TigerData struct {
	ID     string    `json:"id"`
	Name   string    `json:"name"`
	DOB    time.Time `json:"dob"`
	SeenAt time.Time `json:"seenAt"`
}

//@TODO
type ResSightATiger struct {
	TotalCount int            `json:"totalCount"`
	Data       []ResSightData `json:"tigerSights"`
}

type ResSightData struct {
	Coordinates Coordinates `json:"coordinates"`
	ImagePath   string      `json:"imagePath"`
	SeenAt      time.Time   `json:"seenAt"`
}

type ReqUploadImage struct {
	Resource string    `json:"resource"`
	Type     string    `json:"type"`
	Name     string    `json:"name"`
	Image    ImageInfo `json:"image"`
}

type ImageInfo struct {
	Format string `json:"format"`
	Data   string `json:"data"`
}

type ReqCreateTiger struct {
	Name        string      `json:"name"`
	DOB         time.Time   `json:"dob"`
	Coordinates Coordinates `json:"coordinates"`
}

type Coordinates struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type ReqSightATiger struct {
	TigerID     string      `json:"tigerID"`
	Coordinates Coordinates `json:"coordinates"`
	//SeenAt      time.Time   `json:"seenAt"`
	ImagePath string `json:"imagePath"`
}

type ResUploadImage struct {
	ImagePath string `json:"imagePath"`
}
