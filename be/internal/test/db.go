package test

import (
	"flag"
	"github.com/ioartigiano/ioartigiano-be/internal/config"
	"github.com/jinzhu/gorm"
	"testing"
)

var db *gorm.DB
var file = flag.String("config", "../../config/local.json", "path to the config file")

// DB returns the database connection for testing purpose.
func DB(t *testing.T) *gorm.DB {
	if db != nil {
		return db
	}

	cfg := config.Load()
	db, err := gorm.Open(cfg.Database.Dialect(), cfg.Database.ConnectionInfo())
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	return db
}

// ResetTables truncates all data in the specified tables.
func ResetTables(t *testing.T, db *gorm.DB, tables ...string) {
	for _, table := range tables {
		err := db.Exec("TRUNCATE TABLE " + table).Error
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		t.Log("table ", table, " truncated! ")
	}
}

func PopulateTable(t *testing.T, db *gorm.DB, queries ...string) {
	for _, query := range queries {
		err := db.Exec(query).Error
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		t.Log("Table populated by query: ", query)
	}
}
