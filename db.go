package db

// Database interface that could be used to implement N different databases.
type Database interface {
	Insert(collection string, item interface{}) error
	BulkInsert(collection string, items []interface{}) error
}
