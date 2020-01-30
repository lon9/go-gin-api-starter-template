package main

import (
	"flag"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"{{ .ProjectPath }}/config"
	"{{ .ProjectPath }}/database"
	"{{ .ProjectPath }}/server"
)

func main() {

	env := flag.String("e", "development", "")
	flag.Parse()

	config.Init(*env)
	database.Init(false)
	defer database.Close()
	if err := server.Init(); err != nil {
		panic(err)
	}
}
