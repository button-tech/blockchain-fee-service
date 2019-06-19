package server

import (
	"dev.azure.com/fee-service/services/server/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine {
	return gin.New()
}

func RunServer(R *gin.Engine) error {
	R.Use(gin.Recovery())
	R.Use(gin.Logger())
	// TODO: настроить cors
	R.Use(cors.Default())
	{

		api := R.Group("/api/services")
		{
			fee := api.Group("/fee")
			{
				fee.POST("/bitcoin", handlers.GetBitcoinFee)
				fee.POST("/litecoin", handlers.GetLitecoinFee)
				fee.POST("/bitcoinCash", handlers.GetBitcoinCashFee)
				fee.POST("/ethereum", handlers.GetEthereumFee)
				fee.POST("/ethereumClassic", handlers.GetEthereumClassicFee)
			}
		}
	}

	if err := R.Run(":8080"); err != nil {
		return err
	}

	return nil
}
