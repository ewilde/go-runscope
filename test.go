package runscope

import (
	"fmt"
	"time"

)

type Test struct {
	Id            string    `json:"id,omitempty"`
	Bucket        *Bucket   `json:"-"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
}

func (client *Client) CreateTest(test Test) (Test, error) {
	id, error := client.createResource(test, "test", test.Name, "id",
		fmt.Sprintf("/buckets/%s/tests", test.Bucket.Key))
	if error != nil {
		return test, error
	}

	test.Id = id
	return test, nil
}

func (client *Client) ReadTest(test Test) (*Test, error) {
	resource, error := client.readResource("test", test.Id, fmt.Sprintf("/buckets/%s/tests/%s", test.Bucket.Key, test.Id))
	if error != nil {
		return nil, error
	}

	return getTestFromResponse(resource.Data)
}

func (client *Client) UpdateTest(test Test) (response, error) {
	resource, error := client.updateResource(test, "test", test.Id, fmt.Sprintf("/buckets/%s/tests/%s", test.Bucket.Key, test.Id))
	return resource.(response), error
}

func (client *Client) DeleteTest(test Test) error {
	return client.deleteResource("test", test.Id, fmt.Sprintf("/buckets/%s/tests/%s", test.Bucket.Key, test.Id))
}

func getTestFromResponse(response interface{}) (*Test, error) {
	test := new(Test)
	err := Decode(test, response)
	return test, err
}