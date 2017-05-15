package runscope

import (
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

	step := NewTestStep();
	step.StepType = "request"
	step.URL = "http://example.com"
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


