package transaction

import (
	"bwastartup/user"
	"time"
)

type Transaction struct {
	Id         int
	CampaignId int
	UserId     int
	Amount     int
	Status     string
	Code       string
	User       user.User
	CreateAt   time.Time
	UpdateAt   time.Time
}
