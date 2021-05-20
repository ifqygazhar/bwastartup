package campaign

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
