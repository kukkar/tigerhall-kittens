package mongodb

import (
	"context"
	"fmt"

	"go.elastic.co/apm/module/apmmongo"
	"go.mongodb.org/mongo-driver/bson"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongodb driver
type MongoDriver struct {
	conn *mgo.Database
	//session *mgo.Session
	conf *MDBConfig
}

// init method
func (obj *MongoDriver) Init(conf *MDBConfig) *MDBError {
	// set the connection
	tmp, err := mgo.Connect(context.Background(), options.Client().ApplyURI(conf.URL).SetMonitor(apmmongo.CommandMonitor())) //Dial(conf.URL)
	if err != nil {
		return getErrObj(ErrInitialization, err.Error()+"-connection url:"+conf.URL)
	}
	//obj.session = tmp
	obj.conn = tmp.Database(conf.DbName)
	return nil
}

// func (obj *MongoDriver) Copy(safe *Safe) *MSession {
// 	ms := new(MSession)
// 	ms.mgoSession = obj.session.Copy()
// 	if safe != nil {
// 		ms.mgoSession.SetSafe(obj.getMSafeFrom(safe))
// 	}
// 	return ms
// }

// func (obj *MongoDriver) getMSafeFrom(safe *Safe) *mgo.Safe {
// 	mgoSafe := new(mgo.Safe)
// 	mgoSafe.FSync = safe.FSync
// 	mgoSafe.J = safe.J
// 	mgoSafe.W = safe.W
// 	mgoSafe.WMode = safe.WMode
// 	mgoSafe.WTimeout = safe.WTimeout
// 	return mgoSafe
// }

// // Close shuts down the current session.
// func (obj *MongoDriver) Close(session *MSession) {
// 	if session == nil {
// 		return
// 	}
// 	session.mgoSession.Close()
// }

// // CloseMasterSession closes the master session
// func (obj *MongoDriver) CloseMasterSession() {
// 	obj.session.Close()
// }

func (obj *MongoDriver) GetNewID(collection, idName, counterName string) (interface{}, *MDBError) {
	//	obj.session.Refresh()
	// result, retErr := obj.getNewID(obj.conn.Collectionollection(collection), idName, counterName, 1)
	// if retErr != nil {
	// 	return nil, retErr
	// }
	// return bson.M{counterName: result[0]}, nil
	return nil, nil
}

// func (obj *MongoDriver) getNewID(collection *mgo.Collection, idName, counterName string, incValue int) ([]int, *MDBError) {
// 	changeInDocument := mgo.UpdateManyModel {
// 		Update:    bson.M{"$inc": bson.M{counterName: incValue}},
// 		ReturnNew: true,
// 	}
// 	var result bson.M
// 	if _, err := collection.Find(bson.M{"_id": idName}).Apply(changeInDocument, &result); err != nil {

// 		if err.Error() == "not found" {
// 			var value = map[string]interface{}{
// 				"_id":       idName,
// 				counterName: incValue,
// 			}
// 			errM := obj.insert(collection, value)
// 			if errM != nil {
// 				return nil, errM
// 			}

// 			counterValue := value[counterName].(int)
// 			counters := []int{}

// 			for i := 1; i <= counterValue; i++ {
// 				counters = append(counters, i)
// 			}
// 			return counters, nil
// 		}

// 		return nil, getErrObj(ErrFindOneFailure, err.Error())
// 	}

// 	// counter in map of interface
// 	counterValue := result[counterName].(int)
// 	counters := []int{}

// 	for i := counterValue; i > (counterValue - incValue); i-- {
// 		counters = append(counters, i)
// 	}

// 	return counters, nil
// }

//GetNewCustomID is used to increment ID by given incValue and returns array of it
func (obj *MongoDriver) GetNewCustomID(collection, idName, counterName string, incValue int) ([]int, *MDBError) {
	// //	obj.session.Refresh()
	// result, retErr := obj.getNewID(obj.conn.Collection(collection), idName, counterName, incValue)
	// if retErr != nil {
	// 	return nil, retErr
	// }
	return nil, nil
}

// // FindOne queries the mongo DB using the session and returns only single result/collection
// func (obj *MongoDriver) FindOneUsingSession(session *MSession, collection string, query map[string]interface{}) (ret interface{}, aerr *MDBError) {
// 	sess := session.mgoSession
// 	return obj.findOne(sess.DB(obj.conf.DbName).C(collection), bson.M(query))
// }

// FindOne queries the mongo DB and returns only single result/collection
func (obj *MongoDriver) FindOne(ctx context.Context, collection string, query map[string]interface{}) (ret interface{}, aerr *MDBError) {
	//	obj.session.Refresh()
	return obj.findOne(ctx, obj.conn.Collection(collection), bson.M(query))
}

func (obj *MongoDriver) FindOnenLoad(coll string, q map[string]interface{}, payload interface{}) *MDBError {
	//	obj.session.Refresh()
	// if err := obj.conn.Collection(coll).Find(q).One(payload); err != nil {
	// 	return getErrObj(ErrFindOneFailure, err.Error())
	// }
	return nil
}

func (obj *MongoDriver) findOne(ctx context.Context, collection *mgo.Collection, query bson.M) (ret interface{}, aerr *MDBError) {
	if result := collection.FindOne(ctx, query); result.Err() != nil {
		return nil, getErrObj(ErrFindOneFailure, result.Err().Error())
	}
	return ret, nil
}

// // FindAll queries the mongo DB using session and returns all the results
// func (obj *MongoDriver) FindAllUsingSession(session *MSession, collection string, query map[string]interface{}) (ret []interface{}, aerr *MDBError) {
// 	sess := session.mgoSession
// 	return obj.findAll(sess.DB(obj.conf.DbName).C(collection), bson.M(query))
// }

// // FindAll queries the mongo DB and returns all the results
// func (obj *MongoDriver) FindAll(collection string, query map[string]interface{}) (ret []interface{}, aerr *MDBError) {
// 	//	obj.session.Refresh()
// 	return obj.findAll(obj.conn.Collection(collection), bson.M(query))
// }

// func (obj *MongoDriver) findAll(collection *mgo.Collection, query bson.M) (ret []interface{}, aerr *MDBError) {
// 	if err := collection.Find(query).All(&ret); err != nil {
// 		return nil, getErrObj(ErrFindAllFailure, err.Error())
// 	}
// 	return ret, nil
// }

// // FindAll queries the mongo DB and returns the results with limit
// func (obj *MongoDriver) FindWithLimit(collection string, query map[string]interface{}, skip int, limit int) (ret []interface{}, aerr *MDBError) {
// 	//	obj.session.Refresh()
// 	return obj.findWithLimit(obj.conn.Collection(collection), bson.M(query), skip, limit)
// }

// func (obj *MongoDriver) findWithLimit(collection *mgo.Collection, query bson.M, skip int, limit int) (ret []interface{}, aerr *MDBError) {
// 	if err := collection.Find(query).Skip(skip).Limit(limit).All(&ret); err != nil {
// 		return nil, getErrObj(ErrFindAllFailure, err.Error())
// 	}
// 	return ret, nil
// }

// // Insert
// func (obj *MongoDriver) InsertUsingSession(session *MSession, collection string, value interface{}) *MDBError {
// 	sess := session.mgoSession
// 	return obj.insert(sess.DB(obj.conf.DbName).C(collection), value)
// }

func (obj *MongoDriver) Insert(ctx context.Context, collection string, value interface{}) *MDBError {
	//	obj.session.Refresh()
	return obj.insert(ctx, obj.conn.Collection(collection), value)
}

func (obj *MongoDriver) insert(ctx context.Context, collection *mgo.Collection, value interface{}) *MDBError {
	result, err := collection.InsertOne(ctx, value, nil)
	if err != nil {
		return getErrObj(ErrDuplicateKey, err.Error())
	}
	fmt.Println(result)
	// 	//check if its a duplicate key error.
	// if mgo.IsDup(err) {
	// 	return getErrObj(ErrDuplicateKey, err.Error())
	// } else {
	// 	return getErrObj(ErrInsertFailure, err.Error())
	// }

	return nil
}

func (obj *MongoDriver) BulkInsert(ctx context.Context, collection string, value []interface{}) *MDBError {
	//	obj.session.Refresh()
	return obj.insertMany(ctx, obj.conn.Collection(collection), value)
}

func (obj *MongoDriver) insertMany(ctx context.Context, collection *mgo.Collection, value []interface{}) *MDBError {
	result, err := collection.InsertMany(ctx, value, nil)
	if err != nil {
		return getErrObj(ErrInsertFailure, err.Error())
	}
	fmt.Println(result)

	return nil
}

// // Update updates the mongo DB collection passed as an argument
// func (obj *MongoDriver) UpdateUsingSession(session *MSession, collection string, query map[string]interface{}, value interface{}) *MDBError {
// 	sess := session.mgoSession
// 	return obj.update(sess.DB(obj.conf.DbName).C(collection), bson.M(query), value)
// }

// // Update updates the mongo DB collection passed as an argument
// func (obj *MongoDriver) Update(collection string, query map[string]interface{}, value interface{}) *MDBError {
// 	//	obj.session.Refresh()
// 	return obj.update(obj.conn.Collection(collection), bson.M(query), value)
// }

// func (obj *MongoDriver) update(collection *mgo.Collection, query bson.M, value interface{}) *MDBError {
// 	if err := collection.Update(query, value); err != nil {
// 		return getErrObj(ErrUpdateFailure, err.Error())
// 	}
// 	return nil
// }

// // Update updates the mongo DB collection passed as an argument
// func (obj *MongoDriver) UpdateAll(collection string, query map[string]interface{}, value interface{}) (*mgo.ChangeInfo, *MDBError) {
// 	//	obj.session.Refresh()
// 	return obj.updateAll(obj.conn.Collection(collection), bson.M(query), value)
// }

// func (obj *MongoDriver) updateAll(collection *mgo.Collection, query bson.M, value interface{}) (*mgo.ChangeInfo, *MDBError) {
// 	var change *mgo.ChangeInfo
// 	change, err := collection.UpdateAll(query, value)
// 	if err != nil {
// 		return change, getErrObj(ErrUpdateFailure, err.Error())
// 	}
// 	return change, nil
// }

// // Upsert updates/inserts the mongo DB collection passed as an argument
// func (obj *MongoDriver) Upsert(collection string, query map[string]interface{}, value interface{}) *MDBError {
// 	//	obj.session.Refresh()
// 	return obj.upsert(obj.conn.Collection(collection), bson.M(query), value)
// }

// // // UpsertUsingSession updates/inserts the mongo DB collection passed as an argument using the session passed as argument
// // func (obj *MongoDriver) UpsertUsingSession(session *MSession, collection string, query map[string]interface{}, value interface{}) *MDBError {
// // 	sess := session.mgoSession
// // 	return obj.upsert(sess.DB(obj.conf.DbName).C(collection), bson.M(query), value)
// // }

// func (obj *MongoDriver) upsert(collection *mgo.Collection, query bson.M, value interface{}) *MDBError {
// 	if _, err := collection.Upsert(query, value); err != nil {
// 		return getErrObj(ErrUpdateFailure, err.Error())
// 	}
// 	return nil
// }

// // BulkUpsert updates/inserts the bulk of data mongo DB collection passed as an argument
// func (obj *MongoDriver) BulkUpsert(collection string, query []map[string]interface{}, value []interface{}) (BulkResult, *MDBError) {
// 	if len(query) != len(value) {
// 		return BulkResult{}, getErrObj(ErrUpsertFailure, "Bulk.upsert requires an even number of parameters")
// 	}
// 	//	obj.session.Refresh()
// 	return obj.bulkUpsert(obj.conn.Collection(collection), query, value)
// }

// func (obj *MongoDriver) bulkUpsert(collection *mgo.Collection, query []map[string]interface{}, value []interface{}) (BulkResult, *MDBError) {
// 	bulk := collection.Bulk()
// 	for i := range query {
// 		bulk.Upsert(query[i], value[i])
// 	}
// 	var err error
// 	bulkResultmgo, err := bulk.Run()
// 	if err != nil {
// 		return BulkResult{}, getErrObj(ErrUpdateFailure, err.Error())
// 	}
// 	bulkResult := BulkResult{
// 		Matched:  bulkResultmgo.Matched,
// 		Modified: bulkResultmgo.Modified,
// 	}
// 	if err != nil {
// 		errorResult, assrFailed := err.(*mgo.BulkError)
// 		if assrFailed == false {
// 			return BulkResult{}, getErrObj(ErrUpsertFailure, "Unable to assert Error array")
// 		}
// 		errorCases := errorResult.Cases()
// 		bulkResult.Ecases = make([]BulkErrorCase, 0)
// 		for _, eachError := range errorCases {
// 			bulkResult.Ecases = append(bulkResult.Ecases, BulkErrorCase{
// 				Index: eachError.Index,
// 				Err:   eachError.Err,
// 			})
// 		}
// 		return bulkResult, getErrObj(ErrUpsertFailure, err.Error())
// 	}
// 	return bulkResult, nil
// }

// // // Remove deletes the documents using session from the collection passed in the argument
// // func (obj *MongoDriver) RemoveUsingSession(session *MSession, collection string, query map[string]interface{}) *MDBError {
// // 	sess := session.mgoSession
// // 	return obj.remove(sess.DB(obj.conf.DbName).C(collection), bson.M(query))
// // }

// // Remove deletes the documents from the collection passed in the argument
// func (obj *MongoDriver) Remove(collection string, query map[string]interface{}) *MDBError {
// 	//	obj.session.Refresh()
// 	return obj.remove(obj.conn.Collection(collection), bson.M(query))
// }

// func (obj *MongoDriver) remove(collection *mgo.Collection, query bson.M) *MDBError {
// 	if err := collection.Remove(query); err != nil {
// 		return getErrObj(ErrRemoveFailure, err.Error())
// 	}
// 	return nil
// }

// // Remove deletes the documents from the collection passed in the argument
// func (obj *MongoDriver) RemoveAll(collection string, query map[string]interface{}) (*mgo.ChangeInfo, *MDBError) {
// 	//	obj.session.Refresh()
// 	return obj.removeAll(obj.conn.Collection(collection), bson.M(query))
// }

// func (obj *MongoDriver) removeAll(collection *mgo.Collection, query bson.M) (*mgo.ChangeInfo, *MDBError) {
// 	var change *mgo.ChangeInfo
// 	change, err := collection.RemoveAll(query)
// 	if err != nil {
// 		return change, getErrObj(ErrRemoveFailure, err.Error())
// 	}
// 	return change, nil
// }

// // Remove deletes the documents from the collection passed in the argument
// func (obj *MongoDriver) DropCollection(collection string) error {
// 	//	obj.session.Refresh()
// 	return obj.dropCollection(obj.conn.Collection(collection))
// }

// func (obj *MongoDriver) dropCollection(collection *mgo.Collection) error {
// 	err := collection.DropCollection()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // Count returns count of specified collection with where clause
// func (obj *MongoDriver) Count(collection string, query map[string]interface{}) (int, *MDBError) {
// 	//	obj.session.Refresh()
// 	return obj.count(obj.conn.Collection(collection), query)
// }

// func (obj *MongoDriver) count(collection *mgo.Collection, query bson.M) (int, *MDBError) {
// 	count := 0
// 	var err error
// 	if count, err = collection.Find(query).Count(); err != nil {
// 		return 0, getErrObj(ErrFindAllFailure, err.Error())
// 	}
// 	return count, nil
// }

// // FindWithField return given field parameters fields
// func (obj *MongoDriver) FindWithField(collection string, query map[string]interface{}, field map[string]interface{}) (ret []interface{}, aerr *MDBError) {
// 	//	obj.session.Refresh()
// 	return obj.findWithField(obj.conn.Collection(collection), bson.M(query), bson.M(field))
// }

// func (obj *MongoDriver) findWithField(collection *mgo.Collection, query bson.M, field bson.M) (ret []interface{}, aerr *MDBError) {
// 	if err := collection.Find(query).Select(field).All(&ret); err != nil {
// 		return nil, getErrObj(ErrFindAllFailure, err.Error())
// 	}
// 	return ret, nil
// }

// // FindAll queries the mongo DB and returns the sorted results with limit
// func (obj *MongoDriver) FindAndSort(collection string, query map[string]interface{}, selectField map[string]interface{}, sortField string, skip int, limit int) (ret []interface{}, cnt int, aerr *MDBError) {
// 	//	obj.session.Refresh()
// 	return obj.findAndSort(obj.conn.Collection(collection), bson.M(query), bson.M(selectField), sortField, skip, limit)
// }

// func (obj *MongoDriver) findAndSort(collection *mgo.Collection, query bson.M, selectField bson.M, sortField string, skip int, limit int) (ret []interface{}, cnt int, aerr *MDBError) {
// 	if err := collection.Find(query).Select(selectField).Sort(sortField).Skip(skip).Limit(limit).All(&ret); err != nil {
// 		return nil, 0, getErrObj(ErrFindAllFailure, err.Error())
// 	}
// 	count, errC := collection.Find(query).Count()
// 	if errC != nil {
// 		return nil, 0, getErrObj(ErrFindAllFailure, errC.Error())
// 	}
// 	return ret, count, nil
// }

// // FindAll queries the mongo DB and returns the results with limit
// func (obj *MongoDriver) Find(collection string, query map[string]interface{}, selectField map[string]interface{}, skip int, limit int) (ret []interface{}, cnt int, aerr *MDBError) {
// 	//	obj.session.Refresh()
// 	return obj.find(obj.conn.Collection(collection), bson.M(query), bson.M(selectField), skip, limit)
// }

// func (obj *MongoDriver) find(collection *mgo.Collection, query bson.M, selectField bson.M, skip int, limit int) (ret []interface{}, cnt int, aerr *MDBError) {
// 	if err := collection.Find(query).Select(selectField).Skip(skip).Limit(limit).All(&ret); err != nil {
// 		return nil, 0, getErrObj(ErrFindAllFailure, err.Error())
// 	}
// 	count, errC := collection.Find(query).Select(selectField).Count()
// 	if errC != nil {
// 		return nil, 0, getErrObj(ErrFindAllFailure, errC.Error())
// 	}
// 	return ret, count, nil
// }

// //Pipe used for aggregation accepts collection name and query in []bson.M
// func (obj *MongoDriver) Pipe(collection string, query interface{}) (pipe *mgo.Pipe, err error) {
// 	//	obj.session.Refresh()
// 	return obj.pipe(obj.conn.Collection(collection), query)
// }

// func (obj *MongoDriver) pipe(collection *mgo.Collection, query interface{}) (pipe *mgo.Pipe, err error) {
// 	pipe = collection.Pipe(query)
// 	return pipe, nil
// }

// // Update updates the mongo DB collection passed as an argument
// func (obj *MongoDriver) MoveElemToPosition(collection string, docMatcher map[string]interface{}, docField string, elemMatcher map[string]interface{}, position int) *MDBError {
// 	//	obj.session.Refresh()

// 	var filterCondition []interface{}
// 	filterKey := fmt.Sprintf("$$%s.%s",
// 		docField,
// 		elemMatcher["key"].(string),
// 	)
// 	inputField := fmt.Sprintf("$%s",
// 		docField,
// 	)

// 	filterCondition = append(filterCondition, filterKey)
// 	filterCondition = append(filterCondition, elemMatcher["value"])

// 	queryAggregation := []bson.M{}
// 	queryAggregation = append(queryAggregation, bson.M{"$match": docMatcher})
// 	queryAggregation = append(queryAggregation, bson.M{"$project": bson.M{"_id": 0, "lastPulledAt": 1, "data": bson.M{"$filter": bson.M{"input": inputField, "as": docField, "cond": bson.M{"$eq": filterCondition}}}}})
// 	queryAggregation = append(queryAggregation, bson.M{"$unwind": "$data"})

// 	pipe := obj.conn.Collection(collection).Pipe(queryAggregation)
// 	resp := bson.M{}
// 	errA := pipe.One(&resp)
// 	if errA != nil {
// 		return getErrObj(ErrFindAllFailure, errA.Error())
// 	}

// 	if resp != nil {
// 		element := resp["data"]
// 		docMatcher["lastPulledAt"] = resp["lastPulledAt"]

// 		var elementArray []interface{}
// 		elementArray = append(elementArray, element)

// 		var bulkValue []interface{}
// 		v1 := bson.M{"$pull": bson.M{docField: element}}
// 		v2 := bson.M{"$push": bson.M{docField: bson.M{"$each": elementArray, "$position": position}}}
// 		bulkValue = append(bulkValue, v1)
// 		bulkValue = append(bulkValue, v2)

// 		return obj.bulkUpdateOrder(obj.conn.Collection(collection), docMatcher, bulkValue)
// 	}
// 	return getErrObj(ErrResultNotFound, "No Mappings found to sort")
// }

// func (obj *MongoDriver) bulkUpdateOrder(collection *mgo.Collection, query map[string]interface{}, value []interface{}) *MDBError {
// 	bulk := collection.Bulk()
// 	for i := range value {
// 		bulk.Update(query, value[i])
// 	}
// 	result, err := bulk.Run()
// 	if err != nil {
// 		return getErrObj(ErrUpdateFailure, err.Error())
// 	}
// 	if result.Modified == 0 {
// 		return getErrObj(ErrResultNotFound, "Nothing found to update.")
// 	}
// 	return nil
// }

// func (obj *MongoDriver) FindAndModify(collection string, query map[string]interface{}, value interface{}, returnNew bool, upsert bool) (ret interface{}, aerr *MDBError) {
// 	//	obj.session.Refresh()
// 	return obj.findAndModify(obj.conn.Collection(collection), bson.M(query), value, returnNew, upsert)
// }

// func (obj *MongoDriver) findAndModify(collection *mgo.Collection, query bson.M, value interface{}, returnNew bool, upsert bool) (ret interface{}, aerr *MDBError) {
// 	change := mgo.Change{
// 		Update:    value,
// 		ReturnNew: returnNew,
// 		Upsert:    upsert,
// 	}
// 	var doc interface{}
// 	_, err := collection.Find(query).Apply(change, &doc)
// 	if err != nil {
// 		return nil, getErrObj(ErrFindAllFailure, err.Error())
// 	}
// 	return doc, nil

// }

// // BulkUpdate updates/inserts the bulk of data mongo DB collection passed as an argument
// func (obj *MongoDriver) BulkUpdate(collection string, query []map[string]interface{}, value []interface{}) *MDBError {
// 	if len(query) != len(value) {
// 		return getErrObj(ErrUpsertFailure, "Bulk.upsert requires an even number of parameters")
// 	}
// 	//	obj.session.Refresh()
// 	return obj.bulkUpdate(obj.conn.Collection(collection), query, value)
// }

// func (obj *MongoDriver) bulkUpdate(collection *mgo.Collection, query []map[string]interface{}, value []interface{}) *MDBError {
// 	bulk := collection.Bulk()
// 	for i := range query {
// 		bulk.Update(query[i], value[i])
// 	}
// 	if _, err := bulk.Run(); err != nil {
// 		return getErrObj(ErrUpsertFailure, err.Error())
// 	}
// 	return nil
// }

// // BulkUpdate updates/inserts the bulk of data mongo DB collection passed as an argument

// func (obj *MongoDriver) BulkUpdateAll(collection string, query []map[string]interface{}, value []interface{}) (BulkResult, *MDBError) {
// 	if len(query) != len(value) {
// 		return BulkResult{}, getErrObj(ErrUpsertFailure, "Bulk.upsert requires an even number of parameters")
// 	}
// 	//	obj.session.Refresh()
// 	return obj.bulkUpdateAll(obj.conn.Collection(collection), query, value)
// }

// func (obj *MongoDriver) bulkUpdateAll(collection *mgo.Collection, query []map[string]interface{}, value []interface{}) (BulkResult, *MDBError) {
// 	bulk := collection.Bulk()
// 	for i := range query {
// 		bulk.UpdateAll(query[i], value[i])
// 	}
// 	var err error
// 	bulkResultmgo, err := bulk.Run()
// 	if err != nil {
// 		return BulkResult{}, getErrObj(ErrUpdateFailure, err.Error())
// 	}
// 	bulkResult := BulkResult{
// 		Matched:  bulkResultmgo.Matched,
// 		Modified: bulkResultmgo.Modified,
// 	}
// 	if err != nil {
// 		errorResult, assrFailed := err.(*mgo.BulkError)
// 		if assrFailed == false {
// 			return BulkResult{}, getErrObj(ErrUpsertFailure, "Unable to assert Error array")
// 		}
// 		errorCases := errorResult.Cases()
// 		bulkResult.Ecases = make([]BulkErrorCase, 0)
// 		for _, eachError := range errorCases {
// 			bulkResult.Ecases = append(bulkResult.Ecases, BulkErrorCase{
// 				Index: eachError.Index,
// 				Err:   eachError.Err,
// 			})
// 		}
// 		return bulkResult, getErrObj(ErrUpsertFailure, err.Error())
// 	}
// 	return bulkResult, nil

// }

// func (obj *MongoDriver) FindSortnLoad(coll string, query map[string]interface{}, selectField map[string]interface{}, sortField string, skip int, limit int, payload interface{}) (cnt int, aerr *MDBError) {
// 	//	obj.session.Refresh()
// 	if err := obj.conn.Collection(coll).Find(query).Select(selectField).Sort(sortField).Skip(skip).Limit(limit).All(payload); err != nil {
// 		return 0, getErrObj(ErrFindAllFailure, err.Error())
// 	}
// 	count, errC := obj.conn.Collection(coll).Find(query).Count()
// 	if errC != nil {
// 		return 0, getErrObj(ErrFindAllFailure, errC.Error())
// 	}
// 	return count, nil
// }
