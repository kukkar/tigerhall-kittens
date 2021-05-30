package tigerhall

import "time"

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
	ID          string
	Name        string
	DOB         time.Time
	SeenAt      time.Time
	Coordinates Coordinates
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
}

type TigerCollection struct {
	UUID                string
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
