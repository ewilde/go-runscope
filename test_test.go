package runscope

import (
	"testing"
	"encoding/json"
)

func TestCreateTest(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()
	bucket, err := client.CreateBucket(Bucket{Name: "test", Team: Team{Id: teamId}})
	defer client.DeleteBucket(bucket.Key)

	if err != nil {
		t.Error(err)
	}

	test := Test{Name: "tf_test", Description: "This is a tf test", Bucket: bucket}
	test, err = client.CreateTest(test)
	defer client.DeleteTest(test)

	if err != nil {
		t.Error(err)
	}

	if len(test.Id) == 0 {
		t.Error("Test id should not be empty")
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

	newTest := Test{Name: "tf_test", Description: "This is a tf newTest", Bucket: bucket}
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
}

func TestUpdateTest(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()
	bucket, err := client.CreateBucket(Bucket{Name: "test", Team: Team{Id: teamId}})
	defer client.DeleteBucket(bucket.Key)

	if err != nil {
		t.Error(err)
	}

	test := Test{Name: "tf_test", Description: "This is a tf test", Bucket: bucket}
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
      "email": "edward.wilde@form3.tech",
      "id": "8512774f-de31-433e-b068-ed76819b2842",
      "name": "Edward Wilde"
    },
    "default_environment_id": "7e7afae4-23b6-492a-b4b9-75d515b5082b",
    "version": "1.0",
    "exported_at": 1494023235,
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
    "id": "e0699a4b-0141-4fa2-8007-e016acede2bf",
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
}