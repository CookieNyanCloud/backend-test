package main

import (
	"github.com/cookienyancloud/avito-backend-test/internal/app"
	_ "github.com/lib/pq"
)

const configsDir = "configs"

func main() {
	app.Run(configsDir)
}
