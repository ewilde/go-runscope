package runscope

import (
	"testing"
)

func TestGetAccount(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()

	account, err := client.GetAccount()
	if err != nil {
		t.Error(err)
	}
	if account == nil {
		t.Error("Expected to get account information, got nothing")
	}
}
