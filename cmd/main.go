package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rdsdataservice"
	"github.com/gin-gonic/gin"
	"github.com/paujim/pocAurora/controllers"
	"github.com/paujim/pocAurora/repositories"
)

var app *controllers.App

func init() {
	auroraArn := os.Getenv("AURORA_ARN")
	secretArn := os.Getenv("SECRET_ARN")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	app = controllers.NewApp(
		repositories.NewSQLClient(rdsdataservice.New(sess), aws.String(auroraArn), aws.String(secretArn)),
		gin.Default(),
	)
}

func main() {
	router := app.SetupServer()
	router.Run()
}
