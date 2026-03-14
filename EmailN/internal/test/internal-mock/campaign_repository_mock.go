package internalmock

import (
	"emailn/internal/domain/campaign"

	"github.com/stretchr/testify/mock"
)

type CampaignRepositoryMock struct {
	mock.Mock
}

func (m *CampaignRepositoryMock) Create(campaign *campaign.Campaign) error {
	args := m.Called(campaign)
	return args.Error(0)
}

func (m *CampaignRepositoryMock) Update(campaign *campaign.Campaign) error {
	args := m.Called(campaign)
	return args.Error(0)
}

func (m *CampaignRepositoryMock) Get() ([]campaign.Campaign, error) {
	args := m.Called()
	return args.Get(0).([]campaign.Campaign), args.Error(1)
}

func (m *CampaignRepositoryMock) GetBy(id string) (*campaign.Campaign, error) {
	args := m.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*campaign.Campaign), args.Error(1)
}

func (m *CampaignRepositoryMock) Delete(campaign *campaign.Campaign) error {
	args := m.Called(campaign)
	return args.Error(0)
}

func (m *CampaignRepositoryMock) GetCampaignsToBeSent() ([]campaign.Campaign, error) {
	args := m.Called()

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]campaign.Campaign), args.Error(1)
}
