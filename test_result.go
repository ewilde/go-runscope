package runscope

import (
	"encoding/json"
	"fmt"
	"time"
)



// TestResult a list of all results for a given test.
// See https://www.runscope.com/docs/api/results#test-run-list
type TestResult struct {
	Agent             string     `json:"agent,omitempty"`
	AssertionsDefined uint       `json:"assertions_defined,omitempty"`
	AssertionsFailed  uint       `json:"assertions_failed,omitempty"`
	AssertionsPassed  uint       `json:"assertions_passed,omitempty"`
	BucketKey         string     `json:"bucket_key,omitempty"`
	FinishedAt        *time.Time `json:"finished_at,omitempty"`
	Region            string     `json:"region,omitempty"`
	RequestsExecuted  uint       `json:"requests_executed,omitempty"`
	Result            string     `json:"result,omitempty"`
	ScriptsDefined    uint       `json:"scripts_defined,omitempty"`
	ScriptsFailed     uint       `json:"scripts_failed,omitempty"`
	ScriptsPassed     uint       `json:"scripts_passed,omitempty"`
	StartedAt         *time.Time `json:"started_at,omitempty"`
	TestRunID         string     `json:"test_run_id,omitempty"`
	TestRunURL        string     `json:"test_run_url,omitempty"`
	TestID            string     `json:"test_id,omitempty"`
	VariablesDefined  uint       `json:"variables_defined,omitempty"`
	VariablesFailed   uint       `json:"variables_failed,omitempty"`
	VariablesPassed   uint       `json:"variables_passed,omitempty"`
	EnvironmentID     string     `json:"environment_id,omitempty"`
	EnvironmentName   string     `json:"environment_name,omitempty"`
}

// ReadTestResults a list of all results for a given test. See https://www.runscope.com/docs/api/results#test-run-list
func (client *Client) ReadTestResults(bucketKey string, testID string, count uint, since time.Time, before time.Time) ([]*TestResult, error) {
	resource, err := client.readResource("test results", testID,
		fmt.Sprintf("/buckets/%s/tests/%s/results?count=%d&since=%f&before=%f",
			bucketKey, testID, count,
			timeToFloat(since),
			timeToFloat(before)))
	if err != nil {
		return nil, err
	}

	readTestResult, err := getTestResultFromResponse(resource.Data)
	if err != nil {
		return nil, err
	}

	return readTestResult, nil
}

// ReadTestResults a list of all results for a given test. See https://www.runscope.com/docs/api/results#test-run-list
func (client *Client) ReadTestResultsLatest(bucketKey string, testID string, count uint) ([]*TestResult, error) {
	resource, err := client.readResource("test results", testID,
		fmt.Sprintf("/buckets/%s/tests/%s/results?count=%d", bucketKey, testID, count))
	if err != nil {
		return nil, err
	}

	readTestResult, err := getTestResultFromResponse(resource.Data)
	if err != nil {
		return nil, err
	}

	return readTestResult, nil
}

func getTestResultFromResponse(response interface{}) ([]*TestResult, error) {
	var testResults []*TestResult
	err := decode(&testResults, response)
	return testResults, err
}

func (testResult *TestResult) String() string {
	value, err := json.Marshal(testResult)
	if err != nil {
		return ""
	}

	return string(value)
}

func timeToFloat(t time.Time) float64 {
	if t.IsZero() {
		return 0
	}
	return float64(t.UnixNano()) / 1e9
}
