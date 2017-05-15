package runscope

import (
	"testing"
	"encoding/json"
	"time"
)

func TestCreateTest(t *testing.T) {
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

	if len(test.ID) == 0 {
		t.Error("Test id should not be empty")
	}

	if test.CreatedAt.Day() != time.Now().Day() {
		t.Errorf("Expected time %s js not correct", test.CreatedAt.String())
	}
}

func TestReadTest(t *testing.T) {
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

	test.Description = "New description"
	updatedTest, err := client.UpdateTest(test)
	if err != nil {
		t.Error(err)
	}

	if updatedTest.Description != test.Description {
		t.Errorf("Expected description %s, actual %s", test.Description, updatedTest.Description)
	}
}

func TestUpdateTestUsingPartiallyFilledOutObject(t *testing.T) {
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

	testUpdate := &Test{ ID: test.ID, Description: "New description", Bucket: bucket}
	updatedTest, err := client.UpdateTest(testUpdate)
	if err != nil {
		t.Error(err)
	}

	if updatedTest.Description != testUpdate.Description {
		t.Errorf("Expected description %s, actual %s", testUpdate.Description, updatedTest.Description)
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
     "schedules": [
      {
        "environment_id": "d44fe112-74ea-4713-92fc-caa27ef8ce8a",
        "interval": "1m",
        "note": null,
        "version": "1.0",
        "exported_at": 1494623265,
        "id": "c4058b68-3493-44f0-b25e-c33db257e366"
      }
    ],
    "steps": [
      {
        "url": "{{base_url}}/v1/users",
        "variables": [
		{
		    "name": "source_ip",
		    "property": "origin",
		    "source": "response_json"
		}
	    ],
        "args": {},
        "step_type": "request",
        "auth": {},
        "id": "e4044178-3b78-43fd-b67c-3316bfe526a9",
        "note": "some note",
        "headers": {
          "Authorization": [
            "bearer {{token}}"
          ]
        },
        "request_id": "2dbfb5d2-3b5a-499c-9550-b06f9a475feb",
        "assertions": [
          {
            "comparison": "equal_number",
            "value": 200,
            "source": "response_status"
          }
        ],
        "scripts": [
		{
		    "value": "log(\"This is a sample script\");"
		}
	    ],
        "before_scripts": [],
        "data": "",
        "method": "GET"
      }
    ],
    "last_run": {
      "remote_agent_uuid": null,
      "finished_at": 1494623241.385894,
      "error_count": 0,
      "message_success": 1,
      "test_uuid": "7aec8f16-8680-41fe-b0df-4c9be99b3a26",
      "id": "50ded770-f0b5-48ec-91ce-782a932d6b80",
      "extractor_success": 0,
      "uuid": "50ded770-f0b5-48ec-91ce-782a932d6b80",
      "environment_uuid": "d44fe112-74ea-4713-92fc-caa27ef8ce8a",
      "environment_name": "Test Settings",
      "source": "scheduled",
      "remote_agent_name": null,
      "remote_agent": null,
      "status": "completed",
      "bucket_key": "taank6ebawmk",
      "remote_agent_version": "unknown",
      "substitution_success": 0,
      "message_count": 1,
      "script_count": 0,
      "substitution_count": 0,
      "script_success": 0,
      "assertion_count": 1,
      "assertion_success": 1,
      "created_at": 1494623238.460797,
      "messages": [],
      "extractor_count": 0,
      "template_uuids": [
        "699db99e-8c7a-4922-9a0b-a73da87387fb",
        "e4044178-3b78-43fd-b67c-3316bfe526a9"
      ],
      "region": "us1"
    },
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

	if test.CreatedBy.ID != "8512774f-de31-433e-b068-ed76819b2842" {
		t.Errorf("Expected created by %s, actual %s", "8512774f-de31-433e-b068-ed76819b2842", test.CreatedBy.Email)
	}

	if test.DefaultEnvironmentID != "8e7afae4-23b6-492a-b4b9-75d515b5082b" {
		t.Errorf("Expected created by %s, actual %s", "8e7afae4-23b6-492a-b4b9-75d515b5082b", test.DefaultEnvironmentID)
	}

	expectedTime = time.Unix(int64(1294023235), 0)
	if !test.ExportedAt.Equal(expectedTime) {
		t.Errorf("Expected time %s, actual %s",  expectedTime.String(), test.ExportedAt)
	}

	if len(test.Environments) != 1 {
		t.Errorf("Expected %d environments, actual %d",  1, len(test.Environments))
	}

	if test.Environments[0].ID != "7e7afae4-23b6-492a-b4b9-75d515b5082b" {
		t.Errorf("Expected environment id %s, actual %s", "7e7afae4-23b6-492a-b4b9-75d515b5082b", test.Environments[0].ID)
	}

	if test.LastRun == nil {
		t.Error("LastRun nil")
	}

	expectedTime = time.Unix(1494623241, 385894060)
	if !test.LastRun.FinishedAt.Equal(expectedTime) {
		t.Errorf("Expected last run finished at time %s, actual %s",  expectedTime.String(), test.LastRun.FinishedAt)
	}

	if len(test.Steps) != 1 {
		t.Errorf("Expected %d steps, actual %d",  1, len(test.Steps))
	}

	step := test.Steps[0]
	if step.URL != "{{base_url}}/v1/users" {
		t.Errorf("Expected step url %s, actual %s",  "{{base_url}}/v1/users", step.URL)
	}

	if len(step.Variables) != 1 {
		t.Errorf("Expected %d variables, actual %d",  1, len(step.Variables))
	}

	variable := step.Variables[0]
	if variable.Name != "source_ip" {
		t.Errorf("Expected variable name %s, actual %s",  "source_ip", variable.Name)
	}

	if variable.Property != "origin" {
		t.Errorf("Expected variable property %s, actual %s",  "origin", variable.Property)
	}

	if variable.Source != "response_json" {
		t.Errorf("Expected variable source %s, actual %s",  "origin", variable.Source)
	}

	if step.StepType != "request" {
		t.Errorf("Expected step type %s, actual %s",  "request", step.StepType)
	}

	if step.ID != "e4044178-3b78-43fd-b67c-3316bfe526a9" {
		t.Errorf("Expected step type %s, actual %s",  "e4044178-3b78-43fd-b67c-3316bfe526a9", step.StepType)
	}

	if len(step.Headers) != 1 {
		t.Errorf("Expected %d headers, actual %d",  1, len(step.Headers))
	}

	header := step.Headers["Authorization"]
	if len(header) != 1 {
		t.Errorf("Expected %d authorization values, actual %d",  1, len(header))
	}

	if header[0] != "bearer {{token}}" {
		t.Errorf("Expected authorization header %s, actual %s", "bearer {{token}}", header[0])
	}

	if step.RequestId != "2dbfb5d2-3b5a-499c-9550-b06f9a475feb" {
		t.Errorf("Expected step request id %s, actual %s",  "2dbfb5d2-3b5a-499c-9550-b06f9a475feb", step.StepType)
	}

	if len(step.Assertions) != 1 {
		t.Errorf("Expected %d assertions, actual %d",  1, len(step.Assertions))
	}

	assertion := step.Assertions[0]
	if assertion.Comparison != "equal_number" {
		t.Errorf("Expected assertion comparison %s, actual %s", "equal_number", assertion.Comparison)
	}

	if assertion.Value != float64(200) {
		t.Errorf("Expected assertion value %d, actual %d", 200, assertion.Value)
	}

	if assertion.Source != "response_status" {
		t.Errorf("Expected assertion source %s, actual %s", "response_status", assertion.Source)
	}

	if len(step.Scripts) != 1 {
		t.Errorf("Expected %d scripts, actual %d",  1, len(step.Scripts))
	}

	script := step.Scripts[0]

	if script.Value != "log(\"This is a sample script\");" {
		t.Errorf("Expected script value %s, actual %s", "log(\"This is a sample script\");", script.Value)
	}
}