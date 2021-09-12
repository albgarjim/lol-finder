package input

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

const (
	QColId      = "id"
	QColList    = "idList"
)

var list = []string{QColId, QColList}

type URLParams = map[string]interface{}

//ExtractParameters gets the  parameters form the query and stores them in a map
func ExtractParameters(_r *http.Request) (*URLParams, error) {
	vars := mux.Vars(_r)
	params := &URLParams{}

	for _, key := range list {
		val := valid[key](_r.URL.Query().Get(key))
		if val != "" {
			(*params)[key] = val
		}
	}

	for key, val := range vars {
		(*params)[key] = valid[key](val)
	}

	for key, val := range *params {
		log.Info("query: ", key, val)
	}

	return params, nil
}
