package mongodb

// BulkResult holds the results for a bulk operation.
type BulkResult struct {
	Matched  int
	Modified int // Available only for MongoDB 2.6+
	Ecases   []BulkErrorCase
	// Be conservative while we understand exactly how to report these
	// results in a useful and convenient way, and also how to emulate
	// them with prior servers.

}

type BulkErrorCase struct {
	Index int // Position of operation that failed, or -1 if unknown.
	Err   error
}
