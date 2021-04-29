package runscope

import (
	"testing"
)

func TestListRegions(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()

	listRegions, err := client.ListRegions()
	if err != nil {
		t.Error(err)
	}

	if len(listRegions.Regions) == 0 {
		t.Error("Expected results but none found")
	}
}
