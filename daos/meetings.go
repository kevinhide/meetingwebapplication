package daos

import (
	"HelpNow/models"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"
)

//SaveMeetings : "SaveMeetings by Json Input"
func SaveMeetings(meetings *models.Meetings) error {
	db := GetDB()
	defer db.Session.Close()
	err := db.C(models.COLLECTIONNAME).Insert(meetings)
	if err != nil {
		return err
	}
	return nil
}

//GetMeeting : "Get Meeting using Particular UniqueID"
func GetMeeting(uniqueID string) (*models.Meetings, error) {
	db := GetDB()
	defer db.Session.Close()

	var tempmeet *models.Meetings

	query := bson.M{"uniqueID": uniqueID}

	BsonToJSONPrint(query)

	err := db.C(models.COLLECTIONNAME).Find(query).One(&tempmeet)
	if err != nil {
		return nil, nil
	}
	return tempmeet, nil
}

//GetAllMeetings : "To Get Meetings Information by Two Date Ranges"
func GetAllMeetings(start, end string) ([]models.Meetings, error) {
	db := GetDB()
	defer db.Session.Close()

	var tempMeet []models.Meetings

	gtQuery := bson.M{"$gte": start}

	ltQuery := bson.M{"$lte": end}

	pipe := bson.M{
		"$and": []bson.M{
			bson.M{"startTime": gtQuery},
			bson.M{"endTime": ltQuery},
		},
	}

	BsonToJSONPrint(pipe)

	err := db.C(models.COLLECTIONNAME).Find(pipe).All(&tempMeet)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New(err.Error())
	}
	return tempMeet, nil
}

//GetAllUserMeetingsByEmail : "To Get All Meetings by User Email"
func GetAllUserMeetingsByEmail(participant string) ([]models.Meetings, error) {
	db := GetDB()
	defer db.Session.Close()

	var tempMeet []models.Meetings

	pipe := bson.M{
		"$and": []bson.M{
			bson.M{"participants.email": participant},
		},
	}

	BsonToJSONPrint(pipe)

	err := db.C(models.COLLECTIONNAME).Find(pipe).All(&tempMeet)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New(err.Error())
	}
	return tempMeet, nil
}

//CheckUniqueParticipant : ""
func CheckUniqueParticipant(email, RSVP string) (int, error) {

	db := GetDB()
	defer db.Session.Close()

	query := bson.M{
		"$or": []bson.M{
			bson.M{"participants.email": email},
			bson.M{"participants.RSVP": RSVP},
		},
	}

	BsonToJSONPrint(query)

	count, err := db.C(models.COLLECTIONNAME).Find(query).Count()
	log.Println("count==>", count)
	return count, err
}

//BsonToJSONPrint : ""
func BsonToJSONPrint(d interface{}) {
	b, err := json.Marshal(d)
	fmt.Println("err", err, string(b))
}
