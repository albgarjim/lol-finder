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

	if RankCache, err = DBRanksCache(); err != nil {
		log.Fatal("Rank cache not initialized: ", err)
		return
	}

	if ServerCache, err = DBServerCache(); err != nil {
		log.Fatal("Rank cache not initialized: ", err)
		return
	}

	if LadderCache, err = DBLadderCache(); err != nil {
		log.Fatal("Rank cache not initialized: ", err)
		return
	}

	if RoleCache, err = DBRoleCache(); err != nil {
		log.Fatal("Rank cache not initialized: ", err)
		return
	}

	log.Info("Cache initialized!")
}