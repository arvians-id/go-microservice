package main

import (
	"github.com/arvians-id/go-microservice/adapter/pkg/auth"
	"github.com/arvians-id/go-microservice/adapter/pkg/product"
	"github.com/arvians-id/go-microservice/adapter/pkg/user"
	"github.com/arvians-id/go-microservice/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	configuration, err := config.LoadConfig("../config/envs")
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := auth.RegisterRoutes(r, configuration)
	user.RegisterRoutes(r, configuration, authSvc)
	product.RegisterRoutes(r, configuration, authSvc)

	err = r.Run(configuration.Port)
	if err != nil {
		log.Fatalln("Failed at running", err)
	}
}
