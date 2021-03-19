package runscope

import (
	"testing"
)

func TestListResults(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()
	bucket, err := client.CreateBucket(&Bucket{Name: "newTest", Team: &Team{ID: teamID}})
	defer client.DeleteBucket(bucket.Key)

	if err != nil {
		t.Error(err)
	}

	newTest := &Test{Name: "tf_test", Description: "This is a tf newTest", Bucket: bucket}
	newTest, err = client.CreateTest(newTest)
	defer client.DeleteTest(newTest)

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

	step, err = client.CreateTestStep(step, bucket.Key, newTest.ID)
	if err != nil {
		t.Error(err)
	}

	defer client.DeleteTestStep(step, bucket.Key, newTest.ID)
	if err != nil {
		t.Error(err)
	}

	listResults, err := client.ListResults(bucket.Key, newTest.ID)
	if err != nil {
		t.Error(err)
	}

	if len(listResults) == 0 {
		t.Error("Expected results but none found")
	}
}
