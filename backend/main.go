package main

import (
	"ReserveMate/backend/pkg/config"
	"ReserveMate/backend/pkg/infrastructure/datastore"
	"ReserveMate/backend/pkg/infrastructure/router"
	"ReserveMate/backend/pkg/registry"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()

	datastore.AutoMigrate(db)

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
