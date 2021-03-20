package runscope

// Regions represents multiple Runscope regions
type Regions struct {
	Regions []Region `json:"regions"`
}

// Region represents a Runscope region
type Region struct {
	RegionCode      string `json:"region_code"`
	Location        string `json:"location"`
	ServiceProvider string `json:"service_provider"`
	Hostname        string `json:"hostname"`
}

// ListRegions returns all regions known by Runscope https://api.blazemeter.com/api-monitoring/#regions
func (client *Client) ListRegions() (*Regions, error) {
	resource, err := client.readResource("[]regions", "regions", "/regions")
	if err != nil {
		return nil, err
	}

	readResources, error := getRegionsFromResponse(resource.Data)
	if error != nil {
		return nil, error
	}

	return readResources, nil
}

func getRegionsFromResponse(response interface{}) (*Regions, error) {
	var results *Regions
	err := decode(&results, response)
	return results, err
}
