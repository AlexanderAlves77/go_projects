package campaign

import (
	"emailn/internal/domain/campaign"
	"testing"
)

func TestNewCampaign(t *testing.T) {
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	compaign := NewCampaign(name, content, contacts)

	if campaign.Id != "1" {
		t.Errorf("excepted 1", compaign.Id)
	} else if campaign.Name != name {
		t.Errorf("excepted correct name")
	} else if campaign.Content != content {
		t.Errorf("excepted correct content")
	} else if len(campaign.Contact) != len(contacts) {
		t.Errorf("excepted correct contacts")
	}
}
