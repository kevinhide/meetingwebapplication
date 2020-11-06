package daos

import (
	"errors"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type App struct {
	Router *mux.Router
	DB     *mgo.Database
}

//Meetings : ""
type Meetings struct {
	ID                bson.ObjectId `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID          string        `json:"uniqueID" bson:"uniqueID,omitempty"`
	Title             string        `json:"title" bson:"title,omitempty"`
	Participants      Participants  `json:"participants" bson:"participants,omitempty"`
	StartTime         string        `json:"startTime" bson:"startTime,omitempty"`
	EndTime           string        `json:"endTime" bson:"endTime,omitempty"`
	CreationTimestamp time.Time     `json:"creationTimestamp" bson:"creationTimestamp,omitempty"`
}

//Participants : ""
type Participants struct {
	Name  string `json:"name" bson:"name,omitempty"`
	Email string `json:"email" bson:"email,omitempty"`
	RSVP  string `json:"rsvp" bson:"rsvp,omitempty"`
}

func (m *Meetings) createMeeting(db *mgo.Database) error {
	return errors.New("Not implemented")
}
