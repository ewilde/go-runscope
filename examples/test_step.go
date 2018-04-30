package examples

import (
	"github.com/ewilde/go-runscope"
	"log"
)

func createTestWithStep() {
	bucket := createBucket()
	defer client.DeleteBucket(bucket.Key)

	test := &runscope.Test{Name: "tf_test", Description: "This is a tf test", Bucket: bucket}
	test, err := client.CreateTest(test)

	if err != nil {
		log.Fatal(err)
	}

	step := runscope.NewTestStep()
	step.StepType = "request"
	step.URL = "http://example.com"
	step.Method = "GET"
	step.Assertions = []*runscope.Assertion{{
		Source:     "response_status",
		Comparison: "equal_number",
		Value:      200,
	}}

	step, err = client.CreateTestStep(step, bucket.Key, test.ID)
	if err != nil {
		log.Fatal(err)
	}
}
