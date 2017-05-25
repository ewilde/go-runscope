package runscope

import (
	"testing"
	"strings"
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
	step.Assertions = [] Assertion {{
		Source: "response_status",
		Comparison : "equal_number",
		Value: 200,
	}}

	step, err = client.CreateTestStep(step, bucket.Key, test.ID)
	defer client.DeleteTestStep(step, bucket.Key, test.ID)
	if err != nil {
		t.Error(err)
	}

	if len(step.ID) == 0 {
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
	step.StepType = "request"
	step.URL = "http://example.com"
	step.Method = "GET"
	step.Assertions = [] Assertion {{
		Source: "response_status",
		Comparison : "equal_number",
		Value: 200,
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
	step.Assertions = [] Assertion {{
		Source: "response_status",
		Comparison : "equal_number",
		Value: 200,
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
	step.Assertions = [] Assertion {{
		Source: "response_status",
		Comparison : "equal_number",
		Value: 200,
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
	step.Assertions = [] Assertion {{
		Source: "response_status",
		Comparison : "equal_number",
		Value: 200,
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


