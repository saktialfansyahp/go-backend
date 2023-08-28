package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saktialfansyahp/go-rest-api/controllers/productcontrollers"
	"github.com/saktialfansyahp/go-rest-api/models"
)

func DefineRoutes(){
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("api/products", productcontrollers.Index)
	// r.GET("api/products", productcontrollers.Index)
	
	r.Run()
}