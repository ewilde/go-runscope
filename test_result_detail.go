package runscope

import (
	"encoding/json"
	"fmt"
	"time"
)

// TestResultDetail retrieve the details of a given test run by ID.
// see https://www.runscope.com/docs/api/results#test-run-detail
type TestResultDetail struct {
	Agent             string      `json:"agent,omitempty"`
	AssertionsDefined uint        `json:"assertions_defined,omitempty"`
	AssertionsFailed  uint        `json:"assertions_failed,omitempty"`
	AssertionsPassed  uint        `json:"assertions_passed,omitempty"`
	BucketKey         string      `json:"bucket_key,omitempty"`
	FinishedAt        *time.Time  `json:"finished_at,omitempty"`
	Region            string      `json:"region,omitempty"`
	Requests          []*TestResultDetailRequest `json:"requests,omitempty"`
	RequestsExecuted  uint        `json:"requests_executed,omitempty"`
	Result            string      `json:"result,omitempty"`
	ScriptsDefined    uint        `json:"scripts_defined,omitempty"`
	ScriptsFailed     uint        `json:"scripts_failed,omitempty"`
	ScriptsPassed     uint        `json:"scripts_passed,omitempty"`
	StartedAt         *time.Time  `json:"started_at,omitempty"`
	TestRunID         string      `json:"test_run_id,omitempty"`
	TestID            string      `json:"test_id,omitempty"`
	VariablesDefined  uint        `json:"variables_defined,omitempty"`
	VariablesFailed   uint        `json:"variables_failed,omitempty"`
	VariablesPassed   uint        `json:"variables_passed,omitempty"`
}

type TestResultDetailRequest struct {
	UUID              string             `json:"uuid,omitempty"`
	Result            string             `json:"result,omitempty"`
	URL               string             `json:"url,omitempty"`
	Method            string             `json:"method,omitempty"`
	AssertionsDefined uint               `json:"assertions_defined,omitempty"`
	AssertionsFailed  uint               `json:"assertions_failed,omitempty"`
	AssertionsPassed  uint               `json:"assertions_passed,omitempty"`
	ScriptsDefined    uint               `json:"scripts_defined,omitempty"`
	ScriptsFailed     uint               `json:"scripts_failed,omitempty"`
	ScriptsPassed     uint               `json:"scripts_passed,omitempty"`
	VariablesDefined  uint               `json:"variables_defined,omitempty"`
	VariablesFailed   uint               `json:"variables_failed,omitempty"`
	VariablesPassed   uint               `json:"variables_passed,omitempty"`
	Assertions        []*TestResultDetailAssertion `json:"assertions,omitempty"`
	Scripts           []*TestResultDetailScript    `json:"scripts,omitempty"`
	Timings           *TestResultDetailTimings     `json:"timings,omitempty"`
	Variables         []*TestResultDetailVariable  `json:"variables,omitempty"`
}

type TestResultDetailAssertion struct {
	Result      string      `json:"result,omitempty"`
	Source      string      `json:"source,omitempty"`
	Property    string      `json:"property,omitempty"`
	Comparison  string       `json:"comparison,omitempty"`
	TargetValue interface{} `json:"target_value,omitempty"`
	ActualValue interface{} `json:"actual_value,omitempty"`
	Error       string      `json:"error,omitempty"`
}

type TestResultDetailScript struct {
	Result string `json:"result,omitempty"`
	Output string `json:"output,omitempty"`
	Error  string `json:"error,omitempty"`
}

type TestResultDetailTimings struct {
	DialMillisecond            float64 `json:"dial_ms,omitempty"`
	DNSLookupMillisecond       float64 `json:"dns_lookup_ms,omitempty"`
	ReceiveResponseMillisecond float64 `json:"receive_response_ms,omitempty"`
	SendBodyMillisecond        float64 `json:"send_body_ms,omitempty"`
	SendHeadersMillisecond     float64 `json:"send_headers_ms,omitempty"`
	WaitForResponseMillisecond float64 `json:"wait_for_response_ms"`
}

type TestResultDetailVariable struct {
	Result   string      `json:"result,omitempty"`
	Source   string      `json:"source,omitempty"`
	Property string      `json:"property,omitempty"`
	Name     string      `json:"name,omitempty"`
	Value    interface{} `json:"value,omitempty"`
	Error    string      `json:"error,omitempty"`
}

// ReadTestResultDetail. Retrieve the details of a given test run by ID.
// See: https://www.runscope.com/docs/api/results#test-run-detail
func (client *Client) ReadTestResultDetail(bucketKey string, testID string, testRunID string) (*TestResultDetail, error) {
	resource, err := client.readResource("test result detail", testRunID,
		fmt.Sprintf("/buckets/%s/tests/%s/results/%s", bucketKey, testID, testRunID))
	if err != nil {
		return nil, err
	}

	readTestResult, err := getTestResultDetailFromResponse(resource.Data)
	if err != nil {
		return nil, err
	}

	return readTestResult, nil
}

// ReadTestResultDetail. Retrieve the details of a given test run by ID.
// See: https://www.runscope.com/docs/api/results#test-run-detail
func (client *Client) ReadTestResultDetailLatest(bucketKey string, testID string) (*TestResultDetail, error) {
	resource, err := client.readResource("latest test result", testID,
		fmt.Sprintf("/buckets/%s/tests/%s/results/latest", bucketKey, testID))
	if err != nil {
		return nil, err
	}

	readTestResult, err := getTestResultDetailFromResponse(resource.Data)
	if err != nil {
		return nil, err
	}

	return readTestResult, nil
}

func getTestResultDetailFromResponse(response interface{}) (*TestResultDetail, error) {
	var testResultDetail *TestResultDetail
	err := decode(&testResultDetail, response)
	return testResultDetail, err
}

func (testResultDetail *TestResultDetail) String() string {
	value, err := json.Marshal(testResultDetail)
	if err != nil {
		return ""
	}

	return string(value)
}
