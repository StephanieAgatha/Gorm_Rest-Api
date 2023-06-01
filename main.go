// package untuk running gin dan database (kalo bisa email juga hehe)
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api-h8/config"
	"rest-api-h8/controllers"
)

func main() {

	//starting gin
	g := gin.Default()
	//connect to db
	config.DBConnect()
	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Messages": "Hehe",
		})
	})

	//create another routine
	g.GET("/example", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Message": "This is example routing",
		})
	})

	//create products routing
	g.GET("/products", controllers.HandlerGetProduct)

	//create orders routing
	g.POST("/orders", controllers.HandlerPostProduct)

	//create get orders routing
	g.GET("/orders", controllers.HandlerGetOrders)

	//run gin with default port :8080
	g.Run()
}
