package transaction

import (
	"bwastartup/campaign"
	"errors"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionByCampaignId(input GetTransactionCampaignDetailInput) ([]Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionByCampaignId(input GetTransactionCampaignDetailInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindById(input.Id)
	if err != nil {
		return []Transaction{}, err
	}
	if campaign.UserId != input.User.Id {
		return []Transaction{}, errors.New("Not an owner of the campaign")
	}
	transaction, err := s.repository.GetByCampaignId(input.Id)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
