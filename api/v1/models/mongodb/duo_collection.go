package mongo

import (
	"os"
	"fmt"
	inp "lol-api/api/v1/input"
	"context"

	out "lol-api/api/v1/output"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"

	log "github.com/sirupsen/logrus"
)

func DBGetDuoList(_p *inp.URLParams) ([]*out.DuoObjComplete, error) {
	filter := bson.D{{}}
	res := []*out.DuoObjComplete{}

	collection := client.Database(fmt.Sprintf("%sleague-service", os.Getenv("DB_ENV"))).Collection(fmt.Sprintf("%sduo", os.Getenv("DB_ENV")))

	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return res, findError
	}

	for cur.Next(context.TODO()) {
		t := &out.DuoObj{}
		o := &out.DuoObjComplete{}
		err := cur.Decode(&t)
		if err != nil {
			return res, err
		}
		o.ID = t.ID
		o.PlayerData.PlayerName = t.PlayerData.PlayerName
		o.PlayerData.RolesPlayed.Top.RoleInfo = RoleCache[t.PlayerData.RolesPlayed.Top.RoleInfo]
		o.PlayerData.RolesPlayed.Top.Selected = t.PlayerData.RolesPlayed.Top.Selected
		o.PlayerData.RolesPlayed.Jungle.RoleInfo = RoleCache[t.PlayerData.RolesPlayed.Jungle.RoleInfo]
		o.PlayerData.RolesPlayed.Jungle.Selected = t.PlayerData.RolesPlayed.Jungle.Selected
		o.PlayerData.RolesPlayed.Mid.RoleInfo = RoleCache[t.PlayerData.RolesPlayed.Mid.RoleInfo]
		o.PlayerData.RolesPlayed.Mid.Selected = t.PlayerData.RolesPlayed.Mid.Selected
		o.PlayerData.RolesPlayed.Adc.RoleInfo = RoleCache[t.PlayerData.RolesPlayed.Adc.RoleInfo]
		o.PlayerData.RolesPlayed.Adc.Selected = t.PlayerData.RolesPlayed.Adc.Selected
		o.PlayerData.RolesPlayed.Support.RoleInfo = RoleCache[t.PlayerData.RolesPlayed.Support.RoleInfo]
		o.PlayerData.RolesPlayed.Support.Selected = t.PlayerData.RolesPlayed.Support.Selected
		o.PlayerData.RolesPlayed.Fill.RoleInfo = RoleCache[t.PlayerData.RolesPlayed.Fill.RoleInfo]
		o.PlayerData.RolesPlayed.Fill.Selected = t.PlayerData.RolesPlayed.Fill.Selected
		o.PlayerData.SoloQRank = RankCache[t.PlayerData.SoloQRank]

		o.PlayerData.Server = ServerCache[t.PlayerData.Server]
		o.PlayerData.OpggData = t.PlayerData.OpggData

		o.LookingFor.LadderMode =  LadderCache[t.LookingFor.LadderMode]
		o.LookingFor.BuddyRank.MinRank =  RankCache[t.LookingFor.BuddyRank.MinRank]
		o.LookingFor.BuddyRank.MaxRank =  RankCache[t.LookingFor.BuddyRank.MaxRank]

		o.LookingFor.BuddyRoles.Top.RoleInfo = RoleCache[t.LookingFor.BuddyRoles.Top.RoleInfo]
		o.LookingFor.BuddyRoles.Top.Selected = t.LookingFor.BuddyRoles.Top.Selected
		o.LookingFor.BuddyRoles.Jungle.RoleInfo = RoleCache[t.LookingFor.BuddyRoles.Jungle.RoleInfo]
		o.LookingFor.BuddyRoles.Jungle.Selected = t.LookingFor.BuddyRoles.Jungle.Selected
		o.LookingFor.BuddyRoles.Mid.RoleInfo = RoleCache[t.LookingFor.BuddyRoles.Mid.RoleInfo]
		o.LookingFor.BuddyRoles.Mid.Selected = t.LookingFor.BuddyRoles.Mid.Selected
		o.LookingFor.BuddyRoles.Adc.RoleInfo = RoleCache[t.LookingFor.BuddyRoles.Adc.RoleInfo]
		o.LookingFor.BuddyRoles.Adc.Selected = t.LookingFor.BuddyRoles.Adc.Selected
		o.LookingFor.BuddyRoles.Support.RoleInfo = RoleCache[t.LookingFor.BuddyRoles.Support.RoleInfo]
		o.LookingFor.BuddyRoles.Support.Selected = t.LookingFor.BuddyRoles.Support.Selected
		o.LookingFor.BuddyRoles.Fill.RoleInfo = RoleCache[t.LookingFor.BuddyRoles.Fill.RoleInfo]
		o.LookingFor.BuddyRoles.Fill.Selected = t.LookingFor.BuddyRoles.Fill.Selected

		o.LookingFor.ChatUsage = t.LookingFor.ChatUsage
		res = append(res, o)
	}

	cur.Close(context.TODO())
	if len(res) == 0 {
		return res, mongo.ErrNoDocuments
	}

	return res, nil
}

func DBDeleteDuoByID(_p *inp.URLParams) (string, error) {
	objectId, err := GetSingleID((*_p)[inp.QColId])
	if err != nil{
		log.Println("Invalid id")
	}
	filter := bson.M{"_id": objectId}

	collection := client.Database(fmt.Sprintf("%sleague-service", os.Getenv("DB_ENV"))).Collection(fmt.Sprintf("%sduo", os.Getenv("DB_ENV")))

	_, delError := collection.DeleteMany(context.TODO(), filter)
	if delError != nil {
		return "", delError
	}

	return "collection deleted", nil
}

func DBCreateDuo(_o inp.DuoObj) (string, error) {
	collection := client.Database(fmt.Sprintf("%sleague-service", os.Getenv("DB_ENV"))).Collection(fmt.Sprintf("%sduo", os.Getenv("DB_ENV")))

	_, creError := collection.InsertOne(context.TODO(), _o)
	if creError != nil {
		return "", creError
	}

	return "collection created", nil
}
