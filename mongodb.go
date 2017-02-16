package db

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
)

// MongoDB type
type MongoDB struct {
	sess *mgo.Session
	db   *mgo.Database
}

// DbName is exported here to be able to be setted in the implementation
var DbName string
var dbHost string

const (
	// Constants with name of DB collections
	CollectionClients   = "clients"
	CollectionSuppliers = "suppliers"
	CollectionServices  = "services"
	CollectionProjects  = "projects"
	CollectionEmployees = "employees"
)

// NewMongoDB creates a new mongo database based on the host and db name
func NewMongoDB(dbHost string, dbName string) Database {
	sess, err := mgo.DialWithTimeout(dbHost, time.Second)
	if err != nil {
		log.Println("MongoDB not available")
		return nil
	}

	db := sess.DB(dbName)

	log.Println("Using MongoDB:", dbHost, "/", DbName)

	return &MongoDB{sess: sess, db: db}
}

// Insert method implementing the Database interface.
// Used to insert one document at a time
func (m *MongoDB) Insert(collection string, item interface{}) error {
	c := m.db.C(collection)
	_, err := c.Upsert(item, item)
	return err
}

// Bulkinsert method implementing the Databse interface.
// Used to insert more then one document at a time.
func (m *MongoDB) BulkInsert(collection string, items []interface{}) error {
	c := m.db.C(collection)
	b := c.Bulk()
	b.Insert(items...)
	_, err := b.Run()
	return err
}
