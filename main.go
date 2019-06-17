package main

import (
	"log"
	"os"

	"dev.azure.com/fee-service/services/server"
	"github.com/gin-gonic/gin"
)

var (
	ServerInstance *gin.Engine
)

func init() {
	ServerInstance = server.InitServer()
}

func main() {
	if err := server.RunServer(ServerInstance); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
