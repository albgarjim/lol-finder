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
		o.PlayerData.RolesPlayed.Top.RoleInfo = RoleCache[t.PlayerData.RolesPlayed.Top.RoleInfo]
		o.PlayerData.RolesPlayed.Jungle.RoleInfo = RoleCache[t.PlayerData.RolesPlayed.Jungle.RoleInfo]
		o.PlayerData.RolesPlayed.Mid.RoleInfo = RoleCache[t.PlayerData.RolesPlayed.Mid.RoleInfo]
		o.PlayerData.RolesPlayed.Adc.RoleInfo = RoleCache[t.PlayerData.RolesPlayed.Adc.RoleInfo]
		o.PlayerData.RolesPlayed.Support.RoleInfo = RoleCache[t.PlayerData.RolesPlayed.Support.RoleInfo]
		o.PlayerData.SoloQRank = RankCache[t.PlayerData.SoloQRank]
		res = append(res, o)

		o.PlayerData.Server = ServerCache[t.PlayerData.Server]
		o.PlayerData.OpggData = t.PlayerData.OpggData

		o.LookingFor.LadderMode =  LadderCache[t.LookingFor.LadderMode]
		o.LookingFor.BuddyRank.MinRank =  RankCache[t.LookingFor.BuddyRank.MinRank]
		o.LookingFor.BuddyRank.MaxRank =  RankCache[t.LookingFor.BuddyRank.MaxRank]

		o.LookingFor.BuddyRoles.Top.RoleInfo = RoleCache[t.LookingFor.BuddyRoles.Top.RoleInfo]
		o.LookingFor.BuddyRoles.Jungle.RoleInfo = RoleCache[t.LookingFor.BuddyRoles.Jungle.RoleInfo]
		o.LookingFor.BuddyRoles.Mid.RoleInfo = RoleCache[t.LookingFor.BuddyRoles.Mid.RoleInfo]
		o.LookingFor.BuddyRoles.Adc.RoleInfo = RoleCache[t.LookingFor.BuddyRoles.Adc.RoleInfo]
		o.LookingFor.BuddyRoles.Support.RoleInfo = RoleCache[t.LookingFor.BuddyRoles.Support.RoleInfo]

		o.LookingFor.ChatUsage = t.LookingFor.ChatUsage
	}

	cur.Close(context.TODO())
	if len(res) == 0 {
		return res, mongo.ErrNoDocuments
	}

	return res, nil
}
