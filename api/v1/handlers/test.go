package handlers

import (
	"net/http"
	"encoding/json"

	inp "lol-api/api/v1/input"
	out "lol-api/api/v1/output"

	mdb "lol-api/api/v1/models/mongodb"

	log "github.com/sirupsen/logrus"
)

func TestGetAllColl(_w http.ResponseWriter, _r *http.Request) {
	log.Info("-------- Call to TestGetAllColl route --------")
	var err error
	var params *inp.URLParams

	if params, err = inp.ExtractParameters(_r); err != nil {
		log.Error("Incorrect url parameters: ", err)
		out.RespondWithError(_w, err)
		return
	}

	res, err := mdb.TestGetAllCollDB(params)
	if err != nil {
		log.Error("Error getting all collections from database: ", err)
		out.RespondWithError(_w, err)
		return
	}
	log.Info("-------- Finish TestGetAllColl route --------")
	out.RespondWithJSON(_w, http.StatusOK, res)
}


func TestGetCollById(_w http.ResponseWriter, _r *http.Request) {
	log.Info("-------- Call to TestGetCollById route --------")
	var err error
	var params *inp.URLParams

	if params, err = inp.ExtractParameters(_r); err != nil {
		log.Error("Incorrect url parameters: ", err)
		out.RespondWithError(_w, err)
		return
	}

	res, err := mdb.TestGetCollByIdDB(params)
	if err != nil {
		log.Error("Error getting collection from database: ", err)
		out.RespondWithError(_w, err)
		return
	}
	log.Info("-------- Finish TestGetCollById route --------")
	out.RespondWithJSON(_w, http.StatusOK, res)
}

func TestDeleteCollById(_w http.ResponseWriter, _r *http.Request) {
	log.Info("-------- Call to TestDeleteCollById route --------")
	var err error
	var params *inp.URLParams

	if params, err = inp.ExtractParameters(_r); err != nil {
		log.Error("Incorrect url parameters: ", err)
		out.RespondWithError(_w, err)
		return
	}

	res, err := mdb.TestDeleteCollByIdDB(params)
	if err != nil {
		log.Error("Error deleting collection from database: ", err)
		out.RespondWithError(_w, err)
		return
	}
	log.Info("-------- Finish TestDeleteCollById route --------")
	out.RespondWithJSON(_w, http.StatusOK, res)
}

func TestCreateCollById(_w http.ResponseWriter, _r *http.Request) {
	log.Info("-------- Call to TestCreateCollById route --------")
	var err error
	var params *inp.URLParams

	if params, err = inp.ExtractParameters(_r); err != nil {
		log.Error("Incorrect url parameters: ", err)
		out.RespondWithError(_w, err)
		return
	}

	res, err := mdb.TestCreateCollByIdDB(params)
	if err != nil {
		log.Error("Error creating collection from database: ", err)
		out.RespondWithError(_w, err)
		return
	}
	log.Info("-------- Finish TestCreateCollById route --------")
	out.RespondWithJSON(_w, http.StatusOK, res)
}

func TestUpdateCollById(_w http.ResponseWriter, _r *http.Request) {
	log.Info("-------- Call to TestUpdateCollById route --------")
	var err error
	var params *inp.URLParams

	if params, err = inp.ExtractParameters(_r); err != nil {
		log.Error("Incorrect url parameters: ", err)
		out.RespondWithError(_w, err)
		return
	}

	var u inp.TestObj

	if err := json.NewDecoder(_r.Body).Decode(&u); err != nil {
		log.Error("Error decoding body to wishlist struct: ", err)
		out.RespondWithError(_w, out.ErrBodyParams)
		return
	}
	defer _r.Body.Close()

	if err := u.ValidateFields(); err != nil {
		log.Error("Invalid data in wishlist struct: ", err)
		out.RespondWithError(_w, out.ErrBodyParams)
		return
	}

	res, err := mdb.TestUpdateCollByIdDB(params, u)
	if err != nil {
		log.Error("Error updating collection from database: ", err)
		out.RespondWithError(_w, err)
		return
	}
	log.Info("-------- Finish TestUpdateCollById route --------")
	out.RespondWithJSON(_w, http.StatusOK, res)
}
