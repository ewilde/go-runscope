package runscope

import "fmt"

type Test struct {
	Id            string `json:"id,omitempty"`
	BucketId      string `json:"-"`
	Name          string `json:"name"`
	Description   string `json:"description"`
}

func (client *Client) CreateTest(test Test) (Test, error) {
	id, error := client.createResource(test, "test", test.Name, "id",
		fmt.Sprintf("/buckets/%s/tests", test.BucketId))
	if error != nil {
		return test, error
	}

	test.Id = id
	return test, nil
}

func (client *Client) ReadTest(test Test) (response, error) {
	resource, error := client.readResource(response{}, "test", test.Id, fmt.Sprintf("/buckets/%s/tests/%s", test.BucketId, test.Id))
	return resource.(response), error
}

func (client *Client) UpdateTest(test Test) (response, error) {
	resource, error := client.updateResource(test, "test", test.Id, fmt.Sprintf("/buckets/%s/tests/%s", test.BucketId, test.Id))
	return resource.(response), error
}

func (client *Client) DeleteTest(test Test) error {
	return client.deleteResource("test", test.Id, fmt.Sprintf("/buckets/%s/tests/%s", test.BucketId, test.Id))
}