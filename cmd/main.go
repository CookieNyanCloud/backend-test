package main

import(
	"github.com/cookienyancloud/avito-backend-test/internal/app"

)

const configsDir = "configs"

func main() {
	app.Run(configsDir)
}
