package runscope

import (
	"strings"
	"testing"
)

func TestCreateTestStep(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()
	bucket, err := client.CreateBucket(&Bucket{Name: "test", Team: &Team{ID: teamID}})
	defer client.DeleteBucket(bucket.Key)

	if err != nil {
		t.Error(err)
	}

	test := &Test{Name: "tf_test", Description: "This is a tf test", Bucket: bucket}
	test, err = client.CreateTest(test)
	defer client.DeleteTest(test)

	if err != nil {
		t.Error(err)
	}

	step := NewTestStep()
	step.StepType = "request"
	step.URL = "http://example.com"
	step.Method = "GET"
	step.Assertions = []*Assertion{{
		Source:     "response_status",
		Comparison: "equal_number",
		Value:      200,
	}}

	step, err = client.CreateTestStep(step, bucket.Key, test.ID)
	defer client.DeleteTestStep(step, bucket.Key, test.ID)
	if err != nil {
		t.Error(err)
	}

	if len(step.ID) == 0 {
		t.Error("Test step id should not be empty")
	}
	
	test2 := &Test{Name: "tf_test2", Description: "This is a tf test with a subtest step", Bucket: bucket}
	test2, err = client.CreateTest(test2)
	defer client.DeleteTest(test2)

	if err != nil {
		t.Error(err)
	}

	step2 := NewTestStep()
	step2.StepType = "subtest"
	step2.TestUUID = test.ID
	step2.Assertions = []*Assertion{{
		Source:     "response_json",
		Property:   "result",
		Comparison: "equal",
		Value:      "pass",
	}}

	step2, err = client.CreateTestStep(step2, bucket.Key, test2.ID)
	defer client.DeleteTestStep(step2, bucket.Key, test2.ID)
	if err != nil {
		t.Error(err)
	}

	if len(step2.ID) == 0 {
		t.Error("Test step id should not be empty")
	}
}

func TestReadTestStep(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()
	bucket, err := client.CreateBucket(&Bucket{Name: "test", Team: &Team{ID: teamID}})
	defer client.DeleteBucket(bucket.Key)

	if err != nil {
		t.Error(err)
	}

	test := &Test{Name: "tf_test", Description: "This is a tf test", Bucket: bucket}
	test, err = client.CreateTest(test)
	defer client.DeleteTest(test)

	if err != nil {
		t.Error(err)
	}

	step := NewTestStep()
	step.Scripts = []string{"log(\"This is a sample post-request script\");"}
	step.BeforeScripts = []string{"log(\"This is a sample pre-request script\");"}
	step.StepType = "request"
	step.URL = "http://example.com"
	step.Method = "GET"
	step.Assertions = []*Assertion{{
		Source:     "response_status",
		Comparison: "equal_number",
		Value:      200,
	}}

	step, err = client.CreateTestStep(step, bucket.Key, test.ID)
	defer client.DeleteTestStep(step, bucket.Key, test.ID)
	if err != nil {
		t.Error(err)
	}

	readStep, err := client.ReadTestStep(step, bucket.Key, test.ID)
	if err != nil {
		t.Error(err)
	}

	if len(readStep.ID) == 0 {
		t.Error("Test step id should not be empty")
	}

	if readStep.ID != step.ID {
		t.Errorf("Expected step ID %s, actual %s", step.ID, readStep.ID)
	}

	if readStep.Method != step.Method {
		t.Errorf("Expected step method %s, actual %s", step.Method, readStep.Method)
	}

	if readStep.BeforeScripts[0] != step.BeforeScripts[0] {
		t.Errorf("Expected before script %s, actual %s", step.BeforeScripts[0], readStep.BeforeScripts[0])
	}

	if readStep.Scripts[0] != step.Scripts[0] {
		t.Errorf("Expected script %s, actual %s", step.Scripts[0], readStep.Scripts[0])
	}
}

func TestUpdateTestStep(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()
	bucket, err := client.CreateBucket(&Bucket{Name: "test", Team: &Team{ID: teamID}})
	defer client.DeleteBucket(bucket.Key)

	if err != nil {
		t.Error(err)
	}

	test := &Test{Name: "tf_test", Description: "This is a tf test", Bucket: bucket}
	test, err = client.CreateTest(test)
	defer client.DeleteTest(test)

	if err != nil {
		t.Error(err)
	}

	step := NewTestStep()
	step.StepType = "request"
	step.URL = "http://example.com"
	step.Method = "GET"
	step.Assertions = []*Assertion{{
		Source:     "response_status",
		Comparison: "equal_number",
		Value:      200,
	}}

	step, err = client.CreateTestStep(step, bucket.Key, test.ID)
	defer client.DeleteTestStep(step, bucket.Key, test.ID)
	if err != nil {
		t.Error(err)
	}

	step.Method = "POST"
	_, err = client.UpdateTestStep(step, bucket.Key, test.ID)

	readStep, err := client.ReadTestStep(step, bucket.Key, test.ID)
	if err != nil {
		t.Error(err)
	}

	if readStep.Method != "POST" {
		t.Errorf("Expected step method %s, actual %s", "POST", readStep.Method)
	}
}

func TestDeleteTestStep(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()
	bucket, err := client.CreateBucket(&Bucket{Name: "test", Team: &Team{ID: teamID}})
	defer client.DeleteBucket(bucket.Key)

	if err != nil {
		t.Error(err)
	}

	test := &Test{Name: "tf_test", Description: "This is a tf test", Bucket: bucket}
	test, err = client.CreateTest(test)
	defer client.DeleteTest(test)

	if err != nil {
		t.Error(err)
	}

	step := NewTestStep()
	step.StepType = "request"
	step.URL = "http://example.com"
	step.Method = "GET"
	step.Assertions = []*Assertion{{
		Source:     "response_status",
		Comparison: "equal_number",
		Value:      200,
	}}

	step, err = client.CreateTestStep(step, bucket.Key, test.ID)
	if err != nil {
		t.Error(err)
	}

	client.DeleteTestStep(step, bucket.Key, test.ID)
	if err != nil {
		t.Error(err)
	}

	_, err = client.ReadTestStep(step, bucket.Key, test.ID)
	if err == nil {
		t.Error("Should not have found test step after deleting it")
	}

	if !strings.Contains(err.Error(), "404 NOT FOUND") {
		t.Errorf("Expected error to contain %s, actual %s", "404 NOT FOUND", err.Error())
	}
}

func TestValidationRequestTypeMissingMethod(t *testing.T) {

	step := NewTestStep()
	step.StepType = "request"
	step.URL = "http://example.com"
	step.Assertions = []*Assertion{{
		Source:     "response_status",
		Comparison: "equal_number",
		Value:      200,
	}}

	client := clientConfigure()
	_, err := client.CreateTestStep(step, "foo", "ba")
	if err == nil {
		t.Error("Expected validation error for missing method")
	}

	if !strings.Contains(err.Error(), "A request test step must specify 'Method' property") {
		t.Error("Expected validation error for missing method")
	}
}

func TestValidationRequestTypeGetIncludesBody(t *testing.T) {

	step := NewTestStep()
	step.StepType = "request"
	step.URL = "http://example.com"
	step.Assertions = []*Assertion{{
		Source:     "response_status",
		Comparison: "equal_number",
		Value:      200,
	}}
	step.Method = "GET"
	step.Body = "foo"

	client := clientConfigure()
	_, err := client.CreateTestStep(step, "foo", "ba")
	if err == nil {
		t.Error("Expected validation error for request with GET method including body")
	}

	if !strings.Contains(err.Error(), "A request test step that specifies a 'GET' method can not include a body property") {
		t.Error("Expected validation error for request GET with included body")
	}
}
