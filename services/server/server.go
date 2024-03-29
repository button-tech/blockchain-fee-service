package server

import (
	"github.com/button-tech/blockchain-fee-service/services/server/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine {
	g := gin.New()
	gin.SetMode(gin.ReleaseMode)
	return g
}

func RunServer(R *gin.Engine) error {
	R.Use(gin.Recovery())
	R.Use(gin.Logger())
	// TODO: настроить cors
	R.Use(cors.Default())
	{

		api := R.Group("/fee")
		{
			api.POST("/bitcoin", handlers.GetBitcoinFee)
			api.POST("/litecoin", handlers.GetLitecoinFee)
			api.POST("/bitcoinCash", handlers.GetBitcoinCashFee)
			api.POST("/ethereum", handlers.GetEthereumFee)
			api.POST("/ethereumClassic", handlers.GetEthereumClassicFee)
			api.POST("/token", handlers.GetTokenFee)
			api.POST("/waves", handlers.GetWavesFee)
			api.POST("/stellar", handlers.GetStellarFee)
			api.POST("/zilliqa", handlers.GetZilliqaFee)
		}
	}

	if err := R.Run(":8080"); err != nil {
		return err
	}

	return nil
}
