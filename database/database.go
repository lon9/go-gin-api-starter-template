package database

import (
	"github.com/jinzhu/gorm"
	"{{ .ProjectPath }}/config"
)

var d *gorm.DB

// Init initializes database
func Init(isReset bool, models ...interface{}) {
	c := config.GetConfig()
	var err error
	d, err = gorm.Open(c.GetString("db.provider"), c.GetString("db.url"))
	if err != nil {
		panic(err)
	}
	if isReset {
		d.DropTableIfExists(models)
	}
	d.AutoMigrate(models...)
}

// GetDB returns database connection
func GetDB() *gorm.DB {
	return d
}

// Close closes database
func Close() {
	d.Close()
}
