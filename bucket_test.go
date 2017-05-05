package runscope

import (
	"testing"
	"fmt"
	"encoding/json"
)

func TestCreateBucket(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()
	bucket, err := client.CreateBucket(Bucket{Name: "test", Team: Team{Id: teamId}})

	if err != nil {
		t.Error(err)
	}

	client.DeleteBucket(bucket.Key)
}

func TestReadBucket(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()

	createdBucket, err := client.CreateBucket(Bucket{Name: "terraform-client.go-test", Team: Team{Id: teamId}})
	if err != nil {
		t.Error(err)
	}

	readBucket, err := client.ReadBucket(createdBucket.Key)
	if err != nil {
		t.Error(err)
	}

	if readBucket.Key != createdBucket.Key {
		t.Errorf("Bucket createdBucket expected %s was %s.", createdBucket.Key, readBucket.Key)
	}

	if readBucket.TestsUrl != fmt.Sprintf("https://api.runscope.com/buckets/%s/tests", readBucket.Key) {
		t.Errorf("Bucket url expected %s was %s.",
			fmt.Sprintf("https://api.runscope.com/buckets/%s/tests", readBucket.Key), readBucket.TestsUrl)
	}

	client.DeleteBucket(createdBucket.Key)
}


func TestBucketReadFromResponse(t *testing.T) {
	responseBody := `
{
  "meta": {
    "status": "success"
  },
  "data": {
    "verify_ssl": true,
    "trigger_url": "https://api.runscope.com/radar/bucket/f2f4dbbb-7bf0-4528-bf51-eb3d06a20423/trigger",
    "name": "Sample Name",
    "key": "z3n32gktzx94",
    "team": {
      "name": "form3",
      "id": "870ed937-bc6e-4d8b-a9a5-d7f9f2412fa3"
    },
    "default": false,
    "auth_token": null,
    "tests_url": "https://api.runscope.com/buckets/z3n32gktzx94/tests",
    "collections_url": "https://api.runscope.com/buckets/z3n32gktzx94/collections",
    "messages_url": "https://api.runscope.com/buckets/z3n32gktzx94/stream"
  },
  "error": null
}
`
	responseMap := make(map[string]interface{})
	if err := json.Unmarshal([]byte(responseBody), &responseMap); err != nil {
		t.Error(err)
	}

	dataMap := responseMap["data"].(map[string]interface{})
	bucket, err := getBucketFromResponse(dataMap)
	if err != nil {
		t.Error(err)
	}

	if bucket.Name != "Sample Name" {
		t.Errorf("Expected name %s, actual %s", "Sample Name", bucket.Name)
	}

	if len(bucket.TestsUrl) == 0 {
		t.Error("Missing test url")
	}
}