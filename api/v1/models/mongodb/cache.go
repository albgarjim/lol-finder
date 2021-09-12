package mongo

import (
	out "lol-api/api/v1/output"
	// "sync"

	log "github.com/sirupsen/logrus"

)

var RankCache = map[string]out.RankObj{}
// var RankCacheMutex = sync.RWMutex{}

var ServerCache = map[string]out.ServerObj{}

var LadderCache = map[string]out.LadderObj{}

var RoleCache = map[string]out.RoleObj{}

func InitCache() {
	var err error

	if RankCache, err = DBGetRanks(); err != nil {
		log.Fatal("Rank cache not initialized: ", err)
		return
	}

	if ServerCache, err = DBGetServers(); err != nil {
		log.Fatal("Rank cache not initialized: ", err)
		return
	}

	if LadderCache, err = DBGetLadders(); err != nil {
		log.Fatal("Rank cache not initialized: ", err)
		return
	}

	if RoleCache, err = DBGetRoles(); err != nil {
		log.Fatal("Rank cache not initialized: ", err)
		return
	}

	log.Info("Cache initialized!")
}