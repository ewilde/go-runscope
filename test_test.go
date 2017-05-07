package runscope

import (
	"testing"
	"encoding/json"
	"time"
)

func TestCreateTest(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()
	bucket, err := client.CreateBucket(Bucket{Name: "test", Team: Team{Id: teamId}})
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

	if len(test.Id) == 0 {
		t.Error("Test id should not be empty")
	}

	if test.CreatedAt.Day() != time.Now().Day() {
		t.Errorf("Expected time %s js not correct", test.CreatedAt.String())
	}
}

func TestReadTest(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()
	bucket, err := client.CreateBucket(Bucket{Name: "newTest", Team: Team{Id: teamId}})
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

	readTest, err := client.ReadTest(newTest)
	if err != nil {
		t.Error(err)
	}

	if readTest.Name != newTest.Name {
		t.Errorf("Expected name %s, actual %s", newTest.Name, readTest.Name)
	}

	if readTest.CreatedAt.Day() != time.Now().Day() {
		t.Errorf("Expected time %s js not correct", readTest.CreatedAt.String())
	}
}

func TestUpdateTest(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()
	bucket, err := client.CreateBucket(Bucket{Name: "test", Team: Team{Id: teamId}})
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

	test.Description = "New description"
	resource, err := client.UpdateTest(test)
	if err != nil {
		t.Error(err)
	}

	if resource.Data["description"] != test.Description {
		t.Errorf("Expected description %s, actual %s", test.Description, resource.Data["description"])
	}
}


func TestReadFromResponse(t *testing.T) {
	responseBody := `
{
  "meta": {
    "status": "success"
  },
  "data": {
    "trigger_url": "https://api.runscope.com/radar/235d992e-cb62-43b1-9983-dd5864b64d89/trigger",
    "name": "Sample Name",
    "created_at": 1494023235,
    "created_by": {
      "email": "edward.wilde@acme.com",
      "id": "8512774f-de31-433e-b068-ed76819b2842",
      "name": "Edward Wilde"
    },
    "default_environment_id": "8e7afae4-23b6-492a-b4b9-75d515b5082b",
    "version": "1.0",
    "exported_at": 1294023235,
    "environments": [
      {
        "script_library": [],
        "name": "Test Settings",
        "script": null,
        "preserve_cookies": false,
        "test_id": "e0699a4b-0141-4fa2-8007-e016acede2bf",
        "initial_variables": null,
        "integrations": [],
        "auth": null,
        "id": "7e7afae4-23b6-492a-b4b9-75d515b5082b",
        "regions": [
          "us1"
        ],
        "headers": null,
        "verify_ssl": true,
        "version": "1.0",
        "exported_at": 1494023235,
        "retry_on_failure": false,
        "remote_agents": [],
        "webhooks": null,
        "parent_environment_id": null,
        "stop_on_failure": false,
        "emails": {
          "notify_on": null,
          "notify_all": false,
          "recipients": [],
          "notify_threshold": 0
        },
        "client_certificate": null
      }
    ],
    "schedules": [],
    "steps": [],
    "id": "f0699a4b-0141-4fa2-8007-e016acede2bf",
    "description": "My test description"
  },
  "error": null
}
`
	responseMap := new(response)
	if err := json.Unmarshal([]byte(responseBody), &responseMap); err != nil {
		t.Error(err)
	}

	test, err := getTestFromResponse(responseMap.Data)
	if err != nil {
		t.Error(err)
	}

	if test.Name != "Sample Name" {
		t.Errorf("Expected name %s, actual %s", "Sample Name", test.Name)
	}

	expectedTime := time.Time{}
	expectedTime = time.Unix(int64(1494023235), 0)
	if !test.CreatedAt.Equal(expectedTime) {
		t.Errorf("Expected time %s, actual %s",  expectedTime.String(), test.CreatedAt)
	}

	if test.CreatedBy.Name != "Edward Wilde" {
		t.Errorf("Expected created by %s, actual %s", "Edward Wilde", test.CreatedBy.Name)
	}

	if test.CreatedBy.Email != "edward.wilde@acme.com" {
		t.Errorf("Expected created by %s, actual %s", "edward.wilde@acme.com", test.CreatedBy.Email)
	}

	if test.CreatedBy.Id != "8512774f-de31-433e-b068-ed76819b2842" {
		t.Errorf("Expected created by %s, actual %s", "8512774f-de31-433e-b068-ed76819b2842", test.CreatedBy.Email)
	}

	if test.DefaultEnvironment != "8e7afae4-23b6-492a-b4b9-75d515b5082b" {
		t.Errorf("Expected created by %s, actual %s", "8e7afae4-23b6-492a-b4b9-75d515b5082b", test.DefaultEnvironment)
	}

	expectedTime = time.Unix(int64(1294023235), 0)
	if !test.ExportedAt.Equal(expectedTime) {
		t.Errorf("Expected time %s, actual %s",  expectedTime.String(), test.ExportedAt)
	}

}