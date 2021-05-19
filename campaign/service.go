package campaign

type Service interface {
	FindCampaigns(UserId int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindCampaigns(UserId int) ([]Campaign, error) {
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
