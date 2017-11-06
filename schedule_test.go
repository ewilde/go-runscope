package runscope

import (
	"testing"
	"strings"
)

func TestCreateSchedule(t *testing.T) {
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

	environment := &Environment{
		Name: "tf_environment",
		InitialVariables: map[string]string{
			"VarA" : "ValB",
		},
	}
	environment, err = client.CreateTestEnvironment(environment, test)
	defer client.DeleteEnvironment(environment, bucket)

	schedule := NewSchedule()
	schedule.Note = "Daily schedule"
	schedule.Interval = "1d"
	schedule.EnvironmentID = environment.ID

	schedule, err = client.CreateSchedule(schedule, bucket.Key, test.ID)
	defer client.DeleteSchedule(schedule, bucket.Key, test.ID)
	if err != nil {
		t.Error(err)
	}

	if len(schedule.ID) == 0 {
		t.Error("Test schedule id should not be empty")
	}
}

func TestReadSchedule(t *testing.T) {
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

	environment := &Environment{
		Name: "tf_environment",
		InitialVariables: map[string]string{
			"VarA" : "ValB",
		},
	}
	environment, err = client.CreateTestEnvironment(environment, test)
	defer client.DeleteEnvironment(environment, bucket)

	schedule := NewSchedule()
	schedule.Note = "Daily schedule"
	schedule.Interval = "1d"
	schedule.EnvironmentID = environment.ID

	schedule, err = client.CreateSchedule(schedule, bucket.Key, test.ID)
	defer client.DeleteSchedule(schedule, bucket.Key, test.ID)
	if err != nil {
		t.Error(err)
	}

	readSchedule, err := client.ReadSchedule(schedule, bucket.Key, test.ID)
	if err != nil {
		t.Error(err)
	}

	if len(readSchedule.ID) == 0 {
		t.Error("Test schedule id should not be empty")
	}

	if readSchedule.ID != schedule.ID {
		t.Errorf("Expected schedule ID %s, actual %s", schedule.ID, readSchedule.ID)
	}

	if readSchedule.Note != schedule.Note {
		t.Errorf("Expected schedule note %s, actual %s", schedule.Note, readSchedule.Note)
	}
}


func TestUpdateSchedule(t *testing.T) {
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

	environment := &Environment{
		Name: "tf_environment",
		InitialVariables: map[string]string{
			"VarA" : "ValB",
		},
	}
	environment, err = client.CreateTestEnvironment(environment, test)
	defer client.DeleteEnvironment(environment, bucket)

	schedule := NewSchedule()
	schedule.Note = "Daily schedule"
	schedule.Interval = "1d"
	schedule.EnvironmentID = environment.ID

	schedule, err = client.CreateSchedule(schedule, bucket.Key, test.ID)
	defer client.DeleteSchedule(schedule, bucket.Key, test.ID)
	if err != nil {
		t.Error(err)
	}

	schedule.Note = "Updated note field"
	_, err = client.UpdateSchedule(schedule, bucket.Key, test.ID)

	readSchedule, err := client.ReadSchedule(schedule, bucket.Key, test.ID)
	if err != nil {
		t.Error(err)
	}

	if readSchedule.Note != "Updated note field" {
		t.Errorf("Expected schedule note %s, actual %s", "Updated note field", readSchedule.Note)
	}
}

func TestDeleteSchedule(t *testing.T) {
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

	environment := &Environment{
		Name: "tf_environment",
		InitialVariables: map[string]string{
			"VarA" : "ValB",
		},
	}
	environment, err = client.CreateTestEnvironment(environment, test)
	defer client.DeleteEnvironment(environment, bucket)

	schedule := NewSchedule()
	schedule.Note = "Daily schedule"
	schedule.Interval = "1d"
	schedule.EnvironmentID = environment.ID


	schedule, err = client.CreateSchedule(schedule, bucket.Key, test.ID)
	if err != nil {
		t.Error(err)
	}

	client.DeleteSchedule(schedule, bucket.Key, test.ID)
	if err != nil {
		t.Error(err)
	}

	_, err = client.ReadSchedule(schedule, bucket.Key, test.ID)
	if err == nil {
		t.Error("Should not have found test schedule after deleting it")
	}

	if !strings.Contains(err.Error(), "404 NOT FOUND") {
		t.Errorf("Expected error to contain %s, actual %s", "404 NOT FOUND", err.Error())
	}
}


