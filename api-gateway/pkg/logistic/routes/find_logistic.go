package routes

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yusrilsabir22/orderfaz/api-gateway/pkg/logistic/pb"
)

func FindLogistic(ctx *gin.Context, c pb.LogisticServiceClient) {
	dn := ctx.Request.URL.Query().Get("destination_name")
	on := ctx.Request.URL.Query().Get("origin_name")

	if dn == "" || on == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	log.Println("query: ", dn)

	res, err := c.FindOne(context.Background(), &pb.FindOneLogisticRequest{
		OriginName:      on,
		DestinationName: dn,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
