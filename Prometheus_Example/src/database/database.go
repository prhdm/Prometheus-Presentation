package database

type Database struct {
	tables map[string]map[string]interface{}
}

type IDatabase interface {
	CreateTable(tableName string)
	Get(tableName string, key string) interface{}
	Set(tableName string, key string, value interface{})
}

func NewDatabase() *Database {
	database := &Database{tables: make(map[string]map[string]interface{})}
	database.CreateTable("default")
	return database
}

func (d *Database) CreateTable(tableName string) {
	d.tables[tableName] = make(map[string]interface{})
}

func (d *Database) Get(tableName string, key string) interface{} {
	return d.tables[tableName][key]
}

func (d *Database) Set(tableName string, key string, value interface{}) {
	d.tables[tableName][key] = value
}

func (d *Database) Delete(tableName string, key string) interface{} {
	deleted := d.tables[tableName][key]
	delete(d.tables[tableName], key)
	return deleted
}
