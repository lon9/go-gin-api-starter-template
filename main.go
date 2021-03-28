package main

import (
	"flag"

	"{{ .ProjectPath }}/config"
	"{{ .ProjectPath }}/database"
	"{{ .ProjectPath }}/server"
)

func main() {

	env := flag.String("e", "development", "")
	flag.Parse()

	config.Init(*env)
	database.Init(false)
	if err := server.Init(); err != nil {
		panic(err)
	}
}
