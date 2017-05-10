package runscope

import (
	"fmt"
	"time"

)

type Test struct {
	Id                   string    `json:"id,omitempty"`
	Bucket               *Bucket   `json:"-"`
	Name                 string    `json:"name,omitempty"`
	Description          string    `json:"description,omitempty"`
	CreatedAt            time.Time `json:"created_at,omitempty"`
	CreatedBy            Contact   `json:"created_by,omitempty"`
	DefaultEnvironmentId string    `json:"default_environment_id,omitempty"`
	ExportedAt           time.Time `json:"exported_at,omitempty"`
}

type Contact struct {
	Email         string     `json:"email,omitempty"`
        Id            string     `json:"id"`
        Name          string     `json:"name,omitempty"`
}

func NewTest() *Test {
	return &Test { Bucket: &Bucket{}}
}

func (client *Client) CreateTest(test *Test) (*Test, error) {
	newResource, error := client.createResource(test, "test", test.Name,
		fmt.Sprintf("/buckets/%s/tests", test.Bucket.Key))
	if error != nil {
		return nil, error
	}

	newTest, error := getTestFromResponse(newResource.Data)
	if error != nil {
		return nil, error
	}

	newTest.Bucket = test.Bucket
	return newTest, nil
}

func (client *Client) ReadTest(test *Test) (*Test, error) {
	resource, error := client.readResource("test", test.Id, fmt.Sprintf("/buckets/%s/tests/%s", test.Bucket.Key, test.Id))
	if error != nil {
		return nil, error
	}

	readTest, error := getTestFromResponse(resource.Data)
	if error != nil {
		return nil, error
	}

	readTest.Bucket = test.Bucket
	return readTest, nil
}

func (client *Client) UpdateTest(test *Test) (response, error) {
	resource, error := client.updateResource(test, "test", test.Id, fmt.Sprintf("/buckets/%s/tests/%s", test.Bucket.Key, test.Id))
	return resource.(response), error
}

func (client *Client) DeleteTest(test *Test) error {
	return client.deleteResource("test", test.Id, fmt.Sprintf("/buckets/%s/tests/%s", test.Bucket.Key, test.Id))
}

func getTestFromResponse(response interface{}) (*Test, error) {
	test := new(Test)
	err := Decode(test, response)
	return test, err
}