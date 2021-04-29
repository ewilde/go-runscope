package runscope

import (
	"testing"
)

func TestListIntegration(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()
	integrations, err := client.ListIntegrations(teamID)

	if err != nil {
		t.Error(err)
	}

	if len(integrations) <= 0 {
		t.Errorf("Expected some integrations got %d", len(integrations))
	}

	if len(integrations[0].ID) <= 0 {
		t.Errorf("Expected ID got %s", integrations[0].ID)
	}

	if len(integrations[0].IntegrationType) <= 0 {
		t.Errorf("Expected integration type got %s", integrations[0].IntegrationType)
	}

	if len(integrations[0].Description) <= 0 {
		t.Errorf("Expected description got %s", integrations[0].Description)
	}

	if len(integrations[0].UUID) <= 0 {
		t.Errorf("Expected UUID got %s", integrations[0].UUID)
	}
}
