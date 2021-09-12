package mongo

import (
	"reflect"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	log "github.com/sirupsen/logrus"
)

func GetIDs(t interface{}) ([]primitive.ObjectID) {
	var ids []primitive.ObjectID
    switch reflect.TypeOf(t).Kind() {
    case reflect.Slice:
        s := reflect.ValueOf(t)

        for i := 0; i < s.Len(); i++ {
			log.Info("----", s.Index(i))
			id := s.Index(i).Interface().(string)
			objectId, err := primitive.ObjectIDFromHex(id)
			if err != nil{
				log.Println("Invalid id")
			}
			ids = append(ids, objectId)
        }
    }
	return ids
}

func GetSingleID(t interface{}) (primitive.ObjectID, error) {
	str := fmt.Sprintf("%v", t)
	objectId, err := primitive.ObjectIDFromHex(str)
	if err != nil{
		log.Println("Invalid id")
	}
	return objectId, err
}
