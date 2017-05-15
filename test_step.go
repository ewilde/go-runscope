package runscope

import "fmt"

// NewTest creates a new test struct
func NewTestStep() *TestStep {
	return &TestStep {}
}

// CreateTest creates a new runscope test step. See https://www.runscope.com/docs/api/steps#add
func (client *Client) CreateTestStep(testStep *TestStep, bucketKey string, testId string) (*TestStep, error) {
	newResource, error := client.createResource(testStep, "testStep", testStep.ID,
		fmt.Sprintf("/buckets/%s/tests/%s/steps", bucketKey, testId))
	if error != nil {
		return nil, error
	}

	newTestStep, error := getTestStepFromResponse(newResource.Data)
	if error != nil {
		return nil, error
	}

	return newTestStep, nil
}

// ReadTest list details about an existing test. See https://www.runscope.com/docs/api/tests#detail
func (client *Client) ReadTestStep(testStep *TestStep, bucketKey string, testId string) (*TestStep, error) {
	resource, error := client.readResource("testStep", testStep.ID,
		fmt.Sprintf("/buckets/%s/tests/%s/steps/%s", bucketKey, testId, testStep.ID))
	if error != nil {
		return nil, error
	}

	readTestStep, error := getTestStepFromResponse(resource.Data)
	if error != nil {
		return nil, error
	}

	return readTestStep, nil
}

// UpdateTest update an existing test. See https://www.runscope.com/docs/api/tests#modifying
func (client *Client) UpdateTestStep(testStep *TestStep, bucketKey string, testId string) (*TestStep, error) {
	resource, error := client.updateResource(testStep, "testStep", testStep.ID,
		fmt.Sprintf("/buckets/%s/tests/%s/steps/%s", bucketKey, testId, testStep.ID))
	if error != nil {
		return nil, error
	}

	readTestStep, error := getTestStepFromResponse(resource.Data)
	if error != nil {
		return nil, error
	}

	return readTestStep, nil
}

// DeleteTest delete an existing test. See https://www.runscope.com/docs/api/tests#delete
func (client *Client) DeleteTestStep(testStep *TestStep, bucketKey string, testId string) error {
	return client.deleteResource("testStep", testStep.ID,
		fmt.Sprintf("/buckets/%s/tests/%s/steps/%s", bucketKey, testId, testStep.ID))
}

func getTestStepFromResponse(response interface{}) (*TestStep, error) {
	testStep := new(TestStep)
	err := decode(testStep, response)
	return testStep, err
}