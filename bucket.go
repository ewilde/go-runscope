package runscope

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
)

type Bucket struct {
	Name           string `json:"name,omitempty"`
	Key            string `json:"key,omitempty"`
	Default        bool   `json:"default,omitempty"`
	AuthToken      string `json:"auth_token,omitempty"`
	TestsUrl       string `json:"tests_url,omitempty" mapstructure:"tests_url"`
	CollectionsUrl string `json:"collections_url,omitempty"`
	MessagesUrl    string `json:"messages_url,omitempty"`
	TriggerUrl     string `json:"trigger_url,omitempty"`
	VerifySsl      bool   `json:"verify_ssl,omitempty"`
	Team           Team   `json:"team,omitempty"`
}

func (client *Client) CreateBucket(bucket Bucket) (*Bucket, error) {
	log.Printf("[DEBUG] creating bucket %s", bucket.Name)
	data := url.Values{}
	data.Add("name", bucket.Name)
	data.Add("team_uuid", bucket.Team.Id)

	log.Printf("[DEBUG] 	request: POST %s %#v", "/buckets", data)

	req, err := client.newFormUrlEncodedRequest("POST", "/buckets", data)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] %#v", req)
	resp, err := client.Http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	log.Printf("[DEBUG] 	response: %d %s", resp.StatusCode, bodyString)

	if resp.StatusCode >= 300 {
		errorResp := new(errorResponse)
		if err = json.Unmarshal(bodyBytes, &errorResp); err != nil {
			return nil, fmt.Errorf("Error creating bucket: %s", bucket.Name)
		} else {
			return nil, fmt.Errorf("Error creating bucket: %s, status: %d reason: %q", bucket.Name,
				errorResp.Status, errorResp.ErrorMessage)
		}
	} else {
		response := new(response)
		json.Unmarshal(bodyBytes, &response)
		return getBucketFromResponse(response.Data)
	}
}

func (client *Client) ReadBucket(key string) (*Bucket, error) {
	resource, error := client.readResource("bucket", key, fmt.Sprintf("/buckets/%s", key))
	if error != nil {
		return nil, error
	}

	bucket, error := getBucketFromResponse(resource.Data)
	return bucket, error
}

func (client *Client) DeleteBucket(key string) error {
	return client.deleteResource("bucket", key, fmt.Sprintf("/buckets/%s", key))
}

func getBucketFromResponse(response interface{}) (*Bucket, error) {
	bucket := new(Bucket)
	err := Decode(bucket, response)
	return bucket, err
}
