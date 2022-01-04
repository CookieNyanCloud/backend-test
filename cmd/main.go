package main

import (
	"flag"

	"github.com/cookienyancloud/avito-backend-test/internal/app"
	//_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
)

const configsDir = "configs"

func main() {
	var local bool
	flag.BoolVar(&local, "local", false, "host")
	flag.Parse()
	app.Run(configsDir, local)
}
