package transaction

import "time"

type Transaction struct {
	Id         int
	CampaignId int
	UserId     int
	Amount     int
	Status     string
	Code       string
	CreateAt   time.Time
	UpdateAt   time.Time
}
