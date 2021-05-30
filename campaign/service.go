package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(UserId int) ([]Campaign, error)
	GetCampaignById(Input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(InputId GetCampaignDetailInput, InputData CreateCampaignInput) (Campaign, error)
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

func (s *service) GetCampaignById(Input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindById(Input.Id)
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

func (s *service) UpdateCampaign(InputId GetCampaignDetailInput, InputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindById(InputId.Id)
	if err != nil {
		return campaign, err
	}

	if campaign.UserId != InputData.User.Id {
		return campaign, errors.New("Not an owner of the campaign")
	}

	campaign.Name = InputData.Name
	campaign.ShortDescription = InputData.ShortDescription
	campaign.Description = InputData.Description
	campaign.Perks = InputData.Perks
	campaign.GoalAmount = InputData.GoalAmount

	updateCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updateCampaign, err
	}
	return updateCampaign, nil
}
