package mongo

import (
	"fmt"
	"os"
	"time"
	"context"
	"sync"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	// r "github.com/dancannon/gorethink"
	log "github.com/sirupsen/logrus"
)

var mongoOnce sync.Once
var ctx = context.TODO()
var client *mongo.Client

func InitMongoDB() {
	var err error
	mongoOnce.Do(func() {
		var dbTimeout time.Duration = 100

		log.Info("Connecting to mongodb...")

		mongoAddress := fmt.Sprintf("mongodb://%s:%s/", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
		log.Info(mongoAddress)

		clientOptions := options.Client().ApplyURI(mongoAddress)

		for {
			client, err = mongo.Connect(ctx, clientOptions)
			if err != nil {
				log.Error("Database connection denied, attempting to reconnect after ", dbTimeout, " miliseconds")
				time.Sleep(dbTimeout * time.Millisecond)
				dbTimeout *= 2

			} else {
				log.Info("Connection successful")

				break
			}

		}
	})
	log.Info("Connection successful to mongodb")
}
