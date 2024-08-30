package main

import (
	"Cart/controllers"
	"Cart/database"

	"github.com/gin-gonic/gin"
)

func main() {
	Router := gin.Default()
	database.Connectdatabase()
	cnt := &controllers.Basecontrollers{}
	itemcnt := &controllers.Itemcontrollers{}
	cart := &controllers.Cartcontrollers{}
	cnt.Dostuff(Router)
	itemcnt.Itemmain(Router)
	cart.Cartmain(Router)
	Router.Run("localhost:3000")
}
