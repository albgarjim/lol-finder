package mongo

import (
	"os"
	"fmt"
	// inp "lol-api/api/v1/input"
	"context"

	out "lol-api/api/v1/output"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"

	log "github.com/sirupsen/logrus"
)

func DBRanksCache() (map[string]out.RankObj, error) {
	filter := bson.D{{}}
	res := map[string]out.RankObj{}

	collection := client.Database(fmt.Sprintf("%sleague-service", os.Getenv("DB_ENV"))).Collection(fmt.Sprintf("%sranks", os.Getenv("DB_ENV")))

	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return res, findError
	}

	for cur.Next(context.TODO()) {
		t := out.RankObj{}
		err := cur.Decode(&t)
		if err != nil {
			return res, err
		}
		res[t.ID] = t
	}

	cur.Close(context.TODO())
	if len(res) == 0 {
		return res, mongo.ErrNoDocuments
	}
	log.Info("Loaded ", len(res), " ranks in RankObj");

	return res, nil
}

func DBServerCache() (map[string]out.ServerObj, error) {
	filter := bson.D{{}}
	res := map[string]out.ServerObj{}

	collection := client.Database(fmt.Sprintf("%sleague-service", os.Getenv("DB_ENV"))).Collection(fmt.Sprintf("%sservers", os.Getenv("DB_ENV")))

	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return res, findError
	}

	for cur.Next(context.TODO()) {
		t := out.ServerObj{}
		err := cur.Decode(&t)
		if err != nil {
			return res, err
		}
		res[t.ID] = t
	}

	cur.Close(context.TODO())
	if len(res) == 0 {
		return res, mongo.ErrNoDocuments
	}
	log.Info("Loaded ", len(res), " servers in ServerObj");

	return res, nil
}


func DBLadderCache() (map[string]out.LadderObj, error) {
	filter := bson.D{{}}
	res := map[string]out.LadderObj{}

	collection := client.Database(fmt.Sprintf("%sleague-service", os.Getenv("DB_ENV"))).Collection(fmt.Sprintf("%sladders", os.Getenv("DB_ENV")))

	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return res, findError
	}

	for cur.Next(context.TODO()) {
		t := out.LadderObj{}
		err := cur.Decode(&t)
		if err != nil {
			return res, err
		}
		res[t.ID] = t
	}

	cur.Close(context.TODO())
	if len(res) == 0 {
		return res, mongo.ErrNoDocuments
	}
	log.Info("Loaded ", len(res), " ladders in LadderObj");

	return res, nil
}


func DBRoleCache() (map[string]out.RoleObj, error) {
	filter := bson.D{{}}
	res := map[string]out.RoleObj{}

	collection := client.Database(fmt.Sprintf("%sleague-service", os.Getenv("DB_ENV"))).Collection(fmt.Sprintf("%sroles", os.Getenv("DB_ENV")))

	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return res, findError
	}

	for cur.Next(context.TODO()) {
		t := out.RoleObj{}
		err := cur.Decode(&t)
		if err != nil {
			return res, err
		}
		res[t.ID] = t
	}

	cur.Close(context.TODO())
	if len(res) == 0 {
		return res, mongo.ErrNoDocuments
	}
	log.Info("Loaded ", len(res), " roles in RoleObj");

	return res, nil
}
