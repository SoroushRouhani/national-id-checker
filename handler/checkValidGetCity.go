package handler

import (
	"github.com/gin-gonic/gin"
	"project/model"
)

func CheckValidGetCity(c *gin.Context) {

	number := c.Query("id")
	nationalID := model.NewNationalID(number)
	city := nationalID.GetCity()
	if nationalID.IsValid{ 

		c.JSON(200, gin.H{

			"City": city,
			"National-ID" : number,
			"Result": "National-ID is Valid",
		})

	}else {

		c.JSON(404, gin.H{

			"National-ID" : number,
			"Error": "National-ID is Invalid",
		})
	}
}