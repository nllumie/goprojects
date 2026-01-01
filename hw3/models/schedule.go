package models

type Schedule struct {
	ScheduleID int    `json:"schedule_id"`
	GroupID    *int   `json:"group_id"`
	Subject    string `json:"subject"`
	TimeSlot   string `json:"time_slot"`
}
