package app

import (
	"lambda-func/api"
	"lambda-func/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {
	// We init db store gets passed DOWN into the api handler
	db := database.NewDynamoDBClient()
	apihandler := api.NewApiHandler(db)

	return App{
		ApiHandler: apihandler,
	}
}
