package services

import (
	"HelpNow/daos"
	"HelpNow/models"
	"errors"
	"time"
)

//SaveMeetings : "Save Meetings Information"
func SaveMeetings(meetings *models.Meetings) error {

	var err error

	count, err := daos.CheckUniqueParticipant(meetings.Participants.Email, meetings.Participants.RSVP)
	if err != nil {
		return errors.New("Can't find Uniquenessuser")
	}

	if count > 1 {
		return errors.New("Participant Already have meeting on this schedule")
	} else {
		meetings.CreationTimestamp = time.Now()
		err = daos.SaveMeetings(meetings)
		if err != nil {
			return errors.New("Service Error")
		}
	}
	return nil
}

//GetMeeting : "Get Participant Meeting"
func GetMeeting(uniqueID string) (*models.Meetings, error) {
	data, err := daos.GetMeeting(uniqueID)
	if err != nil {
		return data, nil
	}
	return data, nil
}

//GetAllMeetings : "GetAll Participant Meetings"
func GetAllMeetings(start, end string) ([]models.Meetings, error) {
	data, err := daos.GetAllMeetings(start, end)
	if err != nil {
		return data, nil
	}
	return data, nil
}

//GetAllUserMeetingsByEmail : "Get AllUser Meetings By Email"
func GetAllUserMeetingsByEmail(participant string) ([]models.Meetings, error) {
	data, err := daos.GetAllUserMeetingsByEmail(participant)
	if err != nil {
		return data, nil
	}
	return data, nil
}
