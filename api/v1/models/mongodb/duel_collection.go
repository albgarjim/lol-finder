package mongo

import (
	"os"
	"fmt"
	inp "lol-api/api/v1/input"
	"context"

	out "lol-api/api/v1/output"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"

	log "github.com/sirupsen/logrus"
)

func DBGetDuelList(_p *inp.URLParams) ([]*out.TestObj, error) {
	filter := bson.D{{}}
	res := []*out.TestObj{}

	collection := client.Database(fmt.Sprintf("%stest-db", os.Getenv("DB_ENV"))).Collection(fmt.Sprintf("%stest-coll", os.Getenv("DB_ENV")))

	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return res, findError
	}

	for cur.Next(context.TODO()) {
		t := &out.TestObj{}
		err := cur.Decode(&t)
		if err != nil {
			return res, err
		}
		res = append(res, t)
	}

	cur.Close(context.TODO())
	if len(res) == 0 {
		return res, mongo.ErrNoDocuments
	}

	return res, nil
}

func DBGetDuelByID(_p *inp.URLParams) ([]*out.TestObj, error) {
	// idList := GetIDs((*_p)[inp.QColList])
	// log.Info("====", idList);
	// filter := bson.M{"_id": bson.M{"$in": idList}}

	objectId, err := GetSingleID((*_p)[inp.QColId])
	if err != nil{
		log.Println("Invalid id")
	}
	filter := bson.M{"_id": objectId}

	res := []*out.TestObj{}

	collection := client.Database(fmt.Sprintf("%stest-db", os.Getenv("DB_ENV"))).Collection(fmt.Sprintf("%stest-coll", os.Getenv("DB_ENV")))

	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return res, findError
	}

	for cur.Next(context.TODO()) {
		t := &out.TestObj{}
		err := cur.Decode(&t)
		if err != nil {
			return res, err
		}
		res = append(res, t)
	}

	cur.Close(context.TODO())
	if len(res) == 0 {
		return res, mongo.ErrNoDocuments
	}

	return res, nil
}

func DBDeleteDuelByID(_p *inp.URLParams) (string, error) {
	objectId, err := GetSingleID((*_p)[inp.QColId])
	if err != nil{
		log.Println("Invalid id")
	}
	filter := bson.M{"_id": objectId}

	collection := client.Database(fmt.Sprintf("%stest-db", os.Getenv("DB_ENV"))).Collection(fmt.Sprintf("%stest-coll", os.Getenv("DB_ENV")))

	_, delError := collection.DeleteMany(context.TODO(), filter)
	if delError != nil {
		return "", delError
	}

	return "collection deleted", nil
}

func DBCreateDuel(_p *inp.URLParams) (string, error) {
	var res inp.TestObj
	res.Title="create-object"

	collection := client.Database(fmt.Sprintf("%stest-db", os.Getenv("DB_ENV"))).Collection(fmt.Sprintf("%stest-coll", os.Getenv("DB_ENV")))

	_, creError := collection.InsertOne(context.TODO(), res)
	if creError != nil {
		return "", creError
	}

	return "collection created", nil
}

func DBUpdateDuelByID(_p *inp.URLParams, _u inp.TestObj) (string, error) {
	objectId, err := GetSingleID((*_p)[inp.QColId])
	if err != nil{
		log.Println("Invalid id")
	}
	filter := bson.M{"_id": objectId}

	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "title", Value: _u.Title},
	}}}

	collection := client.Database(fmt.Sprintf("%stest-db", os.Getenv("DB_ENV"))).Collection(fmt.Sprintf("%stest-coll", os.Getenv("DB_ENV")))

	_, delError := collection.UpdateOne(context.TODO(), filter, updater)
	if delError != nil {
		return "", delError
	}

	return "collection udpated", nil
}
