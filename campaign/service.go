package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(UserId int) ([]Campaign, error)
	GetCampaignById(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(inputId GetCampaignDetailInput, InputData CreateCampaignInput) (Campaign, error)
	SaveCampaignimage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(UserId int) ([]Campaign, error) {
	if UserId != 0 {
		campaign, err := s.repository.FindByUserId(UserId)
		if err != nil {
			return campaign, err
		}

		return campaign, nil
	}

	campaign, err := s.repository.FindAll()
	if err != nil {
		return campaign, err
	}

	return campaign, nil

}

func (s *service) GetCampaignById(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindById(input.Id)
	if err != nil {
		return campaign, err
	}

	return campaign, err
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserId = input.User.Id
	stringSlug := fmt.Sprint(input.Name, "-", input.User.Id)
	campaign.Slug = slug.Make(stringSlug)

	//buat slug

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

func (s *service) UpdateCampaign(InputId GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindById(InputId.Id)
	if err != nil {
		return campaign, err
	}

	if campaign.UserId != inputData.User.Id {
		return campaign, errors.New("Not an owner of the campaign")
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount

	updateCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updateCampaign, err
	}
	return updateCampaign, nil
}

func (s *service) SaveCampaignimage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {
	campaign, err := s.repository.FindById(input.CampaignId)
	if err != nil {
		return CampaignImage{}, err
	}
	if campaign.UserId != input.User.Id {
		return CampaignImage{}, errors.New("Not an owner of the campaign")
	}
	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1
		_, err := s.repository.MarkAllImageAsNotPrimary(input.CampaignId)
		if err != nil {
			return CampaignImage{}, err
		}
	}

	campaignImage := CampaignImage{}
	campaignImage.CampaignId = input.CampaignId
	campaignImage.IsPrimary = isPrimary
	campaignImage.FileName = fileLocation

	newCampaignImage, err := s.repository.CreateImage(campaignImage)
	if err != nil {
		return newCampaignImage, err
	}
	return newCampaignImage, nil
}
