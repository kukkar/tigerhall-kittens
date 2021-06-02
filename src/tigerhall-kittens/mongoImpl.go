package tigerhall

import (
	"context"
	"fmt"

	"github.com/kukkar/common-golang/pkg/logger"
	"github.com/kukkar/common-golang/pkg/utils/queryparser"
	mfactory "github.com/kukkar/tigerhall-kittens/src/common/factory/mongof"
	"github.com/kukkar/tigerhall-kittens/src/globalconst"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mongoAdapter struct {
	adatper *mfactory.MDB
}

func (this *mongoAdapter) createNewTiger(ctx context.Context,
	req TigerCollection) error {

	dbReq := MongoTigerCollection{
		Name:                req.Name,
		DOB:                 req.DOB,
		LastSeenAt:          req.LastSeenAt,
		LastSeenCoordinates: req.LastSeenCoordinates,
		TigerLastSeenSights: nil,
	}
	dbBytes, err := bson.Marshal(dbReq)
	if err != nil {
		return err
	}
	err = bson.Unmarshal(dbBytes, &dbReq)
	if err != nil {
		return err
	}
	mongErr := this.adatper.Insert(ctx, TigetHallCollection, dbReq)
	if mongErr != nil {
		return mongErr
	}
	return nil
}

func (this *mongoAdapter) addTigerSight(ctx context.Context, id string, req SightData) error {

	where := make(map[string]interface{})
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	where["_id"] = objID

	mongoSightData := MongoTigerSight{
		Coordinates: MongoTigerCoordinates{
			Lat:  req.Coordinates.Lat,
			Long: req.Coordinates.Long,
		},
		ImagePath: req.ImagePath,
		TimeStamp: req.TimeStamp,
	}

	updates := make(map[string]interface{})
	updates["$set"] = map[string]interface{}{"lastSeenCoordinates": mongoSightData.Coordinates}
	updates["$push"] = map[string]interface{}{"tigerLastLocations": mongoSightData}
	updates["$set"] = map[string]interface{}{"lastSeenAt": req.TimeStamp, "lastSeenCoordinates": mongoSightData.Coordinates}

	mdErr := this.adatper.Update(ctx, TigetHallCollection, where, updates)
	if mdErr != nil {
		return fmt.Errorf(mdErr.Error())
	}
	return nil
}

func (this *mongoAdapter) getTigerSights(ctx context.Context, id string,
	sortBy string, sortOrder int, limit, page int) (*TigerCollection, *int, error) {

	rawConn := this.adatper.GetRawConn()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, nil, err
	}
	pipeline := []bson.M{
		bson.M{"$match": bson.M{"_id": objID}},
		bson.M{"$unwind": "$tigerLastLocations"},
		bson.M{"$sort": bson.M{"tigerLastLocations.timeStamp": -1}},
		bson.M{"$skip": page * limit},
		bson.M{"$limit": limit},
	}

	cursor, err := rawConn.Collection(TigetHallCollection).Aggregate(ctx, pipeline)
	if err != nil {
		return nil, nil, err
	}
	var res TigerCollection
	res.TigerLastSeenSights = make([]MongoTigerSight, 0)
	for cursor.Next(ctx) {
		var mongoRecord MongoTigerCollection4Sight
		err := cursor.Decode(&mongoRecord)
		if err != nil {
			return nil, nil, err
		}
		res.TigerLastSeenSights = append(res.TigerLastSeenSights, MongoTigerSight{
			Coordinates: mongoRecord.TigerLastSeenSights.Coordinates,
			ImagePath:   mongoRecord.TigerLastSeenSights.ImagePath,
			TimeStamp:   mongoRecord.TigerLastSeenSights.TimeStamp,
		})
		res.DOB = mongoRecord.DOB
		res.Name = mongoRecord.Name
		res.ID = mongoRecord.Id.Hex()
	}

	data, err := this.getTigerData(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	count := len(data.TigerLastSeenSights)
	return &res, &count, nil
}

func (this *mongoAdapter) getTigerData(ctx context.Context,
	id string) (*TigerCollection, error) {

	whereQuery := make(map[string]interface{})
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	whereQuery["_id"] = objID
	mongoData, mErr := this.adatper.FindOne(ctx, TigetHallCollection, whereQuery)
	if err != nil {
		return nil, fmt.Errorf(mErr.Error())
	}
	bytesData, err := bson.Marshal(mongoData)
	if err != nil {
		return nil, err
	}
	var dbData MongoTigerCollection

	err = bson.Unmarshal(bytesData, &dbData)
	if err != nil {
		return nil, err
	}

	return &TigerCollection{
		ID:                  dbData.Id.Hex(),
		Name:                dbData.Name,
		DOB:                 dbData.DOB,
		LastSeenAt:          dbData.LastSeenAt,
		LastSeenCoordinates: dbData.LastSeenCoordinates,
		TigerLastSeenSights: dbData.TigerLastSeenSights,
	}, nil
}

func (this *mongoAdapter) getTigers(ctx context.Context, q queryparser.QueryParamsList,
	limit, page int, sortBy string, sortOrder string) ([]TigerCollection, int, error) {

	listData := make([]TigerCollection, 0)
	whereQuery, err := q.GetMongoQuery(globalconst.TigerHallQueryMap)
	if err != nil {
		return nil, 0, err
	}
	sortOrderMongo := 1
	if sortOrder == "desc" {
		sortOrderMongo = -1
	}
	logger.Info(fmt.Sprintf("query list %v ", whereQuery))

	mongoData, collCount, mErr := this.adatper.FindSortnLoad(ctx, TigetHallCollection, whereQuery, nil,
		sortBy, sortOrderMongo, page*limit, limit)
	if err != nil {
		return nil, 0, fmt.Errorf(mErr.Error())
	}
	dbData := make([]MongoTigerCollection, 0)

	for _, eachmData := range mongoData {
		var singleRecord MongoTigerCollection
		bytesData, err := bson.Marshal(eachmData)
		if err != nil {
			return nil, 0, err
		}
		err = bson.Unmarshal(bytesData, &singleRecord)
		if err != nil {
			return nil, 0, err
		}
		dbData = append(dbData, singleRecord)

	}
	for _, eachData := range dbData {
		listData = append(listData, TigerCollection{
			ID:                  eachData.Id.Hex(),
			Name:                eachData.Name,
			DOB:                 eachData.DOB,
			LastSeenAt:          eachData.LastSeenAt,
			LastSeenCoordinates: eachData.LastSeenCoordinates,
			TigerLastSeenSights: eachData.TigerLastSeenSights,
		})
	}
	return listData, collCount, nil
}
