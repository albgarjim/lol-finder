package handlers

import (
	"net/http"
	// "encoding/json"

	inp "lol-api/api/v1/input"
	out "lol-api/api/v1/output"

	mdb "lol-api/api/v1/models/mongodb"

	log "github.com/sirupsen/logrus"
)

func GetDuoList(_w http.ResponseWriter, _r *http.Request) {
	log.Info("-------- Call to TestGetAllColl route --------")
	var err error
	var params *inp.URLParams

	if params, err = inp.ExtractParameters(_r); err != nil {
		log.Error("Incorrect url parameters: ", err)
		out.RespondWithError(_w, err)
		return
	}

	res, err := mdb.DBGetDuoList(params)
	if err != nil {
		log.Error("Error getting all collections from database: ", err)
		out.RespondWithError(_w, err)
		return
	}
	log.Info("-------- Finish TestGetAllColl route --------")
	out.RespondWithJSON(_w, http.StatusOK, res)
}
