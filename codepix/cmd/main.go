package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/lucaswilliameufrasio/imersao/codepix-go/app/grpc"
	"github.com/lucaswilliameufrasio/imersao/codepix-go/infra/db"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
