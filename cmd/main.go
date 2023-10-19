package main

import (
	"os"

	"github.com/bertoldoklinger/codepix/application/grpc"
	"github.com/bertoldoklinger/codepix/infra/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}