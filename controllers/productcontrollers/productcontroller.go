package productcontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saktialfansyahp/go-rest-api/models"
	"github.com/saktialfansyahp/go-rest-api/pkg/utils"
)

func Index(c *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)
	token, err := utils.GenerateToken(2)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Product": products, "Token": token})

}