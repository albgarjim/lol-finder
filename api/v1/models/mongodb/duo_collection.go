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

	// log "github.com/sirupsen/logrus"
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
		o.PlayerData.RolesPlayed.Top.LaneInfo = RoleCache[t.PlayerData.RolesPlayed.Top.LaneInfo]
		o.PlayerData.RolesPlayed.Jungle.LaneInfo = RoleCache[t.PlayerData.RolesPlayed.Jungle.LaneInfo]
		o.PlayerData.RolesPlayed.Mid.LaneInfo = RoleCache[t.PlayerData.RolesPlayed.Mid.LaneInfo]
		o.PlayerData.RolesPlayed.Adc.LaneInfo = RoleCache[t.PlayerData.RolesPlayed.Adc.LaneInfo]
		o.PlayerData.RolesPlayed.Support.LaneInfo = RoleCache[t.PlayerData.RolesPlayed.Support.LaneInfo]
		o.PlayerData.SoloQRank = RankCache[t.PlayerData.SoloQRank]
		res = append(res, o)

		o.PlayerData.Server = ServerCache[t.PlayerData.Server]
		o.PlayerData.OpggData = t.PlayerData.OpggData

		o.LookingFor.GameMode =  LadderCache[t.LookingFor.GameMode]
		o.LookingFor.BuddyRank.MinRank =  RankCache[t.LookingFor.BuddyRank.MinRank]
		o.LookingFor.BuddyRank.MaxRank =  RankCache[t.LookingFor.BuddyRank.MaxRank]

		o.LookingFor.BuddyRoles.Top.LaneInfo = RoleCache[t.LookingFor.BuddyRoles.Top.LaneInfo]
		o.LookingFor.BuddyRoles.Jungle.LaneInfo = RoleCache[t.LookingFor.BuddyRoles.Jungle.LaneInfo]
		o.LookingFor.BuddyRoles.Mid.LaneInfo = RoleCache[t.LookingFor.BuddyRoles.Mid.LaneInfo]
		o.LookingFor.BuddyRoles.Adc.LaneInfo = RoleCache[t.LookingFor.BuddyRoles.Adc.LaneInfo]
		o.LookingFor.BuddyRoles.Support.LaneInfo = RoleCache[t.LookingFor.BuddyRoles.Support.LaneInfo]

		o.LookingFor.ChatUsage = t.LookingFor.ChatUsage
	}

	cur.Close(context.TODO())
	if len(res) == 0 {
		return res, mongo.ErrNoDocuments
	}

	return res, nil
}
