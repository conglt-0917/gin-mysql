package main

import (
	"fmt"
	"log"
	"os"

	. "github.com/conglt-0917/gin-mysql/middlewares"
	. "github.com/conglt-0917/gin-mysql/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	//Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, please create one in the root directory")
	}

	if err != nil {
		fmt.Println(err)
	}

	r.Use(CORS())
	SetUpRouters(r)

	port := os.Getenv("PORT")

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	if os.Getenv("TSL") == "TRUE" {
		TSLKeys := &struct {
			CERT string
			KEY  string
		}{}

		TSLKeys.CERT = "./cert/myCA.cer"
		TSLKeys.KEY = "./cert/myCA.key"

		r.RunTLS(":"+port, TSLKeys.CERT, TSLKeys.KEY)

	} else {
		r.Run(":" + port)
	}

}
