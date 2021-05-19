package campaign

import "time"

type Campaign struct {
	Id               int
	UserId           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	Slug             string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	CreateAt         time.Time
	UpdateAt         time.Time
	CampaignImages   []CampaignImage
}

type CampaignImage struct {
	Id         int
	CampaignId int
	FileName   string
	IsPrimary  int
	CreateAt   time.Time
	UpdateAt   time.Time
}
