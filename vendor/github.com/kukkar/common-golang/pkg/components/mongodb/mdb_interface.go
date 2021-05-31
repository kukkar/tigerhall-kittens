package mongodb

import "context"

var _ MDBInterface = (*MongoDriver)(nil)

// mongodbb interface
type MDBInterface interface {
	// init initializes the mongodb instance
	Init(*MDBConfig) *MDBError

	// FindOne returns first matching item
	FindOne(context.Context, string, map[string]interface{}) (interface{}, *MDBError)

	// // FindOnenLoad: find a document and load in supplied payload
	// FindOnenLoad(coll string, q map[string]interface{}, payload interface{}) *MDBError

	// //FindOneUsingSession(*MSession, string, map[string]interface{}) (interface{}, *MDBError)

	// // FindAll returns all matching items
	// FindAll(string, map[string]interface{}) ([]interface{}, *MDBError)

	// //FindAllUsingSession(*MSession, string, map[string]interface{}) ([]interface{}, *MDBError)

	// // FindWithLimit returns all matching items with limit
	// FindWithLimit(string, map[string]interface{}, int, int) ([]interface{}, *MDBError)

	// // FindWithField used to find data and returns only given fields
	// FindWithField(string, map[string]interface{}, map[string]interface{}) ([]interface{}, *MDBError)

	// GetNewID(collection, idName, counterName string) (interface{}, *MDBError)

	// //GetNewCustomID is used to increment ID by given incValue and returns it
	// GetNewCustomID(collection, idName, counterName string, incValue int) ([]int, *MDBError)

	// // Insert add one item
	Insert(context.Context, string, interface{}) *MDBError

	// // Insert add one or more items
	BulkInsert(context.Context, string, []interface{}) *MDBError

	// //	InsertUsingSession(*MSession, string, interface{}) *MDBError

	// // Update modify existing item
	// Update(string, map[string]interface{}, interface{}) *MDBError

	// UpdateAll(string, map[string]interface{}, interface{}) (*mgo.ChangeInfo, *MDBError)

	// // Update one or more items where each item will be query and value
	// // Only update one document of matched criteria.
	// BulkUpdate(collection string, query []map[string]interface{}, value []interface{}) *MDBError

	// // Update one or more items where each item will be query and value
	// // Update all documents of matched criteria.
	// BulkUpdateAll(collection string, query []map[string]interface{}, value []interface{}) (BulkResult, *MDBError)

	// //UpdateUsingSession(*MSession, string, map[string]interface{}, interface{}) *MDBError

	// // Update modify existing item or insert new item if does not exist
	//	Upsert(context.Context, string, map[string]interface{}, interface{}) *MDBError

	// // Upsert one or more items where each item will be query and value
	// BulkUpsert(string, []map[string]interface{}, []interface{}) (BulkResult, *MDBError)

	// //UpsertUsingSession(*MSession, string, map[string]interface{}, interface{}) *MDBError

	// // Remove delete existing item
	// Remove(string, map[string]interface{}) *MDBError

	// RemoveAll(collection string, query map[string]interface{}) (*mgo.ChangeInfo, *MDBError)

	// DropCollection(collection string) error

	// //RemoveUsingSession(*MSession, string, map[string]interface{}) *MDBError

	// // // Close shuts down the current session.
	// // Close(*MSession)

	// // // Copy creates a copy of master session
	// // Copy(safe *Safe) *MSession

	// // // CloseMasterSession closes the master session
	// // CloseMasterSession()

	// // Count
	// Count(string, map[string]interface{}) (int, *MDBError)

	// // Find -- collection,selectedField,limit,skip
	// Find(string, map[string]interface{}, map[string]interface{}, int, int) ([]interface{}, int, *MDBError)

	// // FindAll queries the mongo DB and returns the sorted results with limit
	// FindAndSort(string, map[string]interface{}, map[string]interface{}, string, int, int) ([]interface{}, int, *MDBError)

	// //Pipe is used for aggregation
	// Pipe(string, interface{}) (*mgo.Pipe, error)

	// // Move existing array element to new position
	// MoveElemToPosition(string, map[string]interface{}, string, map[string]interface{}, int) *MDBError

	// FindAndModify(string, map[string]interface{}, interface{}, bool, bool) (interface{}, *MDBError)

	// FindSortnLoad(coll string, q map[string]interface{}, selField map[string]interface{}, order string, page int, limit int, payload interface{}) (int, *MDBError)
}
