package models

type Student struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	Gender    string `json:"gender"`
	GroupID   int    `json:"group_id"`
	GroupName string `json:"group_name"`
}
