package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"{{ .ProjectPath }}/config"
)

var d *gorm.DB

// Init initializes database
func Init(isReset bool, models ...interface{}) {
	c := config.GetConfig()
	var err error
	provider := c.GetString("db.provider")
	var prov gorm.Dialector
	switch provider {
	case "postgres":
		prov = postgres.Open(c.GetString("db.url"))
	case "mysql":
		prov = mysql.Open(c.GetString("db.url"))
	case "sqlite":
		prov = sqlite.Open(c.GetString("db.url"))
	default:
		panic("unsupported db provider")
	}
	d, err = gorm.Open(prov, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if isReset {
		if err := d.Migrator().DropTable(models...); err != nil {
			panic(err)
		}
	}
	if err := d.AutoMigrate(models...); err != nil {
		panic(err)
	}
}

// GetDB returns database connection
func GetDB() *gorm.DB {
	return d
}
