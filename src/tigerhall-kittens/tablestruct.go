package tigerhall

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoTigerCollection struct {
	Id                  primitive.ObjectID `bson:"_id,omitempty"`
	Name                string             `bson:"name"`
	DOB                 time.Time          `bson:"dob"`
	LastSeenAt          time.Time          `bson:"lastSeenAt"`
	LastSeenCoordinates Coordinates        `bson:"lastSeenCoordinates"`
	TigerLastSeenSights []MongoTigerSight  `bson:"tigerLastLocations,omitempty"`
}

type MongoTigerCollection4Sight struct {
	Id                  primitive.ObjectID    `bson:"_id,omitempty"`
	Name                string                `bson:"name"`
	DOB                 time.Time             `bson:"dob"`
	LastSeenAt          time.Time             `bson:"lastSeenAt"`
	LastSeenCoordinates MongoTigerCoordinates `bson:"lastSeenCoordinates"`
	TigerLastSeenSights MongoTigerSight       `bson:"tigerLastLocations,omitempty"`
}

type MongoTigerSight struct {
	Coordinates MongoTigerCoordinates `bson:"coordinates"`
	TimeStamp   time.Time             `bson:"timeStamp"`
	ImagePath   string                `bson:"image"`
}

type MongoTigerCoordinates struct {
	Lat  float64 `bson:"lat"`
	Long float64 `bson:"long"`
}
