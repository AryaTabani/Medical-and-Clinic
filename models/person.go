package models

import "time"

type User struct {
	ID            int64
	First_name    string
	Last_name     string
	Password      string
	Email         string
	Avatar        string
	Phone         string
	Rate          int64
	Job           string
	Token         string
	Refresh_Token string
	Created_at    time.Time
	Updated_at    time.Time
	User_id       string
}
