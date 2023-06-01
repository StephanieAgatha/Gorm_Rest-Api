// i'm struggling :(

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MailerRequest struct {
	BuyerEmail   string `json:"buyer_email"`
	BuyerAddress string `json:"buyer_address"`
	ProductName  string `json:"product_name"`
}

func main() {

	//inisiasi gin
	s := gin.Default()

	//create routing mailer using post
	s.POST("/mailer", func(ctx *gin.Context) {
		var mailerBody MailerRequest

		//handling error
		if err := ctx.ShouldBindJSON(&mailerBody); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, "Mailer Service Running ")
	})

	//run with custom port
	s.Run(":8002")

}
