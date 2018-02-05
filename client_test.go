package runscope

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"testing"
)

var teamID string

func TestDeserializeResult(t *testing.T) {
	responseBody := `
	{
	  "meta": {
	    "status": "success"
	  },
	  "data": {
	    "verify_ssl": true,
	    "trigger_url": "https://api.runscope.com/radar/bucket/2e15499d-2e32-4ea8-b6c9-18468031c491/trigger",
	    "name": "foo",
	    "key": "6t0sd3euxlwa",
	    "team": {
	      "name": "form3",
	      "id": "870ed937-bc6e-4d8b-a9a5-d7f9f2412fa3"
	    },
	    "default": false,
	    "auth_token": null,
	    "tests_url": "https://api.runscope.com/buckets/6t0sd3euxlwa/tests",
	    "collections_url": "https://api.runscope.com/buckets/6t0sd3euxlwa/collections",
	    "messages_url": "https://api.runscope.com/buckets/6t0sd3euxlwa/stream"
	  },
	  "error": null
	}
	`
	response := response{}
	err := json.Unmarshal([]byte(responseBody), &response)
	if err != nil {
		t.Error(err)
	}

	if response.Data.(map[string]interface{})["key"] != "6t0sd3euxlwa" {
		t.Error("Key not deserialized")
	}
}

func clientConfigure() *Client {
	return NewClient(APIURL, os.Getenv("RUNSCOPE_ACCESS_TOKEN"))
}

func testPreCheck(t *testing.T) {
	skip := os.Getenv("RUNSCOPE_ACC") == ""
	if skip {
		t.Log("runscope client.go tests require setting RUNSCOPE")
		t.Skip()
	}

	if v := os.Getenv("RUNSCOPE_ACCESS_TOKEN"); v == "" {
		t.Fatal("RUNSCOPE_ACCESS_TOKEN must be set for acceptance tests")
	}

	if v := os.Getenv("RUNSCOPE_TEAM_ID"); v == "" {
		t.Fatal("RUNSCOPE_TEAM_ID must be set for acceptance tests")
	}

	teamID = os.Getenv("RUNSCOPE_TEAM_ID")
}

func deletePredicate(bucket *Bucket) bool {
	if strings.HasPrefix(bucket.Name, "test") || strings.HasSuffix(bucket.Name, "-test") {
		log.Printf("deleting bucket name: %s key: %s \n", bucket.Name, bucket.Key)
		return true
	}
	return false
}

func TestMain(m *testing.M) {

	client := clientConfigure()
	client.DeleteBuckets(deletePredicate)

	code := m.Run()

	client.DeleteBuckets(deletePredicate)

	os.Exit(code)

}
