package input

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

var valid = map[string]func(string) interface{}{
	QColId:      valColId,
	QColList:    valColList,
}

func valColId(_val string) interface{} {
	if _val == "" {
		log.Info("Filtered colId")
		return ""
	}
	return _val
}

func valColList(_val string) interface{} {
	log.Info("collection list: ", _val)
	if _val == "" {
		return [...]string{""}
	}
	return strings.Split(_val, " ")
}
