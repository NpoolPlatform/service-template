package db

import (
	"github.com/NpoolPlatform/go-service-app-template/pkg/db/ent"

	"github.com/NpoolPlatform/go-service-framework/pkg/app"
)

var myClient *ent.Client

func Init() error {
	myClient = ent.NewClient(ent.Driver(app.Mysql().Driver))
	return nil
}

func Client() *ent.Client {
	return myClient
}
