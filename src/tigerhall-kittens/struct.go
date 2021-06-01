package tigerhall

import (
	"image"
	"time"
)

type ReqCreateTiger struct {
	Name        string
	DOB         time.Time
	SeenAt      time.Time
	Coordinates Coordinates
}

type Coordinates struct {
	Lat  float64
	Long float64
}

type ResListTiger struct {
	ID                  string
	Name                string
	DOB                 time.Time
	SeenAt              time.Time
	LastSeenCoordinates Coordinates
}

type ReqSightOfATiger struct {
	Coordinates Coordinates
	TimeStamp   time.Time
	ImagePath   string
}

type ResListSigntsOfTiger struct {
	Name        string
	DOB         time.Time
	TigerSights []SightData
}

type SightData struct {
	Coordinates Coordinates
	TimeStamp   time.Time
	ImagePath   string
}

type ConfigTigerHall struct {
	StorageAdapter string
}

type TigerCollection struct {
	Name                string
	DOB                 time.Time
	LastSeenAt          time.Time
	LastSeenCoordinates Coordinates
	TigerLastSeenSights []MongoTigerSight
}

type TigerSight struct {
	Coordinates Coordinates
	TimeStamp   time.Time
	ImagePath   string
}

type igerCoordinates struct {
	Lat  float64
	Long float64
}

//
// Image struct
// This acts as the backbone of Image processing.
//
type Image struct {
	Name           string
	Extension      string
	Resource       string
	Type           string
	Status         string
	Variations     []Variation
	Data           image.Image
	storageAdapter ImageStorageAdapter
}
