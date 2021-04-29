package runscope

// Account represents Runscope account
type Account struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    string `json:"id"`
	UUID  string `json:"uuid"`
	Teams []Team `json:"teams"`
}

// GetAccount returns runscope account information https://api.blazemeter.com/api-monitoring/#account
func (client *Client) GetAccount() (*Account, error) {
	resource, err := client.readResource("account", "account", "/account")
	if err != nil {
		return nil, err
	}

	account, err := getAccountFromResponse(resource.Data)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func getAccountFromResponse(response interface{}) (*Account, error) {
	var results *Account
	err := decode(&results, response)
	return results, err
}
