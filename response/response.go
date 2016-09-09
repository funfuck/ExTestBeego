package response

import "time"

type Response struct {
	Success bool
	Desc string
}

type ResponseMember struct{
	Success bool
	Desc string
	Result interface{}
}

type MemberHistory struct {
	ID              uint `gorm:"primary_key"`
	FirstName       string
	LastName        string
	Email           string
	TimeStamp       time.Time `gorm:"timestamp"`
}