package main

import (
	"github.com/mohammad-hamzei/echo-social/db"
	"github.com/mohammad-hamzei/echo-social/route"
	//"echo-gorm/db"
	//"echo-gorm/route"
)

func main() {
	db.Init()
	e := route.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
