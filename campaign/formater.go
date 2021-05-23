package campaign

import "strings"

type CampaignFormater struct {
	Id               int    `json:"id"`
	UserId           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormater {
	campaignFormater := CampaignFormater{}
	campaignFormater.Id = campaign.Id
	campaignFormater.UserId = campaign.UserId
	campaignFormater.Name = campaign.Name
	campaignFormater.ShortDescription = campaign.ShortDescription
	campaignFormater.GoalAmount = campaign.GoalAmount
	campaignFormater.CurrentAmount = campaign.CurrentAmount
	campaignFormater.ImageURL = ""
	campaignFormater.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignFormater.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormater
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormater {

	campaignsFormater := []CampaignFormater{}

	for _, campaigns := range campaigns {
		campaignFormater := FormatCampaign(campaigns)
		campaignsFormater = append(campaignsFormater, campaignFormater)
	}

	return campaignsFormater

}

type CampaignDetailFormater struct {
	Id               int                        `json:"id"`
	Name             string                     `json:"name"`
	ShortDescription string                     `json:"short_description"`
	Description      string                     `json:"description"`
	ImageURL         string                     `json:"image_url"`
	GoalAmount       int                        `json:"goal_amount"`
	CurrentAmount    int                        `json:"current_amount"`
	UserId           int                        `json:"user_id"`
	Slug             string                     `json:"slug"`
	Perks            []string                   `json:"perks"`
	User             CampaignDetailFormaterUser `json:"user"`
	Images           []CampaignImageFormater    `json:"images"`
}
type CampaignDetailFormaterUser struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImageFormater struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormater {
	campaignDetailFormater := CampaignDetailFormater{}
	campaignDetailFormater.Id = campaign.Id
	campaignDetailFormater.Name = campaign.Name
	campaignDetailFormater.ShortDescription = campaign.ShortDescription
	campaignDetailFormater.Description = campaign.Description
	campaignDetailFormater.GoalAmount = campaign.GoalAmount
	campaignDetailFormater.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormater.UserId = campaign.UserId
	campaignDetailFormater.Slug = campaign.Slug

	campaignDetailFormater.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormater.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	campaignDetailFormater.Perks = perks

	user := campaign.User
	campaignFormaterUser := CampaignDetailFormaterUser{}
	campaignFormaterUser.Name = user.Name
	campaignFormaterUser.ImageURL = user.AvatarFileName

	campaignDetailFormater.User = campaignFormaterUser

	images := []CampaignImageFormater{}

	for _, image := range campaign.CampaignImages {
		campaignImageFormater := CampaignImageFormater{}
		campaignImageFormater.ImageURL = image.FileName
		Isprimary := false
		if image.IsPrimary == 1 {
			Isprimary = true
		}
		campaignImageFormater.IsPrimary = Isprimary
		images = append(images, campaignImageFormater)
	}

	campaignDetailFormater.Images = images

	return campaignDetailFormater
}
