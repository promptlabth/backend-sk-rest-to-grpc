package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/promptlabth/orch_sk/app/calculator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	g := gin.Default()

	cred := insecure.NewCredentials()
	cc, err := grpc.Dial("core-grpc-deployment-deploy.petch:8080", grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()
	calculatorClient := calculator.NewCalculatorClient(cc)
	calculatorService := calculator.NewCalculatorService(calculatorClient)

	g.POST("/get", calculatorService.Hello)

	if err := g.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
