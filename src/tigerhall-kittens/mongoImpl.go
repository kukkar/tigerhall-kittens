package tigerhall

import (
	"context"
	"fmt"

	"github.com/kukkar/common-golang/pkg/utils/queryparser"
	mfactory "github.com/kukkar/tigerhall-kittens/src/common/factory/mongof"
	"go.mongodb.org/mongo-driver/bson"
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

func (this *mongoAdapter) getTigers(ctx context.Context, q queryparser.QueryParamsList,
	limit, offset int) ([]TigerCollection, error) {
	return nil, fmt.Errorf("todo")
}
