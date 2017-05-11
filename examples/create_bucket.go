package examples

import "github.com/ewilde/go-runscope"

func createBucket() {
	var accessToken = "{your token}"  // See https://www.runscope.com/applications
	var teamUUID = "{your team uuid}" // See https://www.runscope.com/teams
	var client = runscope.NewClient(runscope.APIURL, accessToken)
	var bucket = &runscope.Bucket{
		Name: "My first bucket",
		Team: &runscope.Team{
			ID: teamUUID,
		},
	}


	client.CreateBucket()
}
