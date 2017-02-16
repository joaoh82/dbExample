# dbExample
Golang example of how to set up a Database interface to be used in an application

## db.go
Database interface with common method to be implement by N different databases

```
Insert(collection string, item interface{}) error
BulkInsert(collection string, items []interface{}) error
```

## mongodb.go
Example of implementation of the Database interface when it comes to MongoDB database

## Created by
João Henrique Machado Silva, joaoh82@gmail.com, @joaoh82
