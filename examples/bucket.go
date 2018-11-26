package examples

import (
	"fmt"
	"github.com/ewilde/go-runscope"
	"log"
)

var accessToken = "{your token}"  // See https://www.runscope.com/applications
var teamUUID = "{your team uuid}" // See https://www.runscope.com/teams
var client = runscope.NewClient(runscope.APIURL, accessToken)

func createBucket() *runscope.Bucket {
	var bucket = &runscope.Bucket{
		Name: "My first bucket",
		Team: &runscope.Team{
			ID: teamUUID,
		},
	}

	bucket, err := client.CreateBucket(bucket)
	if err != nil {
		DebugF(1, "[ERROR] error creating bucket: %s", err)
	}

	fmt.Printf("Bucket created successfully: %s", bucket.String())
	return bucket
}

func readBucket() {
	bucket, err := client.ReadBucket("htqee6p4dhvc")
	if err != nil {
		DebugF(1, "[ERROR] error creating bucket: %s", err)
	}

	fmt.Printf("Bucket read successfully: %s", bucket.String())
}

func deleteBucket() {
	err := client.DeleteBucket("htqee6p4dhvc")
	if err != nil {
		DebugF(1, "[ERROR] error creating bucket: %s", err)
	}
}
