package runscope

import (
	"log"
	"net/url"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type Bucket struct {
	Id   string
	Name string
	Team Team
}

func (client *Client) CreateBucket(bucket Bucket) (string, error) {
	log.Printf("[DEBUG] creating bucket %s", bucket.Name)
	data := url.Values{}
	data.Add("name", bucket.Name)
	data.Add("team_uuid", bucket.Team.Id)

	log.Printf("[DEBUG] 	request: POST %s %#v", "/buckets", data)

	req, err := client.newFormUrlEncodedRequest("POST", "/buckets", data)
	if err != nil {
		return "", err
	}

	resp, err := client.Http.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	log.Printf("[DEBUG] 	response: %d %s", resp.StatusCode, bodyString)

	if resp.StatusCode >= 300 {
		errorResp := new(errorResponse)
		if err = json.Unmarshal(bodyBytes, &errorResp); err != nil {
			return "", fmt.Errorf("Error creating bucket: %s", bucket.Name)
		} else {
			return "", fmt.Errorf("Error creating bucket: %s, status: %d reason: %q", bucket.Name,
				errorResp.Status, errorResp.ErrorMessage)
		}
	} else {
		response := new(response)
		json.Unmarshal(bodyBytes, &response)
		return response.Data["key"].(string), nil
	}
}

func (client *Client) ReadBucket(key string) (response, error) {
	resource, error := client.readResource(response{}, "bucket", key, fmt.Sprintf("/buckets/%s", key))
	return resource.(response), error
}

func (client *Client) DeleteBucket(key string) error {
	return client.deleteResource("bucket", key, fmt.Sprintf("/buckets/%s", key))
}
