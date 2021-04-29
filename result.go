/*
Package runscope implements a client library for the runscope api (https://api.blazemeter.com/api-monitoring/)

*/
package runscope

import "fmt"

type Result struct {
	AssertionsDefined int       `json:"assertions_defined"`
	AssertionsFailed  int       `json:"assertions_failed"`
	AssertionsPassed  int       `json:"assertions_passed"`
	BucketKey         string    `json:"bucket_key"`
	FinishedAt        float64   `json:"finished_at"`
	Region            string    `json:"region"`
	RequestsExecuted  int       `json:"requests_executed"`
	Result            string    `json:"result"`
	ScriptsDefined    int       `json:"scripts_defined"`
	ScriptsFailed     int       `json:"scripts_failed"`
	ScriptsPassed     int       `json:"scripts_passed"`
	StartedAt         float64   `json:"started_at"`
	TestRunID         string    `json:"test_run_id"`
	TestRunURL        string    `json:"test_run_url"`
	TestID            string    `json:"test_id"`
	VariablesDefined  int       `json:"variables_defined"`
	VariablesFailed   int       `json:"variables_failed"`
	VariablesPassed   int       `json:"variables_passed"`
	EnvironmentID     string    `json:"environment_id"`
	EnvironmentName   string    `json:"environment_name"`
	Requests          []Request `json:"requests"`
}

// Request represents the result of a request made by a given test
type Request struct {
	Result            string      `json:"result"`
	URL               string      `json:"url"`
	Method            string      `json:"method"`
	AssertionsDefined int         `json:"assertions_defined"`
	AssertionsFailed  int         `json:"assertions_failed"`
	AssertionsPassed  int         `json:"assertions_passed"`
	ScriptsDefined    int         `json:"scripts_defined"`
	ScriptsFailed     int         `json:"scripts_failed"`
	ScriptsPassed     int         `json:"scripts_passed"`
	VariablesDefined  int         `json:"variables_defined"`
	VariablesFailed   int         `json:"variables_failed"`
	VariablesPassed   int         `json:"variables_passed"`
	Assertions        []Assertion `json:"assertions"`
	Scripts           []Script    `json:"scripts"`
	Variables         []Variable  `json:"variables"`
}

// ListResults list all results for test. https://api.blazemeter.com/api-monitoring/#test-result-list
func (client *Client) ListResults(bucketKey string, testID string) ([]*Result, error) {
	path := fmt.Sprintf("/buckets/%s/tests/%s/results", bucketKey, testID)
	resource, err := client.readResource("[]result", testID, path)
	if err != nil {
		return nil, err
	}

	readResources, error := getResultsFromResponse(resource.Data)
	if error != nil {
		return nil, error
	}

	return readResources, nil
}

// ReadTestResult list details about an existing result. https://api.blazemeter.com/api-monitoring/#test-result-detail
func (client *Client) ReadTestResult(testRunID string, bucketKey string, testID string) (*Result, error) {
	path := fmt.Sprintf("/buckets/%s/tests/%s/results/%s", bucketKey, testID, testRunID)
	resource, error := client.readResource("result", testRunID, path)
	if error != nil {
		return nil, error
	}

	readTestResult, error := getResultFromResponse(resource.Data)
	if error != nil {
		return nil, error
	}

	return readTestResult, nil
}

// ReadTestLatestResult list details about an existing latest result. https://api.blazemeter.com/api-monitoring/#test-result-step-detail
func (client *Client) ReadTestLatestResult(testID string, bucketKey string) (*Result, error) {
	return client.ReadTestResult("latest", bucketKey, testID)
}

// ReadTestStepResult list details about an existing result. https://api.blazemeter.com/api-monitoring/#test-result-step-detail
func (client *Client) ReadTestStepResult(testRunID string, bucketKey string, testID string, testStepID string) (*Result, error) {
	path := fmt.Sprintf("/buckets/%s/tests/%s/results/%s/steps/%s", bucketKey, testID, testRunID, testStepID)
	resource, error := client.readResource("result", testRunID, path)
	if error != nil {
		return nil, error
	}

	readTestResult, error := getResultFromResponse(resource.Data)
	if error != nil {
		return nil, error
	}

	return readTestResult, nil
}

func getResultsFromResponse(response interface{}) ([]*Result, error) {
	var results []*Result
	err := decode(&results, response)
	return results, err
}

func getResultFromResponse(response interface{}) (*Result, error) {
	testResult := new(Result)
	err := decode(testResult, response)
	return testResult, err
}
