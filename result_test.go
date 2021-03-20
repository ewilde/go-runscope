package runscope

import (
	"testing"
	"time"
)

func TestListResults(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()
	bucket, err := client.CreateBucket(&Bucket{Name: "newTest", Team: &Team{ID: teamID}})
	defer client.DeleteBucket(bucket.Key) // nolint: errcheck

	if err != nil {
		t.Error(err)
	}

	newTest := &Test{Name: "tf_test", Description: "This is a tf newTest", Bucket: bucket}
	newTest, err = client.CreateTest(newTest)
	defer client.DeleteTest(newTest) // nolint: errcheck

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

	_, err = client.CreateTestStep(step, bucket.Key, newTest.ID)
	if err != nil {
		t.Error(err)
	}

	defer client.DeleteTestStep(step, bucket.Key, newTest.ID) // nolint: errcheck
	if err != nil {
		t.Error(err)
	}

	schedule := NewSchedule()
	schedule.Note = "Daily schedule"
	schedule.Interval = "1m"
	schedule.EnvironmentID = newTest.DefaultEnvironmentID

	schedule, err = client.CreateSchedule(schedule, bucket.Key, newTest.ID)
	if err != nil {
		t.Error(err)
	}
	defer client.DeleteSchedule(schedule, bucket.Key, newTest.ID) // nolint: errcheck
	time.Sleep(1 * time.Minute)

	listResults, err := client.ListResults(bucket.Key, newTest.ID)
	if err != nil {
		t.Error(err)
	}

	if len(listResults) == 0 {
		t.Error("Expected results but none found")
	}
}

func TestReadTestLatestResult(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()
	bucket, err := client.CreateBucket(&Bucket{Name: "newTest", Team: &Team{ID: teamID}})
	defer client.DeleteBucket(bucket.Key) // nolint: errcheck

	if err != nil {
		t.Error(err)
	}

	newTest := &Test{Name: "tf_test", Description: "This is a tf newTest", Bucket: bucket}
	newTest, err = client.CreateTest(newTest)
	defer client.DeleteTest(newTest) // nolint: errcheck

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

	_, err = client.CreateTestStep(step, bucket.Key, newTest.ID)
	if err != nil {
		t.Error(err)
	}

	defer client.DeleteTestStep(step, bucket.Key, newTest.ID) // nolint: errcheck
	if err != nil {
		t.Error(err)
	}

	schedule := NewSchedule()
	schedule.Note = "Daily schedule"
	schedule.Interval = "1m"
	schedule.EnvironmentID = newTest.DefaultEnvironmentID

	schedule, err = client.CreateSchedule(schedule, bucket.Key, newTest.ID)
	if err != nil {
		t.Error(err)
	}
	defer client.DeleteSchedule(schedule, bucket.Key, newTest.ID) // nolint: errcheck
	time.Sleep(1 * time.Minute)

	testResult, err := client.ReadTestLatestResult(newTest.ID, bucket.Key)
	if err != nil {
		t.Error(err)
	}

	if testResult.Result != "pass" {
		t.Errorf("Expected successfull test %s, actual %s", "pass", testResult.Result)
	}
}
