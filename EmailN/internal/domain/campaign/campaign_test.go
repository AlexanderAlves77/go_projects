package campaign_test

import (
	"emailn/internal/domain/campaign"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	campaign := campaign.NewCampaign(name, content, contacts)

	assert.Equal(campaign.Id, "1")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(campaign.Contacts, len(contacts))
}
