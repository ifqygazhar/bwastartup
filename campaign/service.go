package campaign

type Service interface {
	GetCampaigns(UserId int) ([]Campaign, error)
	GetCampaignById(Input GetCampaignDetailInput) (Campaign, error)
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
