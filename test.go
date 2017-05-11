package runscope

import (
	"fmt"
	"time"
)

// Test represents the details for a runscope test. See https://www.runscope.com/docs/api/tests
type Test struct {
	ID                   string     `json:"id,omitempty"`
	Bucket               *Bucket    `json:"-"`
	Name                 string     `json:"name,omitempty"`
	Description          string     `json:"description,omitempty"`
	CreatedAt            time.Time  `json:"created_at,omitempty"`
	CreatedBy            *Contact   `json:"created_by,omitempty"`
	DefaultEnvironmentID string     `json:"default_environment_id,omitempty"`
	ExportedAt           time.Time  `json:"exported_at,omitempty"`
}

// Contact details
type Contact struct {
	Email string     `json:"email,omitempty"`
        ID    string     `json:"id"`
        Name  string     `json:"name,omitempty"`
}

// NewTest creates a new test struct
func NewTest() *Test {
	return &Test { Bucket: &Bucket{}}
}

// CreateTest creates a new runscope test. See https://www.runscope.com/docs/api/tests#create
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

// ReadTest list details about an existing test. See https://www.runscope.com/docs/api/tests#detail
func (client *Client) ReadTest(test *Test) (*Test, error) {
	resource, error := client.readResource("test", test.ID, fmt.Sprintf("/buckets/%s/tests/%s", test.Bucket.Key, test.ID))
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

// UpdateTest update an existing test. See https://www.runscope.com/docs/api/tests#modifying
func (client *Client) UpdateTest(test *Test) (*Test, error) {
	resource, error := client.updateResource(test, "test", test.ID, fmt.Sprintf("/buckets/%s/tests/%s", test.Bucket.Key, test.ID))
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

// DeleteTest delete an existing test. See https://www.runscope.com/docs/api/tests#delete
func (client *Client) DeleteTest(test *Test) error {
	return client.deleteResource("test", test.ID, fmt.Sprintf("/buckets/%s/tests/%s", test.Bucket.Key, test.ID))
}

func getTestFromResponse(response interface{}) (*Test, error) {
	test := new(Test)
	err := decode(test, response)
	return test, err
}