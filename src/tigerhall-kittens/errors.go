package tigerhall

import "errors"

var ErrTigerIn5KilometerRadios = errors.New("tiger is in 5 kilometer radios from last location")

// File not found
var ErrFileNotFound error = errors.New("File not found")

//Permission Denied
var ErrPermissionDenied error = errors.New("Permission Denied")

//Mongo session not supplied.
var ErrMongoNotSupplied error = errors.New("Mongo Not Supplied")

//Storage Adapter not supplied.
var ErrStorageAdapterNotSupplied error = errors.New("Storage Adapter Not Supplied")

var ErrToBeImplemented error = errors.New("To be Implemented")

// File not found
var ErrNotFound error = errors.New("not found")

// DB execution error
var ErrDBExecution error = errors.New("unable to proferm db task")
